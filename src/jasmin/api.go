package jasmin

import (
	"strings"
)

func Name(s ...string) string {
	return strings.Join(s, "/")
}
func Const(value any, typ Type) Instruction {
	return NewConstInstruction(typ, value)
}

func Store(localNumber int, typ Type) Instruction {
	return NewStoreInstruction(typ, localNumber)
}

func Load(localNumber int, typ Type) Instruction {
	return NewLoadInstruction(typ, localNumber)
}

func GetStatic(field string, typ Type) Instruction {
	return NewGetStaticInstruction(field, typ)
}

func GetField(field string, typ Type) Instruction {
	return NewGetFieldInstruction(field, typ)
}

func PutStatic(field string, typ Type) Instruction {
	return NewPutStaticInstruction(field, typ)
}

func PutField(field string, typ Type) Instruction {
	return NewPutFieldInstruction(field, typ)
}

func InvokeSpecial(function string, t Type) Instruction {
	return NewInvokeSpecialInstruction(function, t.(*MethodType))
}

func InvokeStatic(function string, t Type) Instruction {
	return NewInvokeStaticInstruction(function, t.(*MethodType))
}

func InvokeVirtual(function string, t Type) Instruction {
	return NewInvokeVirtualInstruction(function, t.(*MethodType))
}

func Add(t Type) Instruction {
	return NewAddInstruction(t)
}

func Sub(t Type) Instruction {
	return NewSubInstruction(t)
}

func Mul(t Type) Instruction {
	return NewMulInstruction(t)
}

func New(t Type) Instruction {
	return NewNewInstruction(t)
}

func Dup(t Type) Instruction {
	return NewDupInstruction(t)
}

func Return(t Type) Instruction {
	return NewReturnInstruction(t)
}

func Neg(t Type) Instruction {
	return NewNegInstruction(t)
}

// ---

func MainMethodType() Type {
	return NewMethodType(NewParametersType(NewArrayType(NewReferenceType("java/lang/String"))), NewVoidType())
}

func VoidMethodType() Type {
	return NewMethodType(NewParametersType(), NewVoidType())
}

func UnaryVoidMethodType(parameterType Type) Type {
	return NewMethodType(NewParametersType(parameterType), NewVoidType())
}
