.class public counter/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 10
    .limit locals 2
        new counter/Counter
        dup
        invokespecial counter/Counter/<init>()V
        astore 1
        aload 1
        ldc2_w 9
        invokevirtual counter/Counter/Increase(J)J
        aload 1
        ldc2_w 50
        lneg
        invokevirtual counter/Counter/Increase(J)J
        aload 1
        invokevirtual counter/Counter/Reset()J
.end method
