.class public fibonachi/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method protected static fib(J)J
    .limit stack 6
    .limit locals 2
        lload 0
        ldc2_w 0
        lcmp
        ifgt ELSE_IF_2
        ldc2_w 0
        lreturn
        goto END_IF_1
        ELSE_IF_2:
        lload 0
        ldc2_w 1
        lcmp
        ifne END_IF_1
        ldc2_w 1
        lreturn
        END_IF_1:
        lload 0
        ldc2_w 1
        lsub
        invokestatic fibonachi/MainClass/fib(J)J
        lload 0
        ldc2_w 2
        lsub
        invokestatic fibonachi/MainClass/fib(J)J
        ladd
        lreturn
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 2
    .limit locals 1
        ldc2_w 35
        invokestatic fibonachi/MainClass/fib(J)J
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        return
.end method
