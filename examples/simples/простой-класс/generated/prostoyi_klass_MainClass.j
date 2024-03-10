.class public prostoyi_klass/MainClass
.super java/lang/Object
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method public static main([Ljava/lang/String;)V
    .limit stack 12
    .limit locals 5
        new prostoyi_klass/CHelovek
        dup
        invokespecial prostoyi_klass/CHelovek/<init>()V
        astore 1
        ldc2_w 1000.0
        aload 1
        invokestatic prostoyi_klass/MainClass/Novaya_mashina(DLprostoyi_klass/CHelovek;)Lprostoyi_klass/Mashina;
        astore 2
        ldc2_w 20000.55
        aload 1
        invokestatic prostoyi_klass/MainClass/Novaya_mashina(DLprostoyi_klass/CHelovek;)Lprostoyi_klass/Mashina;
        astore 3
        aload 1
        getfield prostoyi_klass/CHelovek/kol_mashin J
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        aload 1
        getfield prostoyi_klass/CHelovek/sostoyanie D
        invokestatic builtins/Print/print_double(D)V
        invokestatic builtins/Print/println()V
        invokestatic builtins/Print/println()V
        new prostoyi_klass/CHelovek
        dup
        invokespecial prostoyi_klass/CHelovek/<init>()V
        dup
        ldc2_w 30
        putfield prostoyi_klass/CHelovek/vozrast J
        astore 4
        aload 4
        aload 2
        invokevirtual prostoyi_klass/CHelovek/kupit_(Lprostoyi_klass/Mashina;)V
        aload 1
        getfield prostoyi_klass/CHelovek/kol_mashin J
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        aload 1
        getfield prostoyi_klass/CHelovek/sostoyanie D
        invokestatic builtins/Print/print_double(D)V
        invokestatic builtins/Print/println()V
        aload 4
        getfield prostoyi_klass/CHelovek/kol_mashin J
        invokestatic builtins/Print/print_long(J)V
        invokestatic builtins/Print/println()V
        aload 4
        getfield prostoyi_klass/CHelovek/sostoyanie D
        invokestatic builtins/Print/print_double(D)V
        invokestatic builtins/Print/println()V
        return
.end method
.method protected static Novaya_mashina(DLprostoyi_klass/CHelovek;)Lprostoyi_klass/Mashina;
    .limit stack 13
    .limit locals 4
        new prostoyi_klass/Mashina
        dup
        invokespecial prostoyi_klass/Mashina/<init>()V
        dup
        dload 0
        putfield prostoyi_klass/Mashina/tcena D
        dup
        aload 2
        putfield prostoyi_klass/Mashina/xozyain Lprostoyi_klass/CHelovek;
        astore 3
        aload 2
        aload 2
        getfield prostoyi_klass/CHelovek/sostoyanie D
        aload 3
        getfield prostoyi_klass/Mashina/tcena D
        dadd
        putfield prostoyi_klass/CHelovek/sostoyanie D
        aload 2
        aload 2
        getfield prostoyi_klass/CHelovek/kol_mashin J
        ldc2_w 1
        ladd
        putfield prostoyi_klass/CHelovek/kol_mashin J
        aload 3
        areturn
.end method
