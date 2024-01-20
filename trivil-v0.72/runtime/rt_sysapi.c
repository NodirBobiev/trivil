#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <errno.h>
#include "rt_sysapi.h"

#if defined(_WIN32) || defined(_WIN64)

#include <direct.h>

#else

#include <dirent.h>
#include <sys/stat.h>
#include <sys/types.h>

#if defined(__apple_build_version__)
#include <sys/syslimits.h>
#endif

#if defined(__FreeBSD__)
#include <sys/param.h>
#endif

#endif

struct BytesDesc { TInt64 len; TInt64 capacity; TByte* body; };

//=== платформа

EXPORTED TString sysapi_os_kind() {

#if defined(_WIN32) || defined(_WIN64)
    return tri_newString(7, 7, "windows");
#endif

#if defined(__apple_build_version__)
    return tri_newString(6, 6, "darwin");
#endif

#if defined(__linux__)
    return tri_newString(5, 5, "linux");
#endif

#if defined(__FreeBSD__)
    return tri_newString(7, 7, "freebsd");
#endif

}

EXPORTED TBool sysapi_exec(TString cmd) {
    int res = system((char *)cmd->body);
    return res == 0;    
}

//=== вещественные (временно)

TBool sysapi_string_to_float64(TString s, TFloat64* res)  {
    
    char *endptr;
    *res = strtod((char *)s->body, &endptr);
    
    //printf("%s: %f %p %p\n", (char *)s->body, *res, endptr, s->body);
    if (endptr != NULL) {
        if (endptr < (char *)s->body + s->bytes)  return false;
    }

    /* If the result is 0, test for an error */
    if (*res == 0)
    {
        /* If the value provided was out of range, display a warning message */
        if (errno == ERANGE || errno == EINVAL) return false;
    }
    return true;
}

EXPORTED TString sysapi_float64_to_string(TString format, TFloat64 f)  {
    char buf[100];
    int len = snprintf(buf, 100, (char*)format->body, f);
   
    return tri_newString(len, -1, buf); 
}

//==== коды ошибок общие ===

TString error_id(int errcode) {
    char buf[80];
    
    switch (errcode) {

    case ENOENT:
#if defined(_WIN32) || defined(_WIN64)
      strcpy_s(buf, 80, "ФАЙЛ-НЕ-НАЙДЕН"); break;
#else
      strcpy(buf, "ФАЙЛ-НЕ-НАЙДЕН"); break;
#endif
    default:
        sprintf(buf, "ОШИБКА[%d]", errcode); 
    }
    return tri_newString(strlen(buf), -1, buf);
}

//==== папки ====

EXPORTED TString sysapi_exec_path() {
    TString executable = tri_arg(0);

#if defined(_WIN32) || defined(_WIN64)
    size_t len = executable->bytes;
    char* s = nogc_alloc(len + 1);
    strncpy_s(s, len+1, (char*)executable->body, len);
    s[len] = 0; 

    for (int i = 0; i < executable->bytes; i++) {
        if (s[i] == '\\') s[i] = '/';
    }   
    executable =  tri_newString(len, -1, s);
    nogc_free(s);
#else
    char actualPath[PATH_MAX];
    char *ptr;

    ptr = realpath((char*)executable->body, actualPath);

    if (ptr == NULL) {
      return executable;
    } else{
      return tri_newString(strlen(ptr), -1, ptr);;
    }
#endif

    return executable;
}

//==== чтение/запись ====

EXPORTED void* sysapi_fread(void* request, TString filename) {
    
    struct Request* req = request;

    FILE* fp;

#if defined(_WIN32) || defined(_WIN64)
    int errcode = fopen_s(&fp, (char *)filename->body, "rb");
    if (errcode != 0) {
        req->err_id = error_id(errcode);
        return NULL;
    }
#else
    fp = fopen((char *)filename->body, "rb");

    if (fp == NULL) {
        req->err_id = error_id(errno);
        return NULL;
    }
#endif

    fseek(fp, 0, SEEK_END);
    size_t sz = ftell(fp);
    rewind(fp);
    
    struct BytesDesc* bytes = tri_newVector(sizeof(TByte), sz, 0); 
    
    size_t ret = fread(bytes->body, sizeof(TByte), sz, fp);
    
    if (ret != sz) {
            req->err_id = error_id(ferror(fp));
            // TODO: удалить bytes?
            fclose(fp);
            return NULL;
    }

    fclose(fp);
    req->err_id = NULL;
    return bytes;
}

EXPORTED void sysapi_fwrite(void* request, TString filename, void* bytes) {
    
    struct Request* req = request;
    
    FILE* fp;

#if defined(_WIN32) || defined(_WIN64)
    int errcode = fopen_s(&fp, (char *)filename->body, "wb");
    if (errcode != 0) {
        req->err_id = error_id(errcode);
        return;
    }
#else
    fp = fopen((char *)filename->body, "wb");
    if (fp == NULL) {
        req->err_id = error_id(errno);
        return;
    }
#endif

    struct BytesDesc* v = bytes;
    
    size_t ret = fwrite(v->body, sizeof(TByte), v->len, fp);
    
    if (ret != v->len) {
            req->err_id = error_id(ferror(fp));
            // TODO: удалить bytes?
            fclose(fp);
            return;
    }

    fclose(fp);
    req->err_id = NULL;
}

EXPORTED TBool sysapi_make_dir(void* request, TString folder) {
    struct Request* req = request;

    int ret;

#if defined(_WIN32) || defined(_WIN64)
    ret = _mkdir((char*)folder->body);
#else
    ret = mkdir((char*)folder->body, S_IRWXU | S_IRWXG | S_IRWXO); // Права 777, но с учетом umask
#endif

    if (ret == 0) {
        req->err_id = NULL;
	return true;
    } else {
        req->err_id = error_id(errno);
	return false;
    }
}


// ==============   linux     ==============

#ifndef _WIN32

EXPORTED TBool sysapi_is_dir(void* request, TString filename)  {
    struct Request* req = request;    

    struct stat sb;
    req->err_id = NULL;

    // printf("sysapi_is_dir: %s\n", filename->body);
    
    if (stat((char*)filename->body, &sb) == 0 && S_ISDIR(sb.st_mode)) {
	return true;
    } else {
        return false;
    }
}

EXPORTED void* sysapi_dirnames(void* request, TString filename)  {
    struct Request* req = request;

    TInt64 count = 0;

    DIR *d;
    struct dirent *dir;
    d = opendir((char*)filename->body);

    if (d) {
        while ((dir = readdir(d)) != NULL) {
            if (strcmp(dir->d_name, ".") != 0 && (strcmp(dir->d_name, "..") != 0)) {
                count++;
            }
        }
        closedir(d);
    }

    void* list = tri_newVector(sizeof(TString), 0, count);

    d = opendir((char*)filename->body);

    if (d) {
        while ((dir = readdir(d)) != NULL) {
            if (strcmp(dir->d_name, ".") != 0 && (strcmp(dir->d_name, "..") != 0)) {
	        TString s = tri_newString(strlen(dir->d_name), -1, dir->d_name);
	        tri_vectorAppend(list, sizeof(TString), 1, &s);
            }
        }
        closedir(d);
    }

    req->err_id = NULL;
    return list;
}

EXPORTED TString sysapi_abs_path(void* request, TString filename) {
    struct Request* req = request;    
    
    char actualPath[PATH_MAX];
    char *ptr;

    ptr = realpath((char*)filename->body, actualPath);

    if (ptr == NULL) {
      char buf[120];
      sprintf(buf, "Файл %s не найден", filename->body);
      req->err_id = tri_newString(strlen(buf), -1, buf);
      return NULL;
    } else{
      req->err_id = NULL;
      return tri_newString(strlen(ptr), -1, ptr);;
    }
}

EXPORTED TBool sysapi_set_permissions(void* request, TString path, TInt64 permissions) {
    struct Request* req = request;

    int ret = chmod((char*)path->body, (int)(permissions&0x7FFFFFFF));

    if (ret == 0) {
        req->err_id = NULL;
	return true;
    } else {
        req->err_id = error_id(errno);
	return false;
    }
}


// ============== windows ==============
#else
    
#include <windows.h>

TString win_error_id(int errcode) {
    char buf[80];
    
    switch (errcode) {
//    case ENOENT: strcpy_s(buf, 80, "ФАЙЛ-НЕ-НАЙДЕН"); break;
    default:
        sprintf(buf, "ОШИБКА[%d]", errcode); 
    }
    return tri_newString(strlen(buf), -1, buf);
}

EXPORTED TBool sysapi_is_dir(void* request, TString filename)  {
    struct Request* req = request;    
    
    WIN32_FIND_DATA FindFileData;
    
    HANDLE hFind = FindFirstFileA((char*)filename->body, &FindFileData);
    if (hFind == INVALID_HANDLE_VALUE) {
        req->err_id = win_error_id(GetLastError());
        return false;
    } 
    
    if (FindFileData.dwFileAttributes & FILE_ATTRIBUTE_DIRECTORY) {
        return true;
    }
    return false;
}

EXPORTED void* sysapi_dirnames(void* request, TString filename)  {
    
     //printf ("dirnames %s\n", filename->body);    
    
    struct Request* req = request;    
    
    // подготовить образец
    size_t len = filename->bytes;
    char* pattern = nogc_alloc(len + 3);
    
    strncpy_s(pattern, len+3, (char*)filename->body, len);
    pattern[len+0] = '\\';
    pattern[len+1] = '*';
    pattern[len+2] = 0;
    
    for (int i = 0; i < len; i++) {
        if (pattern[i] == '/') pattern[i] = '\\';
    }    

//    printf ("!!!%s\n", pattern);    
    
    //=== Считаю число имен
    WIN32_FIND_DATA FindFileData;
    TInt64 count = 0;
    
    HANDLE hFind = FindFirstFileA(pattern, &FindFileData);
    if (hFind == INVALID_HANDLE_VALUE) {
        req->err_id = win_error_id(GetLastError());
        nogc_free(pattern);
        return NULL;
    } 

    do {
        if (strcmp(FindFileData.cFileName, ".") == 0 || strcmp(FindFileData.cFileName, "..") == 0) {
            // игнорирую
        } else {
            count++;
        }
        //printf("!name: %s\n", FindFileData.cFileName);
    } while (FindNextFile(hFind, &FindFileData) != 0);
    FindClose(hFind);

    //=== Собираю вектор
    hFind = FindFirstFileA(pattern, &FindFileData);
    if (hFind == INVALID_HANDLE_VALUE) {
        req->err_id = win_error_id(GetLastError());
        nogc_free(pattern);
        return NULL;
    } 
    
    void* list = tri_newVector(sizeof(TString), 0, count);

    do {
        if (strcmp(FindFileData.cFileName, ".") == 0 || strcmp(FindFileData.cFileName, "..") == 0) {
            // игнорирую
        } else {
            TString s = tri_newString(strlen(FindFileData.cFileName), -1, FindFileData.cFileName);
            tri_vectorAppend(list, sizeof(TString), 1, &s);            
        }
        //printf("!!name: %s\n", FindFileData.cFileName);
    } while (FindNextFile(hFind, &FindFileData) != 0);
    FindClose(hFind);

    req->err_id = NULL;
    nogc_free(pattern);
    return list;
}

EXPORTED TString sysapi_abs_path(void* request, TString filename) {
    struct Request* req = request;    
    
    DWORD retval = GetFullPathNameA(
        (char*)filename->body,
        0,
        NULL,
        NULL);
    
    if (retval == 0) {
        req->err_id = win_error_id(GetLastError());
        return NULL;
    }

    char* buf = nogc_alloc(retval);

    retval = GetFullPathNameA(
        (char*)filename->body,
        retval,
        buf,
        NULL);

    if (retval == 0) {
        tri_crash("sysapi_abs_path: assert 2nd call of GetFullPathNameA returns error", "");
        return NULL;
    }
    
    for (int i = 0; i < retval; i++) {
        if (buf[i] == '\\') buf[i] = '/';
    }       
    
    TString full =  tri_newString(retval, -1, buf); 
    nogc_free(buf);
    
    return full;
}

EXPORTED TBool sysapi_set_permissions(void* request, TString path, TInt64 permissions) {
    // Не имеет смысла на Windows
    return true;
}

#endif
