.class public print/MainClass
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
    .limit locals 1
        ldc2_w 10
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        ldc2_w 2.2345435
        invokestatic builtins/Print/print_double(D)V
        invokestatic builtins/Print/println()V
        return
.end method
