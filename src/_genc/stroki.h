#ifndef _stroki_H
#define _stroki_H

#include "yunikod.h"
#include "utf8.h"

typedef struct stroki__TSimvoly *stroki__TSimvoly;
typedef struct stroki__TStroki *stroki__TStroki;
typedef struct stroki__TBayity *stroki__TBayity;
typedef struct stroki__TSimvoly {
TInt64 len; TInt64 capacity; TSymbol* body;
} *stroki__TSimvoly;
struct stroki__TRazborshchik_F {
struct stroki__TSborshchik* sb;
stroki__TSimvoly format;
TSymbol sim;
TInt64 N__sim;
TInt64 N__arg;
TBool oshibka;
TString soobshchenie;
TInt64 vid_formata;
TString imya_formata;
TInt64 chislo_dop_argumentov;
TInt64 razmer;
TInt64 tochnost_;
TBool znak;
TBool nuli;
TBool al_t;
TBool zaglavnye;
TBool veshch_eksponenta;
};
struct stroki__TRazborshchik_VT;
typedef struct stroki__TRazborshchik { struct stroki__TRazborshchik_VT* vtable; struct stroki__TRazborshchik_F f;} *stroki__TRazborshchik;

typedef struct stroki__TStroki {
TInt64 len; TInt64 capacity; TString* body;
} *stroki__TStroki;
typedef struct stroki__TBayity {
TInt64 len; TInt64 capacity; TByte* body;
} *stroki__TBayity;
struct stroki__TSborshchik_F {
stroki__TBayity bayity;
TInt64 chislo_simvolov;
};
struct stroki__TSborshchik_VT;
typedef struct stroki__TSborshchik { struct stroki__TSborshchik_VT* vtable; struct stroki__TSborshchik_F f;} *stroki__TSborshchik;

void stroki____s(TString s);
void stroki____tc(TString s, TInt64 tc);
void stroki__TSborshchik_f(stroki__TSborshchik sb, TString fs, TInt64 spisok_len, void* spisok);
void stroki__TRazborshchik_vzyat__simvol(stroki__TRazborshchik r);
TBool stroki__TRazborshchik_sleduyushchiyi_format(stroki__TRazborshchik r);
TBool stroki__TRazborshchik_razobrat__format(stroki__TRazborshchik r);
void stroki__TRazborshchik_razobrat__osnovu(stroki__TRazborshchik r);
TBool stroki__TRazborshchik_priznak(stroki__TRazborshchik r, TSymbol sim);
TInt64 stroki__TRazborshchik_chislo_ili_argument(stroki__TRazborshchik r);
void stroki__TRazborshchik_izvlech__str(stroki__TRazborshchik r);
void stroki__TRazborshchik_izvlech__sim(stroki__TRazborshchik r);
void stroki__TRazborshchik_izvlech__adr(stroki__TRazborshchik r);
void stroki__TRazborshchik_izvlech__tcel(stroki__TRazborshchik r);
void stroki__TRazborshchik_izvlech__shest(stroki__TRazborshchik r, TBool zaglavnye);
void stroki__TRazborshchik_izvlech__veshch(stroki__TRazborshchik r);
void stroki__TRazborshchik_razobrat__razmeshchenie(stroki__TRazborshchik r);
TInt64 stroki__TRazborshchik_razobrat__chislovoyi_argument(stroki__TRazborshchik r, TWord64 argT, TWord64 arg);
void stroki__TRazborshchik_dobavit__argument(stroki__TRazborshchik r, TWord64 argT, TWord64 arg);
void stroki__TRazborshchik_dobavit__po_umolchaniyu(stroki__TRazborshchik r, TWord64 argT, TWord64 arg);
void stroki__TRazborshchik_dobavit__tip(stroki__TRazborshchik r, TWord64 argT, TWord64 arg);
TString stroki__vstroennyyi_tip(TWord64 argT);
void stroki__TRazborshchik_dobavit__str(stroki__TRazborshchik r, TWord64 argT, TWord64 arg);
void stroki__TRazborshchik_dobavit__sim(stroki__TRazborshchik r, TWord64 argT, TWord64 arg);
void stroki__TRazborshchik_dobavit__adr(stroki__TRazborshchik r, TWord64 argT, TWord64 arg);
void stroki__TRazborshchik_dobavit__tcel(stroki__TRazborshchik r, TWord64 argT, TWord64 arg);
void stroki__TRazborshchik_dobavit__shest(stroki__TRazborshchik r, TWord64 argT, TWord64 arg);
void stroki__TRazborshchik_dobavit__veshch(stroki__TRazborshchik r, TWord64 argT, TWord64 arg);
void stroki__TRazborshchik_dobavit__stroku(stroki__TRazborshchik r, TString st);
void stroki__TRazborshchik_dobavit__simvol(stroki__TRazborshchik r, TSymbol sim);
void stroki__TRazborshchik_dobavit__tceloe10(stroki__TRazborshchik r, TInt64 tc);
void stroki__TRazborshchik_dobavit__chislo10(stroki__TRazborshchik r, TWord64 k, TBool neg);
void stroki__TRazborshchik_dobavit__znak(stroki__TRazborshchik r, stroki__TBayity b, TBool neg);
void stroki__TRazborshchik_dobavit__chislo16(stroki__TRazborshchik r, TWord64 k);
void stroki__TRazborshchik_dobavit__prefiks_chisla(stroki__TRazborshchik r, stroki__TBayity b, TSymbol sim);
void stroki__TRazborshchik_dobavit__bayity_reversom(stroki__TRazborshchik r, stroki__TBayity bayity, TInt64 N_);
void stroki__TRazborshchik_dobavit__veshchestvennoe(stroki__TRazborshchik r, TFloat64 v);
TInt64 stroki__indeks(TString s, TInt64 N__start, TString podstroka);
TInt64 stroki__indeks_bayita(TString s8, TInt64 N__start, TByte bayit);
TString stroki__f(TString format, TInt64 argumenty_len, void* argumenty);
TString stroki__soedinit_(TString razdelitel_, TInt64 stroki_len, void* stroki);
TString stroki__sobrat_(TInt64 stroki_len, void* stroki);
TString stroki__izvlech_(TString s, TInt64 N__bayita, TInt64 chislo_bayitov);
TString stroki__izvlech__iz_bayitov(stroki__TBayity bayity, TInt64 N__bayita, TInt64 chislo_bayitov);
stroki__TStroki stroki__razobrat_(TString s, TString razdelitel_);
stroki__TStroki stroki__razobrat_1(TString s, TByte razdelitel_);
TBool stroki__razdelit_(TString s, TString razdelitel_, TString* pervaya, TString* vtoraya);
TBool stroki__est__prefiks(TString s, TString prefiks);
TBool stroki__est__suffiks(TString s, TString suffiks);
TString stroki__obrezat__probel_nye_simvoly(TString s);
TBool stroki__stroka_v_tcel(TString s, TInt64* rez);
TBool stroki__stroka_v_slovo(TString s, TWord64* rez);
TBool stroki__slovo16(TString s8, TInt64 N_, TWord64* rez);
TBool stroki__slovo10(TString s8, TInt64 N_, TWord64* rez);
TBool stroki__tcifraQm(TSymbol sim);
TWord64 stroki__tcifra16(TSymbol sim);
TBool stroki__stroka_v_veshch(TString s, TFloat64* rez);
void stroki__TSborshchik_dobavit__stroku(stroki__TSborshchik sb, TString st);
TString stroki__TSborshchik_stroka(stroki__TSborshchik sb);
TInt64 stroki__TSborshchik_chislo_simvolov(stroki__TSborshchik sb);
void stroki__TSborshchik_dobavit__simvol(stroki__TSborshchik sb, TSymbol si);
TInt64 stroki__TSborshchik_chislo_bayitov(stroki__TSborshchik sb);
void stroki__TSborshchik_dobavit__bayity(stroki__TSborshchik sb, TString s8, TInt64 N_, TInt64 chislo_bayitov);
TString stroki__zamenit__vse(TString s, TString podstroka, TString zamena);
struct stroki__TRazborshchik_VT {
size_t self_size;
void (*__init__)(stroki__TRazborshchik);
void (*vzyat__simvol)(stroki__TRazborshchik);
TBool (*sleduyushchiyi_format)(stroki__TRazborshchik);
TBool (*razobrat__format)(stroki__TRazborshchik);
void (*razobrat__osnovu)(stroki__TRazborshchik);
TBool (*priznak)(stroki__TRazborshchik, TSymbol);
TInt64 (*chislo_ili_argument)(stroki__TRazborshchik);
void (*izvlech__str)(stroki__TRazborshchik);
void (*izvlech__sim)(stroki__TRazborshchik);
void (*izvlech__adr)(stroki__TRazborshchik);
void (*izvlech__tcel)(stroki__TRazborshchik);
void (*izvlech__shest)(stroki__TRazborshchik, TBool);
void (*izvlech__veshch)(stroki__TRazborshchik);
void (*razobrat__razmeshchenie)(stroki__TRazborshchik);
TInt64 (*razobrat__chislovoyi_argument)(stroki__TRazborshchik, TWord64, TWord64);
void (*dobavit__argument)(stroki__TRazborshchik, TWord64, TWord64);
void (*dobavit__po_umolchaniyu)(stroki__TRazborshchik, TWord64, TWord64);
void (*dobavit__tip)(stroki__TRazborshchik, TWord64, TWord64);
void (*dobavit__str)(stroki__TRazborshchik, TWord64, TWord64);
void (*dobavit__sim)(stroki__TRazborshchik, TWord64, TWord64);
void (*dobavit__adr)(stroki__TRazborshchik, TWord64, TWord64);
void (*dobavit__tcel)(stroki__TRazborshchik, TWord64, TWord64);
void (*dobavit__shest)(stroki__TRazborshchik, TWord64, TWord64);
void (*dobavit__veshch)(stroki__TRazborshchik, TWord64, TWord64);
void (*dobavit__stroku)(stroki__TRazborshchik, TString);
void (*dobavit__simvol)(stroki__TRazborshchik, TSymbol);
void (*dobavit__tceloe10)(stroki__TRazborshchik, TInt64);
void (*dobavit__chislo10)(stroki__TRazborshchik, TWord64, TBool);
void (*dobavit__znak)(stroki__TRazborshchik, stroki__TBayity, TBool);
void (*dobavit__chislo16)(stroki__TRazborshchik, TWord64);
void (*dobavit__prefiks_chisla)(stroki__TRazborshchik, stroki__TBayity, TSymbol);
void (*dobavit__bayity_reversom)(stroki__TRazborshchik, stroki__TBayity, TInt64);
void (*dobavit__veshchestvennoe)(stroki__TRazborshchik, TFloat64);
};
extern void * stroki__TRazborshchik_class_info_ptr;
struct stroki__TSborshchik_VT {
size_t self_size;
void (*__init__)(stroki__TSborshchik);
void (*f)(stroki__TSborshchik, TString, TInt64, void*);
void (*dobavit__stroku)(stroki__TSborshchik, TString);
TString (*stroka)(stroki__TSborshchik);
TInt64 (*chislo_simvolov)(stroki__TSborshchik);
void (*dobavit__simvol)(stroki__TSborshchik, TSymbol);
TInt64 (*chislo_bayitov)(stroki__TSborshchik);
void (*dobavit__bayity)(stroki__TSborshchik, TString, TInt64, TInt64);
};
extern void * stroki__TSborshchik_class_info_ptr;
void stroki__init();
#endif