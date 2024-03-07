.class public schetchik/Schetchik
.super java/lang/Object
.field protected znachenie J
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method protected Uvelich(J)J
    .limit stack 6
    .limit locals 3
        aload 0
        aload 0
        getfield schetchik/Schetchik/znachenie J
        lload 1
        ladd
        putfield schetchik/Schetchik/znachenie J
        aload 0
        getfield schetchik/Schetchik/znachenie J
        lreturn
.end method
.method protected Sbros()J
    .limit stack 4
    .limit locals 3
        aload 0
        getfield schetchik/Schetchik/znachenie J
        lstore 1
        aload 0
        ldc2_w 0
        putfield schetchik/Schetchik/znachenie J
        lload 1
        lreturn
.end method
