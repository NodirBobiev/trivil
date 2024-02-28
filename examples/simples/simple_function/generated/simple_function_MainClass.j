.class public simple_function/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
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
    .limit stack 4
    .limit locals 3
        ldc2_w 5
        ldc2_w 6
        invokestatic simple_function/MainClass/add(JJ)J
        lstore 1
.end method
