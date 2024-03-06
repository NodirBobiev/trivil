.class public importer/Counter
.super java/lang/Object
.field protected value J
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method protected Increase(J)J
    .limit stack 6
    .limit locals 3
        aload 0
        aload 0
        getfield importer/Counter/value J
        lload 1
        ladd
        putfield importer/Counter/value J
        aload 0
        getfield importer/Counter/value J
        lreturn
.end method
