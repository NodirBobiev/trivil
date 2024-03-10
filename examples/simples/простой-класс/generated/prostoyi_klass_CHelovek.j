.class public prostoyi_klass/CHelovek
.super java/lang/Object
.field protected kol_mashin J
.field protected sostoyanie D
.field protected vozrast J
.method public <init>()V
    .limit stack 1
    .limit locals 1
        aload 0
        invokespecial java/lang/Object/<init>()V
        return
.end method
.method protected kupit_(Lprostoyi_klass/Mashina;)V
    .limit stack 19
    .limit locals 2
        aload 1
        getfield prostoyi_klass/Mashina/xozyain Lprostoyi_klass/CHelovek;
        aload 1
        getfield prostoyi_klass/Mashina/xozyain Lprostoyi_klass/CHelovek;
        getfield prostoyi_klass/CHelovek/kol_mashin J
        ldc2_w 1
        lsub
        putfield prostoyi_klass/CHelovek/kol_mashin J
        aload 1
        getfield prostoyi_klass/Mashina/xozyain Lprostoyi_klass/CHelovek;
        aload 1
        getfield prostoyi_klass/Mashina/xozyain Lprostoyi_klass/CHelovek;
        getfield prostoyi_klass/CHelovek/sostoyanie D
        aload 1
        getfield prostoyi_klass/Mashina/tcena D
        dsub
        putfield prostoyi_klass/CHelovek/sostoyanie D
        aload 0
        aload 0
        getfield prostoyi_klass/CHelovek/kol_mashin J
        ldc2_w 1
        ladd
        putfield prostoyi_klass/CHelovek/kol_mashin J
        aload 0
        aload 0
        getfield prostoyi_klass/CHelovek/sostoyanie D
        aload 1
        getfield prostoyi_klass/Mashina/tcena D
        dadd
        putfield prostoyi_klass/CHelovek/sostoyanie D
        aload 1
        aload 0
        putfield prostoyi_klass/Mashina/xozyain Lprostoyi_klass/CHelovek;
        return
.end method
