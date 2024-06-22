.class public arifmetika/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method public static ploshchad_(DD)D
    .limit stack 4
    .limit locals 4
        dload 0
        dload 2
        dmul
        dreturn
.end method
.method public static summa(JJ)J
    .limit stack 4
    .limit locals 4
        lload 0
        lload 2
        ladd
        lreturn
.end method
.method public static raznitca(JJ)J
    .limit stack 4
    .limit locals 4
        lload 0
        lload 2
        lsub
        lreturn
.end method
