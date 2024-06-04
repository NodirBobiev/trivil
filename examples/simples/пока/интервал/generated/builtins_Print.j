.class public builtins/Print
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method public static print_double(D)V
    .limit stack 3
    .limit locals 2
        getstatic java/lang/System/out Ljava/io/PrintStream;
        dload 0
        invokevirtual java/io/PrintStream/print(D)V
        return
.end method
.method public static println()V
    .limit stack 1
    .limit locals 0
        getstatic java/lang/System/out Ljava/io/PrintStream;
        invokevirtual java/io/PrintStream/println()V
        return
.end method
.method public static print_long(J)V
    .limit stack 3
    .limit locals 2
        getstatic java/lang/System/out Ljava/io/PrintStream;
        lload 0
        invokevirtual java/io/PrintStream/print(J)V
        return
.end method
