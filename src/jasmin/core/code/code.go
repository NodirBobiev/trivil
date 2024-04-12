package code

type Code int

const (
	Label Code = iota

	Method
	EndMethod
	LimitStack
	LimitLocals
	Class
	Super
	Field

	Iconst
	Ldc2_w

	Istore
	Lstore
	Dstore
	Astore

	Iload
	Lload
	Dload
	Aload

	Iadd
	Ladd
	Dadd

	Isub
	Lsub
	Dsub

	Ineg
	Lneg
	Dneg

	Imul
	Lmul
	Dmul

	Lcmp
	Dcmpl

	New

	Dup

	Getstatic
	Getfield

	Putstatic
	Putfield

	Invokespecial
	Invokestatic
	Invokevirtual

	Return
	Ireturn
	Lreturn
	Dreturn
	Areturn

	Goto

	Ifne
	Ifeq
	Ifle
	Ifgt
	Ifge
	Iflt
)
