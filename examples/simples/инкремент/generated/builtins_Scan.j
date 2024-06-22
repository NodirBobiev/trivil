.class public builtins/Scan
.super java/lang/Object

.field private static scanner Ljava/util/Scanner;

.method public <init>()V
    .limit stack 2
    .limit locals 1
    aload_0
    invokespecial java/lang/Object/<init>()V
    return
.end method

.method public static initScanner()V
    .limit stack 3
    .limit locals 0
    new java/util/Scanner
    dup
    getstatic java/lang/System/in Ljava/io/InputStream;
    invokespecial java/util/Scanner/<init>(Ljava/io/InputStream;)V
    putstatic builtins/Scan/scanner Ljava/util/Scanner;
    return
.end method

.method public static scanln()Ljava/lang/String;
    .limit stack 2
    .limit locals 1
    getstatic builtins/Scan/scanner Ljava/util/Scanner;
    invokevirtual java/util/Scanner/nextLine()Ljava/lang/String;
    areturn
.end method

.method public static scan_int64()J
    .limit stack 2
    .limit locals 1
    getstatic builtins/Scan/scanner Ljava/util/Scanner;
    invokevirtual java/util/Scanner/nextLong()J
    lreturn
.end method

.method public static scan_float64()D
    .limit stack 2
    .limit locals 1
    getstatic builtins/Scan/scanner Ljava/util/Scanner;
    invokevirtual java/util/Scanner/nextDouble()D
    dreturn
.end method

.method public static scan_string()Ljava/lang/String;
    .limit stack 2
    .limit locals 1
    getstatic builtins/Scan/scanner Ljava/util/Scanner;
    invokevirtual java/util/Scanner/next()Ljava/lang/String;
    areturn
.end method
