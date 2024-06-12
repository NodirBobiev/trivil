.class public test/Derived
.super test/Base
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial test/Base/<init>()V
        return
.end method
.method protected Vyvod()V
    .limit stack 1
    .limit locals 1
        ldc "*** From derived class ****"
        invokestatic builtins/Print/print_string(Ljava/lang/String;)V
        return
.end method
