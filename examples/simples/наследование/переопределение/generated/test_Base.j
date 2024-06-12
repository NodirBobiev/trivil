.class public test/Base
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method protected Vyvod()V
    .limit stack 1
    .limit locals 1
        ldc "From base class"
        invokestatic builtins/Print/print_string(Ljava/lang/String;)V
        return
.end method
