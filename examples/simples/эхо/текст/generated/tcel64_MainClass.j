.class public tcel64/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 5
    .limit locals 4
        invokestatic builtins/Scan/scan_int64()J
        lstore 1
        invokestatic builtins/Scan/scanln()Ljava/lang/String;
        WHILE_START_1:
        lload 1
        ldc2_w 0
        lcmp
        ifle WHILE_END_2
        lload 1
        ldc2_w 1
        lsub
        lstore 1
        invokestatic builtins/Scan/scanln()Ljava/lang/String;
        astore 3
        aload 3
        invokestatic builtins/Print/print_string(Ljava/lang/String;)V
        lload 1
        ldc2_w 0
        lcmp
        ifle END_IF_3
        invokestatic builtins/Print/println()V
        END_IF_3:
        goto WHILE_START_1
        WHILE_END_2:
        return
.end method
