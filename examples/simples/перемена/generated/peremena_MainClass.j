.class public peremena/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 5
    .limit locals 7
        ldc2_w 100
        lstore 1
        ldc2_w 1000
        lstore 3
        lload 1
        lload 3
        ladd
        lstore 5
        lload 3
        lload 5
        lsub
        lstore 1
        lload 1
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        return
.end method
