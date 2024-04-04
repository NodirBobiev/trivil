package codes

type Code int

const (
	Label = iota

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

	Ladd
	Dadd

	Lsub
	Dsub

	Lneg
	Dneg

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
	Lreturn
	Dreturn
	Areturn
)
