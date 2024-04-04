package core

import (
	"trivil/jasmin/codes"
	types "trivil/jasmin/types"
)

func Iconst(value int) Instruction {
	return NewInstruction(codes.Iconst, 1, 0, "iconst", value)
}

func Ldc2_w(value any) Instruction {
	return NewInstruction(codes.Ldc2_w, 2, 0, "ldc2_w", value)
}

func Istore(local int) Instruction {
	return NewInstruction(codes.Istore, 0, 1, "istore", local)
}

func Lstore(local int) Instruction {
	return NewInstruction(codes.Lstore, 0, 2, "lstore", local)
}

func Dstore(local int) Instruction {
	return NewInstruction(codes.Dstore, 0, 2, "dstore", local)
}

func Astore(local int) Instruction {
	return NewInstruction(codes.Astore, 0, 1, "astore", local)
}

func Iload(local int) Instruction {
	return NewInstruction(codes.Iload, 1, 0, "iload", local)
}

func Lload(local int) Instruction {
	return NewInstruction(codes.Lload, 2, 0, "lload", local)
}

func Dload(local int) Instruction {
	return NewInstruction(codes.Dload, 2, 0, "dload", local)
}

func Aload(local int) Instruction {
	return NewInstruction(codes.Aload, 1, 0, "aload", local)
}

func Ladd() Instruction {
	return NewInstruction(codes.Ladd, 2, 4, "ladd")
}

func Dadd() Instruction {
	return NewInstruction(codes.Dadd, 2, 4, "dadd")
}

func Lsub() Instruction {
	return NewInstruction(codes.Lsub, 2, 4, "lsub")
}

func Dsub() Instruction {
	return NewInstruction(codes.Dsub, 2, 4, "dsub")
}

func Lmul() Instruction {
	return NewInstruction(codes.Lmul, 2, 4, "lmul")
}

func Dmul() Instruction {
	return NewInstruction(codes.Dmul, 2, 4, "dmul")
}

func Lneg() Instruction {
	return NewInstruction(codes.Lneg, 2, 2, "lneg")
}

func Dneg() Instruction {
	return NewInstruction(codes.Dneg, 2, 2, "dneg")
}

func Lcmp() Instruction {
	return NewInstruction(codes.Lcmp, 2, 2, "lcmp")
}

func Dcmpl() Instruction {
	return NewInstruction(codes.Dcmpl, 2, 4, "dcmpl")
}

func New(class string) Instruction {
	return NewInstruction(codes.New, 1, 0, "new", class)
}

func Dup() Instruction {
	return NewInstruction(codes.Dup, 1, 0, "dup")
}

func Getstatic(field string, fieldType types.Type) Instruction {
	return NewInstruction(codes.Getstatic, fieldType.StackSlot(), 0, "getstatic", field, fieldType)
}

func Getfield(field string, fieldType types.Type) Instruction {
	return NewInstruction(codes.Getfield, fieldType.StackSlot(), 0, "getfield", field, fieldType)
}

func Putfield(field string, fieldType types.Type) Instruction {
	return NewInstruction(codes.Putfield, 0, fieldType.StackSlot(), "putfield", field, fieldType)
}

func Putstatic(field string, fieldType types.Type) Instruction {
	return NewInstruction(codes.Putstatic, 0, fieldType.StackSlot(), "putstatic", field, fieldType)
}

func Invokespecial(fn string, fnType *types.MethodType) Instruction {
	return NewInstruction(codes.Invokespecial, fnType.Return.StackSlot(), fnType.Parameters.StackSlot(), "invokespecial", fn)
}

func Invokestatic(fn string, fnType *types.MethodType) Instruction {
	return NewInstruction(codes.Invokestatic, fnType.Return.StackSlot(), fnType.Parameters.StackSlot(), "invokestatic", fn)
}

func Invokevirtual(fn string, fnType *types.MethodType) Instruction {
	return NewInstruction(codes.Invokevirtual, fnType.Return.StackSlot(), fnType.Parameters.StackSlot(), "invokevirtual", fn)
}

func Return() Instruction {
	return NewInstruction(codes.Return, 0, 0, "return")
}

func Lreturn() Instruction {
	return NewInstruction(codes.Lreturn, 0, 2, "lreturn")
}

func Dreturn() Instruction {
	return NewInstruction(codes.Dreturn, 0, 2, "dreturn")
}

func Areturn() Instruction {
	return NewInstruction(codes.Areturn, 0, 1, "areturn")
}
