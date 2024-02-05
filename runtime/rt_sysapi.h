#ifndef _rt_sysapi_H
#define _rt_sysapi_H

#include "rt_api.h"

//=== платформа

// Выдает вид ОС: windows | linux | unknown
EXPORTED TString sysapi_os_kind();

// Выполняет команду, возращает истину в случае успешного выполнения
EXPORTED TBool sysapi_exec(TString cmd);

//=== вещественные (временно)

// Возвращает истину, если корректное представление вещественного целиком занимает строку
EXPORTED TBool sysapi_string_to_float64(TString s, TFloat64* res) ;

// Выдает строку по указанному формату (как в printf)
EXPORTED TString sysapi_float64_to_string(TString format, TFloat64 f) ;

//==== Используется для работы с файлами
struct Request {
    _BaseObject _base;
    //FILE* handler;
    TString err_id;
};

//==== имена файлов

// Выдает строку - папку программы с правильными разделителями
EXPORTED TString sysapi_exec_path();

// По имени файла выдает аболютное имя файла (от корня файловой системы) 
EXPORTED TString sysapi_abs_path(void* request, TString filename);

//=== чтение/запись

// Читает файл, возвращает дескриптор байтового вектора.
// В случае ошибки, возвращает NULL и выставляет код ошибки в запросе 
EXPORTED void* sysapi_fread(void* request, TString filename);

// Записывает в файл байтовый вектор.
// В случае ошибки выставляет код ошибки в запросе
EXPORTED void sysapi_fwrite(void* request, TString filename, void* bytes);

EXPORTED TBool sysapi_set_permissions(void* request, TString path, TInt64 permissions);

//=== работа с папками

// Выдает true, если файл является папкой 
// В случае ошибки выставляет код ошибки в запросе
// На Линукс true возвращается, только если папка существует в файловой системе
EXPORTED TBool sysapi_is_dir(void* request, TString filename);

// Выдает список имен в папке - список строк []Строка
// В случае ошибки выставляет код ошибки в запросе
EXPORTED void* sysapi_dirnames(void* request, TString filename);

// Создает папку. Возвращает true, если папка создана. Возвращает
// false, если произошла ошибка создания, в том числе, если папка уже
// существует.
EXPORTED TBool sysapi_make_dir(void* request, TString folder);

#endif
