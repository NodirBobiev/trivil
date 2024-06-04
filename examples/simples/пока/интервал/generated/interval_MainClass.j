.class public interval/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 4
    .limit locals 5
        ldc2_w 1
        lstore 1
        ldc2_w 10
        lstore 3
        WHILE_START_1:
        lload 1
        lload 3
        lcmp
        ifge WHILE_END_2
        lload 1
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        lload 1
        ldc2_w 1
        ladd
        lstore 1
        goto WHILE_START_1
        WHILE_END_2:
        return
.end method
