.class public simple_type/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 7
    .limit locals 2
        new simple_type/Human
        dup
        invokespecial simple_type/Human/<init>()V
        dup
        ldc2_w 180.4
        putfield simple_type/Human/height D
        astore 1
        aload 1
        aload 1
        getfield simple_type/Human/height D
        ldc2_w 20.0
        dadd
        putfield simple_type/Human/height D

        ;--- this part is manually added to print the result ---
        getstatic java/lang/System/out Ljava/io/PrintStream;
        aload 1
        getfield simple_type/Human/height D
        invokevirtual java/io/PrintStream/println(D)V
        ;-------------------------------------------------------

        return
.end method
