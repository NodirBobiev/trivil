.class public simple_class/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 8
    .limit locals 2
        new simple_class/Human
        dup
        invokespecial simple_class/Human/<init>()V
        dup
        ldc2_w 180.4
        putfield simple_class/Human/height D
        astore 1
        aload 1
        aload 1
        getfield simple_class/Human/height D
        ldc2_w 20.0
        dadd
        putfield simple_class/Human/height D
.end method
