// WEL - basic definitions for runtime

#ifndef _rt_defs_H
#define _rt_defs_H

#ifdef __linux__
    #define EXPORTED
#elif defined(_WIN32) || defined(_WIN64)
    #define EXPORTED __declspec(dllexport)
#else
    #define EXPORTED
#endif

#endif
