pushd . && cd /home/cyrus/trivil/src/_genc && clang yunikod.c utf8.c stroki.c vyvod.c privet.c /home/cyrus/trivil/runtime/rt_api.c /home/cyrus/trivil/runtime/rt_sysapi.c /home/cyrus/trivil/runtime/rt_syscrash_linux.c -lm -rdynamic -I/home/cyrus/trivil/runtime -o ../privet.exe && popd