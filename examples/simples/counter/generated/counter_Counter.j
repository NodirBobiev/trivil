.class public counter/Counter
.super java/lang/Object
.field protected value J
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
.end method
.method protected Increase(J)J
    .limit stack 6
    .limit locals 3
        aload 0
        aload 0
        getfield counter/Counter/value J
        lload 1
        ladd
        putfield counter/Counter/value J
        aload 0
        getfield counter/Counter/value J
        lreturn
.end method
.method protected Reset()J
    .limit stack 4
    .limit locals 3
        aload 0
        getfield counter/Counter/value J
        lstore 1
        aload 0
        ldc2_w 0
        putfield counter/Counter/value J
        lload 1
        lreturn
.end method
