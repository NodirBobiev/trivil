#include "rt_api.h"
#include "utf8.h"

const TSymbol utf8__Nedop_simvol = 0xFFFD;
const TSymbol utf8__MaksSimvol = 0x10FFFF;
const TSymbol utf8__surrogateMin = 0xD800;
const TSymbol utf8__surrogateMax = 0xDFFF;
TSymbol utf8__dekodirovat__simvol(utf8__TBayity bayity, TInt64* N_, TInt64* chislo_bayitov) {
(*chislo_bayitov) = 0;
if (!((*N_) < (bayity)->len)) {
return utf8__Nedop_simvol;
}
TWord64 tek = (TWord64)(bayity->body[tri_indexcheck((*N_), bayity->len,"стд::юникод/utf8/декодер.tri:20:23")]);
(*N_)++;
(*chislo_bayitov) = 1;
if (tek < 0x80) {
return tri_TWord64_to_TSymbol(tek);
}
if (!((tek >= 0xC2) && (tek <= 0xF4))) {
return utf8__Nedop_simvol;
}
if (tek < 0xE0) {
(*chislo_bayitov) = 2;
if (!((*N_) < (bayity)->len)) {
return utf8__Nedop_simvol;
}
TWord64 tek1 = (TWord64)(bayity->body[tri_indexcheck((*N_), bayity->len,"стд::юникод/utf8/декодер.tri:35:28")]);
(*N_)++;
if (!((tek1 & 0xC0) == 0x80)) {
return utf8__Nedop_simvol;
}
return tri_TWord64_to_TSymbol((((tek & 0x1F) << 6) | (tek1 & 0x3F)));
}
if (tek < 0xF0) {
(*chislo_bayitov) = 3;
if (!(((*N_) + 1) < (bayity)->len)) {
return utf8__Nedop_simvol;
}
TWord64 tek1 = (TWord64)(bayity->body[tri_indexcheck((*N_), bayity->len,"стд::юникод/utf8/декодер.tri:45:28")]);
TWord64 tek2 = (TWord64)(bayity->body[tri_indexcheck(((*N_) + 1), bayity->len,"стд::юникод/utf8/декодер.tri:46:29")]);
(*N_) = ((*N_) + 2);
if (!((tek1 & 0xC0) == 0x80)) {
return utf8__Nedop_simvol;
}
if (!((tek2 & 0xC0) == 0x80)) {
return utf8__Nedop_simvol;
}
if ((tek == 0xED) && (tek1 > 0x9F)) {
return utf8__Nedop_simvol;
}
TWord64 kod = ((((tek & 0xF) << 12) | ((tek1 & 0x3F) << 6)) | (tek2 & 0x3F));
if (!(kod >= 0x800)) {
return utf8__Nedop_simvol;
}
return tri_TWord64_to_TSymbol(kod);
}
(*chislo_bayitov) = 4;
if (!(((*N_) + 2) < (bayity)->len)) {
return utf8__Nedop_simvol;
}
TWord64 tek1 = (TWord64)(bayity->body[tri_indexcheck((*N_), bayity->len,"стд::юникод/utf8/декодер.tri:65:24")]);
TWord64 tek2 = (TWord64)(bayity->body[tri_indexcheck(((*N_) + 1), bayity->len,"стд::юникод/utf8/декодер.tri:66:25")]);
TWord64 tek3 = (TWord64)(bayity->body[tri_indexcheck(((*N_) + 2), bayity->len,"стд::юникод/utf8/декодер.tri:67:25")]);
(*N_) = ((*N_) + 3);
if (!((tek1 & 0xC0) == 0x80)) {
return utf8__Nedop_simvol;
}
if (!((tek2 & 0xC0) == 0x80)) {
return utf8__Nedop_simvol;
}
if (!((tek3 & 0xC0) == 0x80)) {
return utf8__Nedop_simvol;
}
if (tek == 0xF0) {
if (!(tek1 >= 0x90)) {
return utf8__Nedop_simvol;
}
}
else if (tek == 0xF4) {
if (!(tek1 <= 0x8F)) {
return utf8__Nedop_simvol;
}
}
TWord64 kod = (((((tek & 0x7) << 18) | ((tek1 & 0x3F) << 12)) | ((tek2 & 0x3F) << 6)) | (tek3 & 0x3F));
return tri_TWord64_to_TSymbol(kod);
}
void utf8__propustit__simvol(utf8__TBayity bayity, TInt64* N_, TInt64* chislo_bayitov) {
(*chislo_bayitov) = 0;
if (!((*N_) < (bayity)->len)) {
return;
}
TWord64 tek = (TWord64)(bayity->body[tri_indexcheck((*N_), bayity->len,"стд::юникод/utf8/декодер.tri:93:23")]);
(*N_)++;
(*chislo_bayitov) = 1;
if (tek < 0x80) {
return;
}
if (!((tek >= 0xC2) && (tek <= 0xF4))) {
return;
}
if (tek < 0xE0) {
(*chislo_bayitov) = 2;
if (!((*N_) < (bayity)->len)) {
return;
}
(*N_)++;
return;
}
if (tek < 0xF0) {
(*chislo_bayitov) = 3;
if (!(((*N_) + 1) < (bayity)->len)) {
return;
}
(*N_) = ((*N_) + 2);
return;
}
(*chislo_bayitov) = 4;
if (!(((*N_) + 2) < (bayity)->len)) {
return;
}
(*N_) = ((*N_) + 3);
}
TBool utf8__korrektnyyi_simvol(TSymbol s) {
if ((0x0 <= s) && (s < utf8__surrogateMin)) {
return true;
}
else if ((utf8__surrogateMax < s) && (s <= utf8__MaksSimvol)) {
return true;
}
return false;
}
TBool utf8__korrektnaya_stroka(TString s) {
TString s8 = s;
TInt64 N_ = 0;
TInt64 bayitov = 0;
while (N_ < (s8)->bytes) {
TSymbol sim = utf8__dekodirovat__simvol_stroka8(s8, &(N_), &(bayitov));
if (sim == utf8__Nedop_simvol) {
return false;
}
}
return true;
}
TSymbol utf8__dekodirovat__simvol_stroka8(TString bayity, TInt64* N_, TInt64* chislo_bayitov) {
(*chislo_bayitov) = 0;
if (!((*N_) < (bayity)->bytes)) {
return utf8__Nedop_simvol;
}
TWord64 tek = (TWord64)(bayity->body[tri_indexcheck((*N_), bayity->bytes,"стд::юникод/utf8/декодер-стр8.tri:10:23")]);
(*N_)++;
(*chislo_bayitov) = 1;
if (tek < 0x80) {
return tri_TWord64_to_TSymbol(tek);
}
if (!((tek >= 0xC2) && (tek <= 0xF4))) {
return utf8__Nedop_simvol;
}
if (tek < 0xE0) {
(*chislo_bayitov) = 2;
if (!((*N_) < (bayity)->bytes)) {
return utf8__Nedop_simvol;
}
TWord64 tek1 = (TWord64)(bayity->body[tri_indexcheck((*N_), bayity->bytes,"стд::юникод/utf8/декодер-стр8.tri:25:28")]);
(*N_)++;
if (!((tek1 & 0xC0) == 0x80)) {
return utf8__Nedop_simvol;
}
return tri_TWord64_to_TSymbol((((tek & 0x1F) << 6) | (tek1 & 0x3F)));
}
if (tek < 0xF0) {
(*chislo_bayitov) = 3;
if (!(((*N_) + 1) < (bayity)->bytes)) {
return utf8__Nedop_simvol;
}
TWord64 tek1 = (TWord64)(bayity->body[tri_indexcheck((*N_), bayity->bytes,"стд::юникод/utf8/декодер-стр8.tri:35:28")]);
TWord64 tek2 = (TWord64)(bayity->body[tri_indexcheck(((*N_) + 1), bayity->bytes,"стд::юникод/utf8/декодер-стр8.tri:36:29")]);
(*N_) = ((*N_) + 2);
if (!((tek1 & 0xC0) == 0x80)) {
return utf8__Nedop_simvol;
}
if (!((tek2 & 0xC0) == 0x80)) {
return utf8__Nedop_simvol;
}
if ((tek == 0xED) && (tek1 > 0x9F)) {
return utf8__Nedop_simvol;
}
TWord64 kod = ((((tek & 0xF) << 12) | ((tek1 & 0x3F) << 6)) | (tek2 & 0x3F));
if (!(kod >= 0x800)) {
return utf8__Nedop_simvol;
}
return tri_TWord64_to_TSymbol(kod);
}
(*chislo_bayitov) = 4;
if (!(((*N_) + 2) < (bayity)->bytes)) {
return utf8__Nedop_simvol;
}
TWord64 tek1 = (TWord64)(bayity->body[tri_indexcheck((*N_), bayity->bytes,"стд::юникод/utf8/декодер-стр8.tri:55:24")]);
TWord64 tek2 = (TWord64)(bayity->body[tri_indexcheck(((*N_) + 1), bayity->bytes,"стд::юникод/utf8/декодер-стр8.tri:56:25")]);
TWord64 tek3 = (TWord64)(bayity->body[tri_indexcheck(((*N_) + 2), bayity->bytes,"стд::юникод/utf8/декодер-стр8.tri:57:25")]);
(*N_) = ((*N_) + 3);
if (!((tek1 & 0xC0) == 0x80)) {
return utf8__Nedop_simvol;
}
if (!((tek2 & 0xC0) == 0x80)) {
return utf8__Nedop_simvol;
}
if (!((tek3 & 0xC0) == 0x80)) {
return utf8__Nedop_simvol;
}
if (tek == 0xF0) {
if (!(tek1 >= 0x90)) {
return utf8__Nedop_simvol;
}
}
else if (tek == 0xF4) {
if (!(tek1 <= 0x8F)) {
return utf8__Nedop_simvol;
}
}
TWord64 kod = (((((tek & 0x7) << 18) | ((tek1 & 0x3F) << 12)) | ((tek2 & 0x3F) << 6)) | (tek3 & 0x3F));
return tri_TWord64_to_TSymbol(kod);
}
void utf8__propustit__simvol_stroka8(TString bayity, TInt64* N_, TInt64* chislo_bayitov) {
(*chislo_bayitov) = 0;
if (!((*N_) < (bayity)->bytes)) {
return;
}
TWord64 tek = (TWord64)(bayity->body[tri_indexcheck((*N_), bayity->bytes,"стд::юникод/utf8/декодер-стр8.tri:83:23")]);
(*N_)++;
(*chislo_bayitov) = 1;
if (tek < 0x80) {
return;
}
if (!((tek >= 0xC2) && (tek <= 0xF4))) {
return;
}
if (tek < 0xE0) {
(*chislo_bayitov) = 2;
if (!((*N_) < (bayity)->bytes)) {
return;
}
(*N_)++;
return;
}
if (tek < 0xF0) {
(*chislo_bayitov) = 3;
if (!(((*N_) + 1) < (bayity)->bytes)) {
return;
}
(*N_) = ((*N_) + 2);
return;
}
(*chislo_bayitov) = 4;
if (!(((*N_) + 2) < (bayity)->bytes)) {
return;
}
(*N_) = ((*N_) + 3);
}
static TBool init_done = false;
void utf8__init() {
if (init_done) return;
init_done = true;
}