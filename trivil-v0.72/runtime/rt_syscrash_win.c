#include <signal.h>
#include <stdio.h>
#include "rt_defs.h"

// ============== windows ==============
  
#include <windows.h>
#include <imagehlp.h>
#include <stdlib.h>
#include <stdbool.h>


LONG WINAPI windows_crash_handler(EXCEPTION_POINTERS* ExceptionInfo) {
  switch (ExceptionInfo->ExceptionRecord->ExceptionCode) {
  case EXCEPTION_ACCESS_VIOLATION:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_ACCESS_VIOLATION");
    break;
  case EXCEPTION_ARRAY_BOUNDS_EXCEEDED:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_ARRAY_BOUNDS_EXCEEDED");
    break;
  case EXCEPTION_BREAKPOINT:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_BREAKPOINT");
    break;
  case EXCEPTION_DATATYPE_MISALIGNMENT:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_DATATYPE_MISALIGNMENT");
    break;
  case EXCEPTION_FLT_DENORMAL_OPERAND:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_FLT_DENORMAL_OPERAND");
    break;
  case EXCEPTION_FLT_DIVIDE_BY_ZERO:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_FLT_DIVIDE_BY_ZERO");
    break;
  case EXCEPTION_FLT_INEXACT_RESULT:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_FLT_INEXACT_RESULT");
    break;
  case EXCEPTION_FLT_INVALID_OPERATION:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_FLT_INVALID_OPERATION");
    break;
  case EXCEPTION_FLT_OVERFLOW:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_FLT_OVERFLOW");
    break;
  case EXCEPTION_FLT_STACK_CHECK:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_FLT_STACK_CHECK");
    break;
  case EXCEPTION_FLT_UNDERFLOW:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_FLT_UNDERFLOW");
    break;
  case EXCEPTION_ILLEGAL_INSTRUCTION:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_ILLEGAL_INSTRUCTION");
    break;
  case EXCEPTION_IN_PAGE_ERROR:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_IN_PAGE_ERROR");
    break;
  case EXCEPTION_INT_DIVIDE_BY_ZERO:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_INT_DIVIDE_BY_ZERO");
    break;
  case EXCEPTION_INT_OVERFLOW:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_INT_OVERFLOW");
    break;
  case EXCEPTION_INVALID_DISPOSITION:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_INVALID_DISPOSITION");
    break;
  case EXCEPTION_NONCONTINUABLE_EXCEPTION:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_NONCONTINUABLE_EXCEPTION");
    break;
  case EXCEPTION_PRIV_INSTRUCTION:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_PRIV_INSTRUCTION");
    break;
  case EXCEPTION_SINGLE_STEP:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_SINGLE_STEP");
    break;
  case EXCEPTION_STACK_OVERFLOW:
    fprintf(stderr, "Sys crash: %s\n", "EXCEPTION_STACK_OVERFLOW");
    break;
    /*
      case :
      fprintf(stderr, "Sys crash: %s\n", "");
      break;
    */

  default:
    fprintf(stderr, "Unrecognized exception: %lu\n", ExceptionInfo->ExceptionRecord->ExceptionCode);
    break;
  }
    
  return EXCEPTION_EXECUTE_HANDLER;
}

EXPORTED void register_default_crash_handler() {
  SetUnhandledExceptionFilter(windows_crash_handler);
}

//==== stack trace

void setupStackFrame(CONTEXT *context, STACKFRAME *frame, DWORD* machineType) {
#ifdef _M_IX86    
    *machineType = IMAGE_FILE_MACHINE_I386;
    frame->AddrPC.Offset = context->Eip;
    frame->AddrPC.Mode = AddrModeFlat;
    frame->AddrFrame.Offset = context->Ebp;
    frame->AddrFrame.Mode = AddrModeFlat;
    frame->AddrStack.Offset = context->Esp;
    frame->AddrStack.Mode = AddrModeFlat;
#elif _M_X64
    *machineType = IMAGE_FILE_MACHINE_AMD64;
    frame->AddrPC.Offset = context->Rip;
    frame->AddrPC.Mode = AddrModeFlat;
    frame->AddrFrame.Offset = context->Rbp;
    frame->AddrFrame.Mode = AddrModeFlat;
    frame->AddrStack.Offset = context->Rsp;
    frame->AddrStack.Mode = AddrModeFlat;
#else
    #error "Platform not supported!"
#endif
}

char* prepareFuncName(char *rawName) {
    char* resName = (char*)malloc(MAX_SYM_NAME * sizeof(char)); 
//    resName[0] = '\0';
//    char* id = (char*)malloc((strlen(rawName) + 1) *sizeof(char));
//    size_t len = strlen(rawName);
//    char* funcName = (char*)malloc((len+1) *sizeof(char));
    
    strncpy_s(resName, MAX_SYM_NAME, rawName, strlen(rawName));
    
    return resName;
}

char* getFuncName(HANDLE curProcess, DWORD64 adr) {
    DWORD64 place = 0;
    char buf[sizeof(SYMBOL_INFO) + MAX_SYM_NAME * sizeof(TCHAR)];
    PSYMBOL_INFO pSymbol = (PSYMBOL_INFO)buf;
    
    pSymbol->SizeOfStruct = sizeof(SYMBOL_INFO);
    pSymbol->MaxNameLen = MAX_SYM_NAME;

    if (!SymFromAddr(curProcess, adr, &place, pSymbol)) {
        DWORD error = GetLastError();
        fprintf(stderr, "printStack: SymFromAddr error: %lu\n", error);
        return NULL;
    }
    return prepareFuncName(pSymbol->Name); 
}

void printStackFrame(char* name, int count, _Bool printReturnAddr, unsigned long long returnAddr) {
    fprintf(stderr, "%2d ", count);
    if (printReturnAddr) {
        fprintf(stderr, "(at 0x%p) ", (void*) returnAddr);
    }
    fprintf(stderr, "%s\n", name);
}

EXPORTED void  printStack(_Bool printReturnAddr, int maxFuncs) {
        // setup context
        CONTEXT context = {0};
        context.ContextFlags = CONTEXT_FULL;
        RtlCaptureContext(&context);
        
        SymSetOptions(!SYMOPT_UNDNAME); // all symbols are presented in undecorated form
        HANDLE curProcess = GetCurrentProcess();
        HANDLE curThread = GetCurrentThread();
        
        if (!SymInitialize(curProcess, NULL, TRUE)) {
            DWORD error = GetLastError();
            fprintf(stderr, "printStack: SymInitialize error: %lu\n", error);
            return;
        }
        
        STACKFRAME frame = {0};
        DWORD machineType = 0;
        DWORD count = 1;
        setupStackFrame(&context, &frame, &machineType);
        
        while (StackWalk(machineType, curProcess, curThread, 
            &frame, &context, 0, SymFunctionTableAccess, SymGetModuleBase, 0) ) {

            char* funcName = getFuncName(curProcess, frame.AddrPC.Offset);
            
            if (funcName == NULL) {
                break;
            }
            printStackFrame(funcName, count, printReturnAddr, frame.AddrReturn.Offset);
            
            if (count >= maxFuncs || strcmp(funcName, "main") == 0) {
                free(funcName);    
                break;
            }
            
            free(funcName);    
            count++;
        }
        SymCleanup(curProcess);
}
