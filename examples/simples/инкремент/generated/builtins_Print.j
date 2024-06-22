.class public builtins/Print
.super java/lang/Object

.method public <init>()V
    .limit stack 1
    .limit locals 1
    aload_0
    invokespecial java/lang/Object/<init>()V
    return
.end method

.method public static println()V
    .limit stack 2
    .limit locals 0
    getstatic java/lang/System/out Ljava/io/PrintStream;
    invokevirtual java/io/PrintStream/println()V
    return
.end method

.method public static print_int64(J)V
    .limit stack 3
    .limit locals 2
    getstatic java/lang/System/out Ljava/io/PrintStream;
    lload_0
    invokevirtual java/io/PrintStream/print(J)V
    return
.end method

.method public static print_float64(D)V
    .limit stack 3
    .limit locals 2
    getstatic java/lang/System/out Ljava/io/PrintStream;
    dload_0
    invokevirtual java/io/PrintStream/print(D)V
    return
.end method

.method public static print_string(Ljava/lang/String;)V
    .limit stack 2
    .limit locals 1
    getstatic java/lang/System/out Ljava/io/PrintStream;
    aload_0
    invokevirtual java/io/PrintStream/print(Ljava/lang/String;)V
    return
.end method

.method public static print_bool(Z)V
    .limit stack 2
    .limit locals 1
    getstatic java/lang/System/out Ljava/io/PrintStream;
    iload_0
    invokevirtual java/io/PrintStream/print(Z)V
    return
.end method
