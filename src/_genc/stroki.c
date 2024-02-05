#include "rt_api.h"
#include "stroki.h"
#include "rt_sysapi.h"

//--- globals
static TString strlit1 = NULL;
static TString strlit2 = NULL;
static TString strlit3 = NULL;
static TString strlit5 = NULL;
static TString strlit6 = NULL;
static TString strlit8 = NULL;
static TString strlit13 = NULL;
static TString strlit17 = NULL;
static TString strlit18 = NULL;
static TString strlit19 = NULL;
static TString strlit20 = NULL;
static TString strlit21 = NULL;
static TString strlit22 = NULL;
static TString strlit23 = NULL;
static TString strlit24 = NULL;
static TString strlit25 = NULL;
static TString strlit26 = NULL;
static TString strlit28 = NULL;
static TString strlit31 = NULL;
static TString strlit33 = NULL;
static TString strlit34 = NULL;
static TString strlit35 = NULL;
static TString strlit37 = NULL;
static TString strlit38 = NULL;
static TString strlit40 = NULL;
static TString strlit41 = NULL;
static TString strlit42 = NULL;
static TString strlit43 = NULL;
static TString strlit44 = NULL;
static TString strlit45 = NULL;
static TString strlit46 = NULL;
static TString strlit47 = NULL;
static TString strlit48 = NULL;
static TString strlit51 = NULL;
static TString strlit54 = NULL;
static TString strlit55 = NULL;
static TString strlit58 = NULL;
static TString strlit61 = NULL;
static TString strlit63 = NULL;
static TString strlit65 = NULL;
static TString strlit83 = NULL;
static TString strlit85 = NULL;
//--- end globals

const TInt64 stroki___neop = 0;
const TInt64 stroki___umolchanie = 1;
const TInt64 stroki___tip = 2;
const TInt64 stroki___str = 3;
const TInt64 stroki___sim = 4;
const TInt64 stroki___adr = 5;
const TInt64 stroki___tcel = 6;
const TInt64 stroki___shest = 7;
const TInt64 stroki___veshch = 8;
const TSymbol stroki__KON_STR = 0x0;
const TByte stroki__PROBEL = 0x20;
const TInt64 stroki__maks_TCel64 = 9223372036854775807;
TWord64 stroki__maks_Slovo64;
void stroki____s(TString s) {
print_string(tri_newLiteralString(&strlit1, -1, 1, "!"));
print_string(s);
println();
}
void stroki____tc(TString s, TInt64 tc) {
print_string(tri_newLiteralString(&strlit2, -1, 1, "!"));
print_string(s);
print_string(tri_newLiteralString(&strlit3, -1, 1, "="));
print_int64(tc);
println();
}
void stroki__TSborshchik_f(stroki__TSborshchik sb, TString fs, TInt64 spisok_len, void* spisok) {
stroki__TRazborshchik loc4 = tri_newObject(stroki__TRazborshchik_class_info_ptr);loc4->f.sb = sb; loc4->f.format = tri_TString_to_Symbols(fs);
stroki__TRazborshchik r = loc4;
r->vtable->vzyat__simvol((stroki__TRazborshchik)r);
while (r->vtable->sleduyushchiyi_format((stroki__TRazborshchik)r)) {
if (!(((r)->f.N__arg + (r)->f.chislo_dop_argumentov) < spisok_len)) {
(r)->f.oshibka = true;
(r)->f.soobshchenie = tri_newLiteralString(&strlit5, -1, 24, "не достаточно аргументов");
break;
}
if ((r)->f.chislo_dop_argumentov > 0) {
(r)->f.razmer = r->vtable->razobrat__chislovoyi_argument((stroki__TRazborshchik)r, ((TWord64*)(spisok))[tri_indexcheck(((r)->f.N__arg + 1), spisok_len, "стд::строки/формат.tri:111:65") << 1], ((TWord64*)(spisok))[(tri_indexcheck(((r)->f.N__arg + 1), spisok_len, "стд::строки/формат.tri:111:93") << 1)+1]);
if (!(!((r)->f.oshibka))) {
break;
}
}
r->vtable->dobavit__argument((stroki__TRazborshchik)r, ((TWord64*)(spisok))[tri_indexcheck((r)->f.N__arg, spisok_len, "стд::строки/формат.tri:115:39") << 1], ((TWord64*)(spisok))[(tri_indexcheck((r)->f.N__arg, spisok_len, "стд::строки/формат.tri:115:63") << 1)+1]);
(r)->f.N__arg = (((r)->f.N__arg + 1) + (r)->f.chislo_dop_argumentov);
}
if (!((r)->f.oshibka) && ((r)->f.N__arg < spisok_len)) {
(r)->f.oshibka = true;
(r)->f.soobshchenie = tri_newLiteralString(&strlit6, -1, 24, "слишком много аргументов");
}
if ((r)->f.oshibka) {
stroki__TSborshchik loc7 = (r)->f.sb;
loc7->vtable->dobavit__stroku((stroki__TSborshchik)loc7, tri_newLiteralString(&strlit8, -1, 8, "!формат:"));
stroki__TSborshchik loc9 = (r)->f.sb;
loc9->vtable->dobavit__stroku((stroki__TSborshchik)loc9, (r)->f.soobshchenie);
stroki__TSborshchik loc10 = (r)->f.sb;
loc10->vtable->dobavit__simvol((stroki__TSborshchik)loc10, 0x21);
}
}
void stroki__TRazborshchik_vzyat__simvol(stroki__TRazborshchik r) {
if (!((r)->f.N__sim < ((r)->f.format)->len)) {
(r)->f.sim = stroki__KON_STR;
return;
}
stroki__TSimvoly loc11 = (r)->f.format;
(r)->f.sim = loc11->body[tri_indexcheck((r)->f.N__sim, loc11->len,"стд::строки/формат.tri:137:25")];
(r)->f.N__sim++;
}
TBool stroki__TRazborshchik_sleduyushchiyi_format(stroki__TRazborshchik r) {
while (!((r)->f.oshibka) && ((r)->f.sim != stroki__KON_STR)) {
if (!((r)->f.sim != 0x0)) {
tri_crash((char *)"0x0 в форматной строке","стд::строки/формат.tri:144:39");
}
if ((r)->f.sim == 0x24) {
if (r->vtable->razobrat__format((stroki__TRazborshchik)r)) {
return true;
}
}
stroki__TSborshchik loc12 = (r)->f.sb;
loc12->vtable->dobavit__simvol((stroki__TSborshchik)loc12, (r)->f.sim);
r->vtable->vzyat__simvol((stroki__TRazborshchik)r);
}
return false;
}
TBool stroki__TRazborshchik_razobrat__format(stroki__TRazborshchik r) {
(r)->f.chislo_dop_argumentov = 0;
(r)->f.razmer = -(1);
(r)->f.tochnost_ = -(1);
(r)->f.znak = false;
(r)->f.nuli = false;
(r)->f.al_t = false;
(r)->f.zaglavnye = false;
(r)->f.veshch_eksponenta = false;
r->vtable->vzyat__simvol((stroki__TRazborshchik)r);
if ((r)->f.sim == 0x24) {
return false;
}
r->vtable->razobrat__osnovu((stroki__TRazborshchik)r);
if (!(!((r)->f.oshibka))) {
return false;
}
if ((r)->f.sim == 0x3B) {
r->vtable->vzyat__simvol((stroki__TRazborshchik)r);
return true;
}
if ((r)->f.sim == 0x3A) {
r->vtable->vzyat__simvol((stroki__TRazborshchik)r);
r->vtable->razobrat__razmeshchenie((stroki__TRazborshchik)r);
}
return !((r)->f.oshibka);
}
void stroki__TRazborshchik_razobrat__osnovu(stroki__TRazborshchik r) {
if (!((r)->f.sim != stroki__KON_STR)) {
(r)->f.soobshchenie = tri_newLiteralString(&strlit13, -1, 14, "формат оборван");
(r)->f.oshibka = true;
return;
}
stroki__TSimvoly loc14 = tri_newVector(sizeof(TSymbol), 0, 0);
stroki__TSimvoly simvoly = loc14;
TInt64 N_ = 0;
while (yunikod__bukvaQm((r)->f.sim)) {
TSymbol loc15[1] = {(r)->f.sim};
tri_vectorAppend(simvoly, sizeof(TSymbol), 1, loc15);
N_++;
r->vtable->vzyat__simvol((stroki__TRazborshchik)r);
}
TString imya = tri_Symbols_to_TString(simvoly);
(r)->f.imya_formata = imya;
TString loc16 = imya;
if (tri_equalStrings(loc16, tri_emptyString())) {
(r)->f.vid_formata = stroki___umolchanie;
(r)->f.imya_formata = tri_newLiteralString(&strlit17, -1, 2, "$;");
}
else if (tri_equalStrings(loc16, tri_newLiteralString(&strlit18, -1, 3, "тип"))) {
(r)->f.vid_formata = stroki___tip;
}
else if (tri_equalStrings(loc16, tri_newLiteralString(&strlit19, -1, 3, "стр"))) {
r->vtable->izvlech__str((stroki__TRazborshchik)r);
}
else if (tri_equalStrings(loc16, tri_newLiteralString(&strlit20, -1, 3, "сим"))) {
r->vtable->izvlech__sim((stroki__TRazborshchik)r);
}
else if (tri_equalStrings(loc16, tri_newLiteralString(&strlit21, -1, 3, "адр"))) {
r->vtable->izvlech__adr((stroki__TRazborshchik)r);
}
else if (tri_equalStrings(loc16, tri_newLiteralString(&strlit22, -1, 3, "цел"))) {
r->vtable->izvlech__tcel((stroki__TRazborshchik)r);
}
else if (tri_equalStrings(loc16, tri_newLiteralString(&strlit23, -1, 3, "вещ"))) {
r->vtable->izvlech__veshch((stroki__TRazborshchik)r);
}
else if (tri_equalStrings(loc16, tri_newLiteralString(&strlit24, -1, 1, "ш"))) {
r->vtable->izvlech__shest((stroki__TRazborshchik)r, false);
}
else if (tri_equalStrings(loc16, tri_newLiteralString(&strlit25, -1, 1, "Ш"))) {
r->vtable->izvlech__shest((stroki__TRazborshchik)r, true);
}
else {
(r)->f.soobshchenie = tri_newLiteralString(&strlit26, -1, 18, "неизвестный формат");
(r)->f.oshibka = false;
}
}
TBool stroki__TRazborshchik_priznak(stroki__TRazborshchik r, TSymbol sim) {
if ((r)->f.sim == sim) {
r->vtable->vzyat__simvol((stroki__TRazborshchik)r);
return true;
}
return false;
}
TInt64 stroki__TRazborshchik_chislo_ili_argument(stroki__TRazborshchik r) {
if ((r)->f.sim == 0x2A) {
r->vtable->vzyat__simvol((stroki__TRazborshchik)r);
(r)->f.chislo_dop_argumentov++;
return 0;
}
if (!(yunikod__tcifraQm((r)->f.sim))) {
return -(1);
}
TInt64 N_ = 0;
while (true) {
TInt64 tcifra = ((TInt64)((r)->f.sim) - 48);
N_ = ((N_ * 10) + tcifra);
r->vtable->vzyat__simvol((stroki__TRazborshchik)r);
if (!(yunikod__tcifraQm((r)->f.sim))) {
break;
}
}
return N_;
}
void stroki__TRazborshchik_izvlech__str(stroki__TRazborshchik r) {
(r)->f.vid_formata = stroki___str;
(r)->f.al_t = r->vtable->priznak((stroki__TRazborshchik)r, 0x23);
(r)->f.razmer = r->vtable->chislo_ili_argument((stroki__TRazborshchik)r);
}
void stroki__TRazborshchik_izvlech__sim(stroki__TRazborshchik r) {
(r)->f.vid_formata = stroki___sim;
(r)->f.al_t = r->vtable->priznak((stroki__TRazborshchik)r, 0x23);
}
void stroki__TRazborshchik_izvlech__adr(stroki__TRazborshchik r) {
(r)->f.vid_formata = stroki___adr;
(r)->f.al_t = true;
(r)->f.nuli = r->vtable->priznak((stroki__TRazborshchik)r, 0x30);
(r)->f.razmer = r->vtable->chislo_ili_argument((stroki__TRazborshchik)r);
}
void stroki__TRazborshchik_izvlech__tcel(stroki__TRazborshchik r) {
(r)->f.vid_formata = stroki___tcel;
(r)->f.znak = r->vtable->priznak((stroki__TRazborshchik)r, 0x2B);
(r)->f.nuli = r->vtable->priznak((stroki__TRazborshchik)r, 0x30);
(r)->f.razmer = r->vtable->chislo_ili_argument((stroki__TRazborshchik)r);
}
void stroki__TRazborshchik_izvlech__shest(stroki__TRazborshchik r, TBool zaglavnye) {
(r)->f.vid_formata = stroki___shest;
(r)->f.al_t = r->vtable->priznak((stroki__TRazborshchik)r, 0x23);
(r)->f.nuli = r->vtable->priznak((stroki__TRazborshchik)r, 0x30);
(r)->f.zaglavnye = zaglavnye;
(r)->f.razmer = r->vtable->chislo_ili_argument((stroki__TRazborshchik)r);
}
void stroki__TRazborshchik_izvlech__veshch(stroki__TRazborshchik r) {
(r)->f.vid_formata = stroki___veshch;
(r)->f.znak = r->vtable->priznak((stroki__TRazborshchik)r, 0x2B);
(r)->f.nuli = r->vtable->priznak((stroki__TRazborshchik)r, 0x30);
(r)->f.razmer = r->vtable->chislo_ili_argument((stroki__TRazborshchik)r);
if ((r)->f.sim == 0x2E) {
r->vtable->vzyat__simvol((stroki__TRazborshchik)r);
r->vtable->chislo_ili_argument((stroki__TRazborshchik)r);
}
if ((r)->f.sim == 0x65) {
(r)->f.veshch_eksponenta = true;
r->vtable->vzyat__simvol((stroki__TRazborshchik)r);
}
else if ((r)->f.sim == 0x45) {
(r)->f.veshch_eksponenta = true;
(r)->f.zaglavnye = true;
r->vtable->vzyat__simvol((stroki__TRazborshchik)r);
}
}
void stroki__TRazborshchik_razobrat__razmeshchenie(stroki__TRazborshchik r) {
tri_crash((char *)"не сделано","стд::строки/формат.tri:324:6");
}
TInt64 stroki__TRazborshchik_razobrat__chislovoyi_argument(stroki__TRazborshchik r, TWord64 argT, TWord64 arg) {
TWord64 loc27 = argT;
if (loc27 == tri_tagTInt64()) {
return ((TUnion64)arg).i;
}
else if (loc27 == tri_tagTWord64() || loc27 == tri_tagTByte()) {
return tri_TWord64_to_TInt64(arg);
}
else {
(r)->f.oshibka = true;
(r)->f.soobshchenie = tri_newLiteralString(&strlit28, -1, 50, "дополнительный аргумент для '*' не является числом");
return 0;
}
}
void stroki__TRazborshchik_dobavit__argument(stroki__TRazborshchik r, TWord64 argT, TWord64 arg) {
TInt64 loc29 = (r)->f.vid_formata;
if (loc29 == stroki___umolchanie) {
r->vtable->dobavit__po_umolchaniyu((stroki__TRazborshchik)r, argT, arg);
}
else if (loc29 == stroki___tip) {
r->vtable->dobavit__tip((stroki__TRazborshchik)r, argT, arg);
}
else if (loc29 == stroki___str) {
r->vtable->dobavit__str((stroki__TRazborshchik)r, argT, arg);
}
else if (loc29 == stroki___sim) {
r->vtable->dobavit__sim((stroki__TRazborshchik)r, argT, arg);
}
else if (loc29 == stroki___adr) {
r->vtable->dobavit__adr((stroki__TRazborshchik)r, argT, arg);
}
else if (loc29 == stroki___tcel) {
r->vtable->dobavit__tcel((stroki__TRazborshchik)r, argT, arg);
}
else if (loc29 == stroki___shest) {
r->vtable->dobavit__shest((stroki__TRazborshchik)r, argT, arg);
}
else if (loc29 == stroki___veshch) {
r->vtable->dobavit__veshch((stroki__TRazborshchik)r, argT, arg);
}
else {
TString loc30[2] ={tri_newLiteralString(&strlit31, -1, 18, "формат не сделан: "), (r)->f.imya_formata};
tri_crash((char *)stroki__soedinit_(tri_emptyString(), 2, &loc30)->body,"стд::строки/формат.tri:354:9");
}
}
void stroki__TRazborshchik_dobavit__po_umolchaniyu(stroki__TRazborshchik r, TWord64 argT, TWord64 arg) {
TWord64 loc32 = argT;
if (loc32 == tri_tagTInt64()) {
r->vtable->dobavit__tceloe10((stroki__TRazborshchik)r, ((TUnion64)arg).i);
}
else if (loc32 == tri_tagTWord64()) {
r->vtable->dobavit__chislo10((stroki__TRazborshchik)r, arg, false);
}
else if (loc32 == tri_tagTSymbol() || loc32 == tri_tagTByte()) {
(r)->f.al_t = true;
r->vtable->dobavit__chislo16((stroki__TRazborshchik)r, arg);
}
else if (loc32 == tri_tagTFloat64()) {
r->vtable->dobavit__veshchestvennoe((stroki__TRazborshchik)r, ((TUnion64)arg).f);
}
else if (loc32 == tri_tagTBool()) {
if (arg == 0x0) {
r->vtable->dobavit__stroku((stroki__TRazborshchik)r, tri_newLiteralString(&strlit33, -1, 4, "ложь"));
}
else {
r->vtable->dobavit__stroku((stroki__TRazborshchik)r, tri_newLiteralString(&strlit34, -1, 6, "истина"));
}
}
else if (loc32 == tri_tagTString()) {
r->vtable->dobavit__stroku((stroki__TRazborshchik)r, (TString)((TUnion64)arg).a);
}
else if (loc32 == tri_tagTNull()) {
r->vtable->dobavit__stroku((stroki__TRazborshchik)r, tri_newLiteralString(&strlit35, -1, 5, "пусто"));
}
else {
stroki__TSborshchik loc36 = (r)->f.sb;
loc36->vtable->dobavit__stroku((stroki__TSborshchik)loc36, tri_newLiteralString(&strlit37, -1, 45, " *формат по умолчанию ($;) не сделано: тег 0x"));
r->vtable->dobavit__chislo16((stroki__TRazborshchik)r, argT);
}
}
void stroki__TRazborshchik_dobavit__tip(stroki__TRazborshchik r, TWord64 argT, TWord64 arg) {
if (tri_isClassTag(argT)) {
r->vtable->dobavit__stroku((stroki__TRazborshchik)r, tri_newLiteralString(&strlit38, -1, 6, "класс "));
r->vtable->dobavit__stroku((stroki__TRazborshchik)r, tri_className(argT));
return;
}
TString imya = stroki__vstroennyyi_tip(argT);
if (!tri_equalStrings(imya, tri_emptyString())) {
r->vtable->dobavit__stroku((stroki__TRazborshchik)r, imya);
return;
}
(r)->f.oshibka = true;
TString loc39[2] ={tri_newLiteralString(&strlit40, -1, 19, "'тип' не реализован"), (r)->f.imya_formata};
(r)->f.soobshchenie = stroki__soedinit_(tri_emptyString(), 2, &loc39);
}
TString stroki__vstroennyyi_tip(TWord64 argT) {
if (argT == tri_tagTByte()) {
return tri_newLiteralString(&strlit41, -1, 4, "Байт");
}
else if (argT == tri_tagTInt64()) {
return tri_newLiteralString(&strlit42, -1, 5, "Цел64");
}
else if (argT == tri_tagTWord64()) {
return tri_newLiteralString(&strlit43, -1, 7, "Слово64");
}
else if (argT == tri_tagTFloat64()) {
return tri_newLiteralString(&strlit44, -1, 5, "Вещ64");
}
else if (argT == tri_tagTBool()) {
return tri_newLiteralString(&strlit45, -1, 3, "Лог");
}
else if (argT == tri_tagTSymbol()) {
return tri_newLiteralString(&strlit46, -1, 6, "Символ");
}
else if (argT == tri_tagTString()) {
return tri_newLiteralString(&strlit47, -1, 6, "Строка");
}
else if (argT == tri_tagTNull()) {
return tri_newLiteralString(&strlit48, -1, 5, "Пусто");
}
return tri_emptyString();
}
void stroki__TRazborshchik_dobavit__str(stroki__TRazborshchik r, TWord64 argT, TWord64 arg) {
TWord64 loc49 = argT;
if (loc49 == tri_tagTString()) {
r->vtable->dobavit__stroku((stroki__TRazborshchik)r, (TString)((TUnion64)arg).a);
}
else {
(r)->f.oshibka = true;
TString loc50[2] ={tri_newLiteralString(&strlit51, -1, 25, "неверный тип для формата "), (r)->f.imya_formata};
(r)->f.soobshchenie = stroki__soedinit_(tri_emptyString(), 2, &loc50);
}
}
void stroki__TRazborshchik_dobavit__sim(stroki__TRazborshchik r, TWord64 argT, TWord64 arg) {
TWord64 loc52 = argT;
if (loc52 == tri_tagTSymbol()) {
r->vtable->dobavit__simvol((stroki__TRazborshchik)r, tri_TWord64_to_TSymbol(arg));
}
else {
(r)->f.oshibka = true;
TString loc53[2] ={tri_newLiteralString(&strlit54, -1, 25, "неверный тип для формата "), (r)->f.imya_formata};
(r)->f.soobshchenie = stroki__soedinit_(tri_emptyString(), 2, &loc53);
}
}
void stroki__TRazborshchik_dobavit__adr(stroki__TRazborshchik r, TWord64 argT, TWord64 arg) {
if (argT == tri_tagTString()) {
r->vtable->dobavit__chislo16((stroki__TRazborshchik)r, arg);
}
else if (tri_isClassTag(argT)) {
r->vtable->dobavit__chislo16((stroki__TRazborshchik)r, arg);
}
else {
(r)->f.oshibka = true;
(r)->f.soobshchenie = tri_newLiteralString(&strlit55, -1, 28, "неверный тип для формата адр");
}
}
void stroki__TRazborshchik_dobavit__tcel(stroki__TRazborshchik r, TWord64 argT, TWord64 arg) {
TWord64 loc56 = argT;
if (loc56 == tri_tagTInt64()) {
r->vtable->dobavit__tceloe10((stroki__TRazborshchik)r, ((TUnion64)arg).i);
}
else if (loc56 == tri_tagTWord64()) {
r->vtable->dobavit__chislo10((stroki__TRazborshchik)r, arg, false);
}
else if (loc56 == tri_tagTSymbol() || loc56 == tri_tagTByte()) {
r->vtable->dobavit__chislo10((stroki__TRazborshchik)r, arg, false);
}
else {
(r)->f.oshibka = true;
TString loc57[2] ={tri_newLiteralString(&strlit58, -1, 25, "неверный тип для формата "), (r)->f.imya_formata};
(r)->f.soobshchenie = stroki__soedinit_(tri_emptyString(), 2, &loc57);
}
}
void stroki__TRazborshchik_dobavit__shest(stroki__TRazborshchik r, TWord64 argT, TWord64 arg) {
TWord64 loc59 = argT;
if (loc59 == tri_tagTInt64() || loc59 == tri_tagTWord64() || loc59 == tri_tagTSymbol() || loc59 == tri_tagTByte() || loc59 == tri_tagTString()) {
r->vtable->dobavit__chislo16((stroki__TRazborshchik)r, arg);
}
else {
(r)->f.oshibka = true;
TString loc60[2] ={tri_newLiteralString(&strlit61, -1, 25, "неверный тип для формата "), (r)->f.imya_formata};
(r)->f.soobshchenie = stroki__soedinit_(tri_emptyString(), 2, &loc60);
}
}
void stroki__TRazborshchik_dobavit__veshch(stroki__TRazborshchik r, TWord64 argT, TWord64 arg) {
TWord64 loc62 = argT;
if (loc62 == tri_tagTFloat64()) {
r->vtable->dobavit__veshchestvennoe((stroki__TRazborshchik)r, ((TUnion64)arg).f);
}
else {
(r)->f.oshibka = true;
TWord64 loc64[6] = {tri_tagTString(), ((TUnion64)(void*)tri_newLiteralString(&strlit65, -1, 25, "неверный тип для формата ")).w, tri_tagTString(), ((TUnion64)(void*)(r)->f.imya_formata).w, tri_tagTWord64(), argT};
(r)->f.soobshchenie = stroki__f(tri_newLiteralString(&strlit63, -1, 10, "$; $;: $ш;"), 3, &loc64);
}
}
void stroki__TRazborshchik_dobavit__stroku(stroki__TRazborshchik r, TString st) {
if (!(!((r)->f.al_t))) {
tri_crash((char *)"не реализовано - # для строки ","стд::строки/формат.tri:508:25");
}
TInt64 ogranichenie = (r)->f.razmer;
if (ogranichenie == 0) {
return;
}
if ((ogranichenie > 0) && (ogranichenie < tri_lenString(st))) {
stroki__TBayity bayity = tri_TString_to_Bytes(st);
TInt64 N_ = 0;
TInt64 bayitov = 0;
TInt64 chislo_bayitov = 0;
while (ogranichenie > 0) {
utf8__propustit__simvol((utf8__TBayity)bayity, &(N_), &(bayitov));
chislo_bayitov = (chislo_bayitov + bayitov);
ogranichenie--;
}
TString podstroka = stroki__izvlech_(st, 0, chislo_bayitov);
stroki__TSborshchik loc66 = (r)->f.sb;
loc66->vtable->dobavit__stroku((stroki__TSborshchik)loc66, podstroka);
}
else {
stroki__TSborshchik loc67 = (r)->f.sb;
loc67->vtable->dobavit__stroku((stroki__TSborshchik)loc67, st);
}
}
void stroki__TRazborshchik_dobavit__simvol(stroki__TRazborshchik r, TSymbol sim) {
stroki__TSborshchik loc68 = (r)->f.sb;
loc68->vtable->dobavit__simvol((stroki__TSborshchik)loc68, sim);
}
void stroki__TRazborshchik_dobavit__tceloe10(stroki__TRazborshchik r, TInt64 tc) {
r->vtable->dobavit__chislo10((stroki__TRazborshchik)r, ((TUnion64)(TInt64)tc).w, (tc < 0));
}
void stroki__TRazborshchik_dobavit__chislo10(stroki__TRazborshchik r, TWord64 k, TBool neg) {
if (neg) {
k = -(k);
}
stroki__TBayity loc69 = tri_newVector(sizeof(TByte), 0, 21);
stroki__TBayity b = loc69;
while (true) {
TWord64 ost = (k % 0xA);
k = (k / 0xA);
TByte loc70[1] = {tri_TWord64_to_TByte((ost + 0x30))};
tri_vectorAppend(b, sizeof(TByte), 1, loc70);
if (k == 0x0) {
break;
}
}
TInt64 shirina = (r)->f.razmer;
if (shirina > (b)->len) {
shirina = (shirina - (b)->len);
if ((r)->f.nuli) {
shirina--;
while (shirina > 0) {
TByte loc71[1] = {0x30};
tri_vectorAppend(b, sizeof(TByte), 1, loc71);
shirina--;
}
r->vtable->dobavit__znak((stroki__TRazborshchik)r, b, neg);
}
else {
r->vtable->dobavit__znak((stroki__TRazborshchik)r, b, neg);
while (shirina > 0) {
TByte loc72[1] = {stroki__PROBEL};
tri_vectorAppend(b, sizeof(TByte), 1, loc72);
shirina--;
}
}
}
else {
r->vtable->dobavit__znak((stroki__TRazborshchik)r, b, neg);
}
r->vtable->dobavit__bayity_reversom((stroki__TRazborshchik)r, b, (b)->len);
}
void stroki__TRazborshchik_dobavit__znak(stroki__TRazborshchik r, stroki__TBayity b, TBool neg) {
if (neg) {
TByte loc73[1] = {0x2D};
tri_vectorAppend(b, sizeof(TByte), 1, loc73);
}
else if ((r)->f.znak) {
TByte loc74[1] = {0x2B};
tri_vectorAppend(b, sizeof(TByte), 1, loc74);
}
}
void stroki__TRazborshchik_dobavit__chislo16(stroki__TRazborshchik r, TWord64 k) {
stroki__TBayity loc75 = tri_newVector(sizeof(TByte), 0, 17);
stroki__TBayity b = loc75;
TByte sim = 0;
while (true) {
TWord64 ost = (k % 0x10);
k = (k / 0x10);
if (ost < 0xA) {
sim = tri_TWord64_to_TByte((ost + 0x30));
}
else {
sim = tri_TWord64_to_TByte(((ost - 0xA) + 0x41));
}
TByte loc76[1] = {sim};
tri_vectorAppend(b, sizeof(TByte), 1, loc76);
if (k == 0x0) {
break;
}
}
TInt64 shirina = (r)->f.razmer;
if (shirina > (b)->len) {
shirina = (shirina - (b)->len);
if ((r)->f.nuli) {
if ((r)->f.al_t) {
shirina = (shirina - 2);
}
while (shirina > 0) {
TByte loc77[1] = {0x30};
tri_vectorAppend(b, sizeof(TByte), 1, loc77);
shirina--;
}
r->vtable->dobavit__prefiks_chisla((stroki__TRazborshchik)r, b, 0x78);
}
else {
r->vtable->dobavit__prefiks_chisla((stroki__TRazborshchik)r, b, 0x78);
while (shirina > 0) {
TByte loc78[1] = {stroki__PROBEL};
tri_vectorAppend(b, sizeof(TByte), 1, loc78);
shirina--;
}
}
}
else {
r->vtable->dobavit__prefiks_chisla((stroki__TRazborshchik)r, b, 0x78);
}
r->vtable->dobavit__bayity_reversom((stroki__TRazborshchik)r, b, (b)->len);
}
void stroki__TRazborshchik_dobavit__prefiks_chisla(stroki__TRazborshchik r, stroki__TBayity b, TSymbol sim) {
if ((r)->f.al_t) {
TByte loc79[1] = {tri_TSymbol_to_TByte(sim)};
tri_vectorAppend(b, sizeof(TByte), 1, loc79);
TByte loc80[1] = {0x30};
tri_vectorAppend(b, sizeof(TByte), 1, loc80);
}
}
void stroki__TRazborshchik_dobavit__bayity_reversom(stroki__TRazborshchik r, stroki__TBayity bayity, TInt64 N_) {
((r)->f.sb)->f.chislo_simvolov = (((r)->f.sb)->f.chislo_simvolov + N_);
N_--;
while (N_ >= 0) {
TByte loc81[1] = {bayity->body[tri_indexcheck(N_, bayity->len,"стд::строки/формат.tri:637:35")]};
tri_vectorAppend(((r)->f.sb)->f.bayity, sizeof(TByte), 1, loc81);
N_--;
}
}
void stroki__TRazborshchik_dobavit__veshchestvennoe(stroki__TRazborshchik r, TFloat64 v) {
if ((r)->f.veshch_eksponenta) {
stroki__TSborshchik loc82 = (r)->f.sb;
loc82->vtable->dobavit__stroku((stroki__TSborshchik)loc82, sysapi_float64_to_string(tri_newLiteralString(&strlit83, -1, 2, "%e"), v));
}
else {
stroki__TSborshchik loc84 = (r)->f.sb;
loc84->vtable->dobavit__stroku((stroki__TSborshchik)loc84, sysapi_float64_to_string(tri_newLiteralString(&strlit85, -1, 2, "%g"), v));
}
}
TInt64 stroki__indeks(TString s, TInt64 N__start, TString podstroka) {
if (!(N__start >= 0)) {
tri_crash((char *)"№-старт < 0","стд::строки/индексы.tri:7:29");
}
TString s8 = s;
TString p8 = podstroka;
TInt64 dl = (p8)->bytes;
if (dl == 0) {
return N__start;
}
else if ((N__start + dl) > (s8)->bytes) {
return -(1);
}
else if (dl == 1) {
return stroki__indeks_bayita(s8, N__start, p8->body[tri_indexcheck(0, p8->bytes,"стд::строки/индексы.tri:17:56")]);
}
else if ((N__start == 0) && (dl == (s8)->bytes)) {
if (tri_equalStrings(s, podstroka)) {
return 0;
}
return -(1);
}
TInt64 stopor = ((s8)->bytes - dl);
TByte b0 = p8->body[tri_indexcheck(0, p8->bytes,"стд::строки/индексы.tri:26:19")];
while (N__start <= stopor) {
if (s8->body[tri_indexcheck(N__start, s8->bytes,"стд::строки/индексы.tri:29:17")] == b0) {
TInt64 N_ = 1;
while ((N_ < dl) && (s8->body[tri_indexcheck((N__start + N_), s8->bytes,"стд::строки/индексы.tri:31:38")] == p8->body[tri_indexcheck(N_, p8->bytes,"стд::строки/индексы.tri:31:48")])) {
N_++;
}
if (N_ == dl) {
return N__start;
}
}
N__start++;
}
return -(1);
}
TInt64 stroki__indeks_bayita(TString s8, TInt64 N__start, TByte bayit) {
if (!(N__start >= 0)) {
tri_crash((char *)"№-старт < 0","стд::строки/индексы.tri:44:29");
}
while (N__start < (s8)->bytes) {
if (s8->body[tri_indexcheck(N__start, s8->bytes,"стд::строки/индексы.tri:47:17")] == bayit) {
return N__start;
}
N__start++;
}
return -(1);
}
TString stroki__f(TString format, TInt64 argumenty_len, void* argumenty) {
stroki__TSborshchik loc86 = tri_newObject(stroki__TSborshchik_class_info_ptr);
stroki__TSborshchik sb = loc86;
sb->vtable->f((stroki__TSborshchik)sb, format, argumenty_len, argumenty);
return sb->vtable->stroka((stroki__TSborshchik)sb);
}
TString stroki__soedinit_(TString razdelitel_, TInt64 stroki_len, void* stroki) {
if (!(stroki_len > 0)) {
return tri_emptyString();
}
TInt64 razmer = 0;
TInt64 N_ = 0;
while (N_ < stroki_len) {
razmer = (razmer + (((TString*)stroki)[tri_indexcheck(N_, stroki_len, "стд::строки/строковые.tri:23:41")])->bytes);
N_++;
}
razmer = (razmer + ((stroki_len - 1) * (razdelitel_)->bytes));
stroki__TBayity loc88 = tri_newVector(sizeof(TByte), 0, razmer);
stroki__TSborshchik loc87 = tri_newObject(stroki__TSborshchik_class_info_ptr);loc87->f.bayity = loc88;
stroki__TSborshchik sb = loc87;
N_ = 0;
while (N_ < (stroki_len - 1)) {
sb->vtable->dobavit__stroku((stroki__TSborshchik)sb, ((TString*)stroki)[tri_indexcheck(N_, stroki_len, "стд::строки/строковые.tri:31:35")]);
sb->vtable->dobavit__stroku((stroki__TSborshchik)sb, razdelitel_);
N_++;
}
sb->vtable->dobavit__stroku((stroki__TSborshchik)sb, ((TString*)stroki)[tri_indexcheck(N_, stroki_len, "стд::строки/строковые.tri:35:31")]);
return sb->vtable->stroka((stroki__TSborshchik)sb);
}
TString stroki__sobrat_(TInt64 stroki_len, void* stroki) {
if (!(stroki_len > 0)) {
return tri_emptyString();
}
TInt64 razmer = 0;
TInt64 N_ = 0;
while (N_ < stroki_len) {
razmer = (razmer + (((TString*)stroki)[tri_indexcheck(N_, stroki_len, "стд::строки/строковые.tri:47:41")])->bytes);
N_++;
}
stroki__TBayity loc90 = tri_newVector(sizeof(TByte), 0, razmer);
stroki__TSborshchik loc89 = tri_newObject(stroki__TSborshchik_class_info_ptr);loc89->f.bayity = loc90;
stroki__TSborshchik sb = loc89;
N_ = 0;
while (N_ < stroki_len) {
sb->vtable->dobavit__stroku((stroki__TSborshchik)sb, ((TString*)stroki)[tri_indexcheck(N_, stroki_len, "стд::строки/строковые.tri:54:35")]);
N_++;
}
return sb->vtable->stroka((stroki__TSborshchik)sb);
}
TString stroki__izvlech_(TString s, TInt64 N__bayita, TInt64 chislo_bayitov) {
return tri_substring(s, N__bayita, chislo_bayitov);
}
TString stroki__izvlech__iz_bayitov(stroki__TBayity bayity, TInt64 N__bayita, TInt64 chislo_bayitov) {
return tri_substring_from_bytes(bayity, N__bayita, chislo_bayitov);
}
stroki__TStroki stroki__razobrat_(TString s, TString razdelitel_) {
TString r8 = razdelitel_;
switch ((r8)->bytes) {
case 0: 
{
stroki__TStroki loc91 = tri_newVector(sizeof(TString), 1, 0);loc91->body[0] = s;
return loc91;
}
break;
case 1: 
{
return stroki__razobrat_1(s, r8->body[tri_indexcheck(0, r8->bytes,"стд::строки/строковые.tri:80:34")]);
}
break;
default:{
tri_crash((char *)"не реализовано - длина разделителя > 1","стд::строки/строковые.tri:82:9");
}
}
stroki__TStroki loc92 = tri_newVector(sizeof(TString), 0, 0);
return loc92;
}
stroki__TStroki stroki__razobrat_1(TString s, TByte razdelitel_) {
stroki__TStroki loc93 = tri_newVector(sizeof(TString), 0, 0);
stroki__TStroki rez = loc93;
TString s8 = s;
TInt64 N_ = 0;
TInt64 N__razd = -(1);
while (N_ < (s8)->bytes) {
if (s8->body[tri_indexcheck(N_, s8->bytes,"стд::строки/строковые.tri:94:17")] == razdelitel_) {
if ((N__razd + 1) == N_) {
TString loc94[1] = {tri_emptyString()};
tri_vectorAppend(rez, sizeof(TString), 1, loc94);
}
else {
TString loc95[1] = {tri_substring(s8, (N__razd + 1), ((N_ - N__razd) - 1))};
tri_vectorAppend(rez, sizeof(TString), 1, loc95);
}
N__razd = N_;
}
N_++;
}
TString loc96[1] = {tri_substring(s8, (N__razd + 1), (((s8)->bytes - N__razd) - 1))};
tri_vectorAppend(rez, sizeof(TString), 1, loc96);
return rez;
}
TBool stroki__razdelit_(TString s, TString razdelitel_, TString* pervaya, TString* vtoraya) {
TInt64 i = stroki__indeks(s, 0, razdelitel_);
if (!(i >= 0)) {
(*pervaya) = s;
(*vtoraya) = tri_emptyString();
return false;
}
TString s8 = s;
TString r8 = razdelitel_;
(*pervaya) = tri_substring(s8, 0, i);
(*vtoraya) = tri_substring(s8, (i + (r8)->bytes), ((s8)->bytes - (i + (r8)->bytes)));
return true;
}
TBool stroki__est__prefiks(TString s, TString prefiks) {
TString s8 = s;
TString p8 = prefiks;
if (!((p8)->bytes <= (s8)->bytes)) {
return false;
}
TInt64 N_ = 0;
while (N_ < (p8)->bytes) {
if (s8->body[tri_indexcheck(N_, s8->bytes,"стд::строки/строковые.tri:140:17")] != p8->body[tri_indexcheck(N_, p8->bytes,"стд::строки/строковые.tri:140:25")]) {
return false;
}
N_++;
}
return true;
}
TBool stroki__est__suffiks(TString s, TString suffiks) {
TString s8 = s;
TString k8 = suffiks;
if (!((k8)->bytes <= (s8)->bytes)) {
return false;
}
TInt64 N_1 = ((k8)->bytes - 1);
TInt64 N_2 = ((s8)->bytes - 1);
while (N_1 >= 0) {
if (s8->body[tri_indexcheck(N_2, s8->bytes,"стд::строки/строковые.tri:156:17")] != k8->body[tri_indexcheck(N_1, k8->bytes,"стд::строки/строковые.tri:156:26")]) {
return false;
}
N_1--;
N_2--;
}
return true;
}
TString stroki__obrezat__probel_nye_simvoly(TString s) {
TString s8 = s;
TInt64 N_1 = 0;
while (N_1 < (s8)->bytes) {
if (!(s8->body[tri_indexcheck(N_1, s8->bytes,"стд::строки/строковые.tri:172:17")] < 0x80)) {
break;
}
if (!(yunikod__probel_nyyi_simvolQm((TSymbol)(s8->body[tri_indexcheck(N_1, s8->bytes,"стд::строки/строковые.tri:173:43")])))) {
break;
}
N_1++;
}
TInt64 N_2 = (s8)->bytes;
while (N_2 > N_1) {
TByte bayit = s8->body[tri_indexcheck((N_2 - 1), s8->bytes,"стд::строки/строковые.tri:179:27")];
if (!(bayit < 0x80)) {
break;
}
if (!(yunikod__probel_nyyi_simvolQm((TSymbol)(bayit)))) {
break;
}
N_2--;
}
return tri_substring(s8, N_1, (N_2 - N_1));
}
TBool stroki__stroka_v_tcel(TString s, TInt64* rez) {
(*rez) = 0;
TString s8 = s;
if (!((s8)->bytes > 0)) {
return false;
}
TInt64 N_ = 0;
TBool neg = (s8->body[tri_indexcheck(0, s8->bytes,"стд::строки/числа.tri:17:20")] == 0x2D);
if (neg || (s8->body[tri_indexcheck(0, s8->bytes,"стд::строки/числа.tri:18:19")] == 0x2B)) {
N_++;
if (!(N_ < (s8)->bytes)) {
return false;
}
}
TSymbol sim = (TSymbol)(s8->body[tri_indexcheck(N_, s8->bytes,"стд::строки/числа.tri:23:21")]);
if (!(stroki__tcifraQm(sim))) {
return false;
}
TInt64 chislo = 0;
while (true) {
TInt64 tcifra = ((TInt64)(sim) - 48);
if (!(chislo <= ((stroki__maks_TCel64 - tcifra) / 10))) {
return false;
}
chislo = ((chislo * 10) + tcifra);
N_++;
if (N_ >= (s8)->bytes) {
if (neg) {
chislo = -(chislo);
}
(*rez) = chislo;
return true;
}
sim = (TSymbol)(s8->body[tri_indexcheck(N_, s8->bytes,"стд::строки/числа.tri:42:19")]);
if (!(stroki__tcifraQm(sim))) {
break;
}
}
return false;
}
TBool stroki__stroka_v_slovo(TString s, TWord64* rez) {
(*rez) = 0;
TString s8 = s;
if (!((s8)->bytes > 0)) {
return false;
}
if (s8->body[tri_indexcheck(0, s8->bytes,"стд::строки/числа.tri:57:13")] == 0x30) {
if (!((s8)->bytes > 1)) {
return true;
}
TByte sled = s8->body[tri_indexcheck(1, s8->bytes,"стд::строки/числа.tri:60:25")];
if ((sled == 0x78) || (sled == 0x58)) {
if (!((s8)->bytes > 2)) {
return false;
}
return stroki__slovo16(s8, 2, &((*rez)));
}
}
return stroki__slovo10(s8, 0, &((*rez)));
}
TBool stroki__slovo16(TString s8, TInt64 N_, TWord64* rez) {
TSymbol sim = (TSymbol)(s8->body[tri_indexcheck(N_, s8->bytes,"стд::строки/числа.tri:76:21")]);
TWord64 tcifra = stroki__tcifra16(sim);
if (!(tcifra < 0x10)) {
return false;
}
TWord64 chislo = 0;
while (true) {
if (!(chislo <= ((stroki__maks_Slovo64 - tcifra) / 0x10))) {
return false;
}
chislo = ((chislo * 0x10) + tcifra);
N_++;
if (N_ >= (s8)->bytes) {
(*rez) = chislo;
return true;
}
sim = (TSymbol)(s8->body[tri_indexcheck(N_, s8->bytes,"стд::строки/числа.tri:93:19")]);
tcifra = stroki__tcifra16(sim);
if (!(tcifra < 0x10)) {
return false;
}
}
return false;
}
TBool stroki__slovo10(TString s8, TInt64 N_, TWord64* rez) {
TSymbol sim = (TSymbol)(s8->body[tri_indexcheck(N_, s8->bytes,"стд::строки/числа.tri:105:21")]);
if (!(stroki__tcifraQm(sim))) {
return false;
}
TWord64 chislo = 0x0;
while (true) {
TWord64 tcifra = ((TWord64)(sim) - 0x30);
if (!(chislo <= ((stroki__maks_Slovo64 - tcifra) / 0xA))) {
return false;
}
chislo = ((chislo * 0xA) + tcifra);
N_++;
if (N_ >= (s8)->bytes) {
(*rez) = chislo;
return true;
}
sim = (TSymbol)(s8->body[tri_indexcheck(N_, s8->bytes,"стд::строки/числа.tri:121:19")]);
if (!(stroki__tcifraQm(sim))) {
break;
}
}
return false;
}
TBool stroki__tcifraQm(TSymbol sim) {
return ((sim >= 0x30) && (sim <= 0x39));
}
TWord64 stroki__tcifra16(TSymbol sim) {
if ((sim >= 0x30) && (sim <= 0x39)) {
return ((TWord64)(sim) - 0x30);
}
else if ((sim >= 0x41) && (sim <= 0x46)) {
return ((0xA + (TWord64)(sim)) - 0x41);
}
else if ((sim >= 0x61) && (sim <= 0x66)) {
return ((0xA + (TWord64)(sim)) - 0x61);
}
else {
return 16;
}
}
TBool stroki__stroka_v_veshch(TString s, TFloat64* rez) {
return sysapi_string_to_float64(s, &((*rez)));
}
void stroki__TSborshchik_dobavit__stroku(stroki__TSborshchik sb, TString st) {
TString st8 = st;
(sb)->f.chislo_simvolov = ((sb)->f.chislo_simvolov + (st8)->bytes);
TString loc97 = st8;
tri_vectorAppend((sb)->f.bayity, sizeof(TByte), loc97->bytes, loc97->body);
}
TString stroki__TSborshchik_stroka(stroki__TSborshchik sb) {
return tri_Bytes_to_TString((sb)->f.bayity);
}
TInt64 stroki__TSborshchik_chislo_simvolov(stroki__TSborshchik sb) {
return (sb)->f.chislo_simvolov;
}
void stroki__TSborshchik_dobavit__simvol(stroki__TSborshchik sb, TSymbol si) {
tri_vectorAppend_TSymbol_to_Bytes((sb)->f.bayity, si);
(sb)->f.chislo_simvolov++;
}
TInt64 stroki__TSborshchik_chislo_bayitov(stroki__TSborshchik sb) {
return ((sb)->f.bayity)->len;
}
void stroki__TSborshchik_dobavit__bayity(stroki__TSborshchik sb, TString s8, TInt64 N_, TInt64 chislo_bayitov) {
TInt64 stopor = (N_ + chislo_bayitov);
while (N_ < stopor) {
TByte loc98[1] = {s8->body[tri_indexcheck(N_, s8->bytes,"стд::строки/сборщик.tri:41:30")]};
tri_vectorAppend((sb)->f.bayity, sizeof(TByte), 1, loc98);
N_++;
}
}
TString stroki__zamenit__vse(TString s, TString podstroka, TString zamena) {
if (!(tri_lenString(podstroka) > 0)) {
return s;
}
TString s8 = s;
TString p8 = podstroka;
TInt64 N_ = stroki__indeks(s, 0, podstroka);
if (!(N_ >= 0)) {
return s;
}
stroki__TSborshchik loc99 = tri_newObject(stroki__TSborshchik_class_info_ptr);
stroki__TSborshchik sb = loc99;
TInt64 N__start = 0;
while (true) {
sb->vtable->dobavit__bayity((stroki__TSborshchik)sb, s8, N__start, (N_ - N__start));
sb->vtable->dobavit__stroku((stroki__TSborshchik)sb, zamena);
N__start = (N_ + (p8)->bytes);
N_ = stroki__indeks(s, N__start, podstroka);
if (!(N_ >= 0)) {
break;
}
}
sb->vtable->dobavit__bayity((stroki__TSborshchik)sb, s8, N__start, ((s8)->bytes - N__start));
return sb->vtable->stroka((stroki__TSborshchik)sb);
}
struct stroki__TRazborshchik_Meta {
size_t object_size;
void* base;
TString name;
};
void stroki__TRazborshchik__init__(stroki__TRazborshchik o) {
o->f.sim = stroki__KON_STR;
o->f.N__sim = 0;
o->f.N__arg = 0;
o->f.oshibka = false;
o->f.soobshchenie = tri_emptyString();
o->f.vid_formata = stroki___neop;
o->f.imya_formata = tri_emptyString();
o->f.chislo_dop_argumentov = 0;
o->f.razmer = -(1);
o->f.tochnost_ = -(1);
o->f.znak = false;
o->f.nuli = false;
o->f.al_t = false;
o->f.zaglavnye = false;
o->f.veshch_eksponenta = false;
}
struct { struct stroki__TRazborshchik_VT vt; struct stroki__TRazborshchik_Meta meta; } stroki__TRazborshchik_class_info;
void * stroki__TRazborshchik_class_info_ptr;
void stroki__TRazborshchik_init() {
stroki__TRazborshchik_class_info.vt.self_size = sizeof(struct stroki__TRazborshchik_VT);
stroki__TRazborshchik_class_info.vt.__init__ = &stroki__TRazborshchik__init__;
stroki__TRazborshchik_class_info.vt.vzyat__simvol = &stroki__TRazborshchik_vzyat__simvol;
stroki__TRazborshchik_class_info.vt.sleduyushchiyi_format = &stroki__TRazborshchik_sleduyushchiyi_format;
stroki__TRazborshchik_class_info.vt.razobrat__format = &stroki__TRazborshchik_razobrat__format;
stroki__TRazborshchik_class_info.vt.razobrat__osnovu = &stroki__TRazborshchik_razobrat__osnovu;
stroki__TRazborshchik_class_info.vt.priznak = &stroki__TRazborshchik_priznak;
stroki__TRazborshchik_class_info.vt.chislo_ili_argument = &stroki__TRazborshchik_chislo_ili_argument;
stroki__TRazborshchik_class_info.vt.izvlech__str = &stroki__TRazborshchik_izvlech__str;
stroki__TRazborshchik_class_info.vt.izvlech__sim = &stroki__TRazborshchik_izvlech__sim;
stroki__TRazborshchik_class_info.vt.izvlech__adr = &stroki__TRazborshchik_izvlech__adr;
stroki__TRazborshchik_class_info.vt.izvlech__tcel = &stroki__TRazborshchik_izvlech__tcel;
stroki__TRazborshchik_class_info.vt.izvlech__shest = &stroki__TRazborshchik_izvlech__shest;
stroki__TRazborshchik_class_info.vt.izvlech__veshch = &stroki__TRazborshchik_izvlech__veshch;
stroki__TRazborshchik_class_info.vt.razobrat__razmeshchenie = &stroki__TRazborshchik_razobrat__razmeshchenie;
stroki__TRazborshchik_class_info.vt.razobrat__chislovoyi_argument = &stroki__TRazborshchik_razobrat__chislovoyi_argument;
stroki__TRazborshchik_class_info.vt.dobavit__argument = &stroki__TRazborshchik_dobavit__argument;
stroki__TRazborshchik_class_info.vt.dobavit__po_umolchaniyu = &stroki__TRazborshchik_dobavit__po_umolchaniyu;
stroki__TRazborshchik_class_info.vt.dobavit__tip = &stroki__TRazborshchik_dobavit__tip;
stroki__TRazborshchik_class_info.vt.dobavit__str = &stroki__TRazborshchik_dobavit__str;
stroki__TRazborshchik_class_info.vt.dobavit__sim = &stroki__TRazborshchik_dobavit__sim;
stroki__TRazborshchik_class_info.vt.dobavit__adr = &stroki__TRazborshchik_dobavit__adr;
stroki__TRazborshchik_class_info.vt.dobavit__tcel = &stroki__TRazborshchik_dobavit__tcel;
stroki__TRazborshchik_class_info.vt.dobavit__shest = &stroki__TRazborshchik_dobavit__shest;
stroki__TRazborshchik_class_info.vt.dobavit__veshch = &stroki__TRazborshchik_dobavit__veshch;
stroki__TRazborshchik_class_info.vt.dobavit__stroku = &stroki__TRazborshchik_dobavit__stroku;
stroki__TRazborshchik_class_info.vt.dobavit__simvol = &stroki__TRazborshchik_dobavit__simvol;
stroki__TRazborshchik_class_info.vt.dobavit__tceloe10 = &stroki__TRazborshchik_dobavit__tceloe10;
stroki__TRazborshchik_class_info.vt.dobavit__chislo10 = &stroki__TRazborshchik_dobavit__chislo10;
stroki__TRazborshchik_class_info.vt.dobavit__znak = &stroki__TRazborshchik_dobavit__znak;
stroki__TRazborshchik_class_info.vt.dobavit__chislo16 = &stroki__TRazborshchik_dobavit__chislo16;
stroki__TRazborshchik_class_info.vt.dobavit__prefiks_chisla = &stroki__TRazborshchik_dobavit__prefiks_chisla;
stroki__TRazborshchik_class_info.vt.dobavit__bayity_reversom = &stroki__TRazborshchik_dobavit__bayity_reversom;
stroki__TRazborshchik_class_info.vt.dobavit__veshchestvennoe = &stroki__TRazborshchik_dobavit__veshchestvennoe;
stroki__TRazborshchik_class_info.meta.object_size = sizeof(struct stroki__TRazborshchik);
stroki__TRazborshchik_class_info.meta.base = NULL;
stroki__TRazborshchik_class_info.meta.name = tri_newString(18, -1, "Разборщик");
stroki__TRazborshchik_class_info_ptr = &stroki__TRazborshchik_class_info;
}
struct stroki__TSborshchik_Meta {
size_t object_size;
void* base;
TString name;
};
void stroki__TSborshchik__init__(stroki__TSborshchik o) {
stroki__TBayity loc100 = tri_newVector(sizeof(TByte), 0, 0);
o->f.bayity = loc100;
o->f.chislo_simvolov = 0;
}
struct { struct stroki__TSborshchik_VT vt; struct stroki__TSborshchik_Meta meta; } stroki__TSborshchik_class_info;
void * stroki__TSborshchik_class_info_ptr;
void stroki__TSborshchik_init() {
stroki__TSborshchik_class_info.vt.self_size = sizeof(struct stroki__TSborshchik_VT);
stroki__TSborshchik_class_info.vt.__init__ = &stroki__TSborshchik__init__;
stroki__TSborshchik_class_info.vt.f = &stroki__TSborshchik_f;
stroki__TSborshchik_class_info.vt.dobavit__stroku = &stroki__TSborshchik_dobavit__stroku;
stroki__TSborshchik_class_info.vt.stroka = &stroki__TSborshchik_stroka;
stroki__TSborshchik_class_info.vt.chislo_simvolov = &stroki__TSborshchik_chislo_simvolov;
stroki__TSborshchik_class_info.vt.dobavit__simvol = &stroki__TSborshchik_dobavit__simvol;
stroki__TSborshchik_class_info.vt.chislo_bayitov = &stroki__TSborshchik_chislo_bayitov;
stroki__TSborshchik_class_info.vt.dobavit__bayity = &stroki__TSborshchik_dobavit__bayity;
stroki__TSborshchik_class_info.meta.object_size = sizeof(struct stroki__TSborshchik);
stroki__TSborshchik_class_info.meta.base = NULL;
stroki__TSborshchik_class_info.meta.name = tri_newString(14, -1, "Сборщик");
stroki__TSborshchik_class_info_ptr = &stroki__TSborshchik_class_info;
}
static TBool init_done = false;
void stroki__init() {
if (init_done) return;
init_done = true;
yunikod__init();
utf8__init();
stroki__TRazborshchik_init();
stroki__TSborshchik_init();
stroki__maks_Slovo64 = 0xFFFFFFFFFFFFFFFF;
}