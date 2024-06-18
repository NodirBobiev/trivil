.class public Main
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 1
    .limit locals 1
        invokestatic builtins/Scan/initScanner()V
        aload 0
        invokestatic tekst/MainClass/main([Ljava/lang/String;)V
        return
.end method
