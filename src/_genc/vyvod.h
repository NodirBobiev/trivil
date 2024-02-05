#ifndef _vyvod_H
#define _vyvod_H

#include "stroki.h"

void vyvod__simvol(TSymbol s);
void vyvod__f(TString format, TInt64 spisok_len, void* spisok);
void vyvod__init();
#endif