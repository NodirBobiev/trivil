#include "rt_api.h"
#include "vyvod.h"

//--- globals
static TString strlit1 = NULL;
//--- end globals

void vyvod__simvol(TSymbol s) {
TWord64 loc2[2] = {tri_tagTSymbol(), s};
vyvod__f(tri_newLiteralString(&strlit1, -1, 5, "$сим;"), 1, &loc2);
}
void vyvod__f(TString format, TInt64 spisok_len, void* spisok) {
stroki__TSborshchik loc3 = tri_newObject(stroki__TSborshchik_class_info_ptr);
stroki__TSborshchik sb = loc3;
sb->vtable->f((stroki__TSborshchik)sb, format, spisok_len, spisok);
print_string(sb->vtable->stroka((stroki__TSborshchik)sb));
}
static TBool init_done = false;
void vyvod__init() {
if (init_done) return;
init_done = true;
stroki__init();
}