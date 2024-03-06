.class public importer/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method protected static add(JJ)J
    .limit stack 4
    .limit locals 4
        lload 0
        lload 2
        ladd
        lreturn
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 8
    .limit locals 2
        new importer/Counter
        dup
        invokespecial importer/Counter/<init>()V
        astore 1
        aload 1
        ldc2_w 9
        invokevirtual importer/Counter/Increase(J)J
        ldc2_w 10
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        invokestatic print/MainClass/cat()V
        ldc2_w 7
        ldc2_w 9
        invokestatic importer/MainClass/add(JJ)J
        return
.end method
