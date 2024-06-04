.class public diapazon/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method protected static rekursiya(JJJ)V
    .limit stack 7
    .limit locals 6
        lload 0
        lload 2
        lcmp
        ifle END_IF_1
        return
        END_IF_1:
        lload 0
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        lload 0
        lload 4
        ladd
        lload 2
        lload 4
        invokestatic diapazon/MainClass/rekursiya(JJJ)V
        return
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 6
    .limit locals 1
        ldc2_w 50
        lneg
        ldc2_w 50
        ldc2_w 3
        invokestatic diapazon/MainClass/rekursiya(JJJ)V
        return
.end method
