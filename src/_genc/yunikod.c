#include "rt_api.h"
#include "yunikod.h"

const TSymbol yunikod__MaksSimvol = 0x10FFFF;
TBool yunikod__bukvaQm(TSymbol s) {
if (((((s >= 0x430) && (s <= 0x44F)) || ((s >= 0x410) && (s <= 0x42F))) || (s == 0x451)) || (s == 0x401)) {
return true;
}
if (((s >= 0x61) && (s <= 0x7A)) || ((s >= 0x41) && (s <= 0x5A))) {
return true;
}
return false;
}
TBool yunikod__tcifraQm(TSymbol s) {
return ((s >= 0x30) && (s <= 0x39));
}
TBool yunikod__upravlyayushchiyi_simvolQm(TSymbol s) {
return ((s <= 0x1F) || ((s >= 0x7F) && (s <= 0x9F)));
}
TBool yunikod__probel_nyyi_simvolQm(TSymbol s) {
if (!(s <= 0x20)) {
return false;
}
return ((((s == 0x20) || (s == 0x9)) || (s == 0xA)) || (s == 0xD));
}
static TBool init_done = false;
void yunikod__init() {
if (init_done) return;
init_done = true;
}