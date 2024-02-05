#include "rt_api.h"
#include "privet.h"

//--- globals
static TString strlit1 = NULL;
static TString strlit3 = NULL;
static TString strlit5 = NULL;
//--- end globals

void privet__privet(TString imya) {
if (imya == NULL) {
TWord64 loc2[0] = {};
vyvod__f(tri_newLiteralString(&strlit1, -1, 15, "привет аноним!\n"), 0, &loc2);
}
else {
TWord64 loc4[2] = {tri_tagTString(), ((TUnion64)(void*)(TString)tri_nilcheck(imya,"../examples/privet/privet.tri:9:36")).w};
vyvod__f(tri_newLiteralString(&strlit3, -1, 11, "привет $;!\n"), 1, &loc4);
}
}
int main(int argc, char *argv[]) {
tri_init(argc, argv);
vyvod__init();
privet__privet((TString)tri_newLiteralString(&strlit5, -1, 4, "Вася"));
privet__privet((TString)NULL);
  return 0;
}