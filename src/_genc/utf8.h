#ifndef _utf8_H
#define _utf8_H


typedef struct utf8__TBayity *utf8__TBayity;
typedef struct utf8__TBayity {
TInt64 len; TInt64 capacity; TByte* body;
} *utf8__TBayity;
extern const TSymbol utf8__Nedop_simvol;
TSymbol utf8__dekodirovat__simvol(utf8__TBayity bayity, TInt64* N_, TInt64* chislo_bayitov);
void utf8__propustit__simvol(utf8__TBayity bayity, TInt64* N_, TInt64* chislo_bayitov);
TBool utf8__korrektnyyi_simvol(TSymbol s);
TBool utf8__korrektnaya_stroka(TString s);
TSymbol utf8__dekodirovat__simvol_stroka8(TString bayity, TInt64* N_, TInt64* chislo_bayitov);
void utf8__propustit__simvol_stroka8(TString bayity, TInt64* N_, TInt64* chislo_bayitov);
void utf8__init();
#endif