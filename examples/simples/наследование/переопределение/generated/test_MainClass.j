.class public test/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 2
    .limit locals 2
        new test/Derived
        dup
        invokespecial test/Derived/<init>()V
        astore 1
        aload 1
        invokevirtual test/Base/Vyvod()V
        invokestatic builtins/Print/println()V
        return
.end method
