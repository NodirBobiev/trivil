package jasmin

import (
	"fmt"
)

type Primitive int

const (
	VoidPrimitive Primitive = iota + 1
	BytePrimitive
	CharPrimitive
	ShortPrimitive
	IntPrimitive
	LongPrimitive
	FloatPrimitive
	DoublePrimitive
	BooleanPrimitive
)

type Type interface {
	String() string
	StackSlot() int
}

type PrimitiveType interface {
	Type
	PrimitiveType()
}

// ---

type TypeBase struct {
	Description string
	Slots       int
}

func NewTypeBase(desc string, slots int) TypeBase {
	return TypeBase{
		Description: desc,
		Slots:       slots,
	}
}
func (t *TypeBase) String() string {
	return t.Description
}
func (t *TypeBase) StackSlot() int {
	return t.Slots
}

// PrimitiveType

type VoidType struct {
	TypeBase
}

func NewVoidType() *VoidType {
	return &VoidType{
		TypeBase: NewTypeBase("V", 0),
	}
}
func (i *VoidType) PrimitiveType() {}

// ---

type IntType struct {
	TypeBase
}

func NewIntType() *IntType {
	return &IntType{
		TypeBase: NewTypeBase("I", 1),
	}
}
func (i *IntType) PrimitiveType() {}

// ---

type LongType struct {
	TypeBase
}

func NewLongType() *LongType {
	return &LongType{
		TypeBase: NewTypeBase("J", 2),
	}
}
func (i *LongType) PrimitiveType() {}

// ---

type DoubleType struct {
	TypeBase
}

func NewDoubleType() *DoubleType {
	return &DoubleType{
		TypeBase: NewTypeBase("D", 2),
	}
}
func (i *DoubleType) PrimitiveType() {}

// ReferenceType

type ReferenceType struct {
	TypeBase
	Class string
}

func NewReferenceType(class string) *ReferenceType {
	return &ReferenceType{
		TypeBase: NewTypeBase(fmt.Sprintf("L%s;", class), 1),
		Class:    class,
	}
}

func NewStringType() *ReferenceType {
	return NewReferenceType("java/lang/String")
}

// Array Type

type ArrayType struct {
	TypeBase
	Type Type
}

func NewArrayType(t Type) *ArrayType {
	return &ArrayType{
		TypeBase: NewTypeBase(fmt.Sprintf("[%s", t), 1),
		Type:     t,
	}
}

// Function Type

type MethodType struct {
	TypeBase
	Parameters Type
	Return     Type
}

func NewMethodType(parameters Type, returnType Type) *MethodType {
	return &MethodType{
		TypeBase:   NewTypeBase(fmt.Sprintf("%s%s", parameters, returnType), 1),
		Parameters: parameters,
		Return:     returnType,
	}
}

// ParametersType

type ParametersType []Type

func NewParametersType(t ...Type) *ParametersType {
	result := &ParametersType{}
	*result = append(*result, t...)
	return result
}
func (p *ParametersType) StackSlot() (total int) {
	for _, t := range *p {
		total += t.StackSlot()
	}
	return
}
func (p *ParametersType) String() (result string) {
	for _, p := range *p {
		result += p.String()
	}
	return "(" + result + ")"
}
