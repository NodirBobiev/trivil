.class public importer/MainClass
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
    .limit locals 1
        ldc2_w 5
        ldc2_w 6
        invokestatic arifmetika/MainClass/summa(JJ)J
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        ldc2_w 5
        ldc2_w 6
        invokestatic arifmetika/MainClass/raznitca(JJ)J
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        ldc2_w 5.5
        ldc2_w 6.5
        invokestatic arifmetika/MainClass/ploshchad_(DD)D
        invokestatic builtins/Print/print_double(D)V
        invokestatic builtins/Print/println()V
        return
.end method
