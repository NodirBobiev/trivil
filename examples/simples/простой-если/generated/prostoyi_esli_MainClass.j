.class public prostoyi_esli/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 4
    .limit locals 5
        ldc2_w 100
        lstore 1
        ldc2_w 200
        lstore 3
        lload 1
        lload 3
        lcmp
        ifne ELSE_IF_2
        ldc2_w 1110
        invokestatic builtins/Print/print_long(J)V
        goto END_IF_1
        ELSE_IF_2:
        lload 1
        lload 3
        lcmp
        ifeq END_IF_1
        ldc2_w 1120
        invokestatic builtins/Print/print_long(J)V
        END_IF_1:
        lload 1
        lload 3
        lcmp
        ifeq ELSE_IF_4
        ldc2_w 1210
        invokestatic builtins/Print/print_long(J)V
        goto END_IF_3
        ELSE_IF_4:
        lload 1
        lload 3
        lcmp
        ifne END_IF_3
        ldc2_w 1220
        invokestatic builtins/Print/print_long(J)V
        END_IF_3:
        lload 1
        lload 3
        lcmp
        ifle ELSE_IF_6
        ldc2_w 2110
        invokestatic builtins/Print/print_long(J)V
        goto END_IF_5
        ELSE_IF_6:
        lload 1
        lload 3
        lcmp
        ifge ELSE_IF_7
        ldc2_w 2120
        invokestatic builtins/Print/print_long(J)V
        goto END_IF_5
        ELSE_IF_7:
        ldc2_w 2130
        invokestatic builtins/Print/print_long(J)V
        END_IF_5:
        lload 3
        lload 1
        lcmp
        ifle ELSE_IF_9
        ldc2_w 2210
        invokestatic builtins/Print/print_long(J)V
        goto END_IF_8
        ELSE_IF_9:
        lload 3
        lload 1
        lcmp
        ifge ELSE_IF_10
        ldc2_w 2220
        invokestatic builtins/Print/print_long(J)V
        goto END_IF_8
        ELSE_IF_10:
        ldc2_w 2230
        invokestatic builtins/Print/print_long(J)V
        END_IF_8:
        lload 1
        lload 3
        lcmp
        ifle ELSE_IF_12
        ldc2_w 2310
        invokestatic builtins/Print/print_long(J)V
        goto END_IF_11
        ELSE_IF_12:
        lload 3
        lload 1
        lcmp
        ifge ELSE_IF_13
        ldc2_w 2320
        invokestatic builtins/Print/print_long(J)V
        goto END_IF_11
        ELSE_IF_13:
        ldc2_w 2330
        invokestatic builtins/Print/print_long(J)V
        END_IF_11:
        lload 3
        lload 1
        lcmp
        iflt ELSE_IF_15
        ldc2_w 3110
        invokestatic builtins/Print/print_long(J)V
        goto END_IF_14
        ELSE_IF_15:
        lload 3
        lload 1
        lcmp
        ifgt END_IF_14
        ldc2_w 3120
        invokestatic builtins/Print/print_long(J)V
        END_IF_14:
        lload 1
        lload 3
        lcmp
        iflt ELSE_IF_17
        ldc2_w 3210
        invokestatic builtins/Print/print_long(J)V
        goto END_IF_16
        ELSE_IF_17:
        lload 1
        lload 3
        lcmp
        ifgt END_IF_16
        ldc2_w 3220
        invokestatic builtins/Print/print_long(J)V
        END_IF_16:
        return
.end method
