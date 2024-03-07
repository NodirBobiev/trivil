.class public schetchik/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 15
    .limit locals 2
        new schetchik/Schetchik
        dup
        invokespecial schetchik/Schetchik/<init>()V
        astore 1
        aload 1
        ldc2_w 9
        invokevirtual schetchik/Schetchik/Uvelich(J)J
        aload 1
        getfield schetchik/Schetchik/znachenie J
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        aload 1
        ldc2_w 50
        lneg
        invokevirtual schetchik/Schetchik/Uvelich(J)J
        aload 1
        getfield schetchik/Schetchik/znachenie J
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        aload 1
        invokevirtual schetchik/Schetchik/Sbros()J
        aload 1
        getfield schetchik/Schetchik/znachenie J
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        return
.end method
