package jasmin

import (
	"fmt"
)

type Instruction interface {
	String() string // full representation of instruction
	In() int        // how many entries are added onto the stack
	Out() int       // how many entries are removed from the stack
}

type Sequence []Instruction

func NewSequence(i ...Instruction) Sequence {
	return i
}

type InstructionBase struct {
	in  int
	out int
}

func NewInstructionBase(in, out int) InstructionBase {
	return InstructionBase{
		in:  in,
		out: out,
	}
}
func (i *InstructionBase) String() string {
	return ""
}
func (i *InstructionBase) In() int {
	return i.in
}
func (i *InstructionBase) Out() int {
	return i.out
}

// -----

type ConstInstruction struct {
	InstructionBase
	Type  Type
	Value any
}

func NewConstInstruction(t Type, v any) *ConstInstruction {
	return &ConstInstruction{
		InstructionBase: NewInstructionBase(t.StackSlot(), 0),
		Type:            t,
		Value:           v,
	}
}
func (c *ConstInstruction) String() string {
	switch c.Type.(type) {
	case *IntType:
		return fmt.Sprintf("iconst %v", c.Value)
	case *LongType:
		return fmt.Sprintf("ldc2_w %v", c.Value)
	case *DoubleType:
		return fmt.Sprintf("ldc2_w %v", c.Value)
	default:
		panic(fmt.Sprintf("unexpected type: %+v", c.Type))
	}
}

// ----

type StoreInstruction struct {
	InstructionBase
	Type Type
	Var  int
}

func NewStoreInstruction(t Type, v int) *StoreInstruction {
	return &StoreInstruction{
		InstructionBase: NewInstructionBase(0, t.StackSlot()),
		Type:            t,
		Var:             v,
	}
}

func (s *StoreInstruction) String() string {
	switch s.Type.(type) {
	case *IntType:
		return fmt.Sprintf("istore %v", s.Var)
	case *LongType:
		return fmt.Sprintf("lstore %v", s.Var)
	case *DoubleType:
		return fmt.Sprintf("dstore %v", s.Var)
	case *ReferenceType:
		return fmt.Sprintf("astore %v", s.Var)
	default:
		panic(fmt.Sprintf("unexpected type: %+v", s.Type))
	}
}

// ----

type LoadInstruction struct {
	InstructionBase
	Type Type
	Var  int
}

func NewLoadInstruction(t Type, l int) *LoadInstruction {
	return &LoadInstruction{
		InstructionBase: NewInstructionBase(t.StackSlot(), 0),
		Type:            t,
		Var:             l,
	}
}

func (l *LoadInstruction) String() string {
	switch l.Type.(type) {
	case *IntType:
		return fmt.Sprintf("iload %v", l.Var)
	case *LongType:
		return fmt.Sprintf("lload %v", l.Var)
	case *DoubleType:
		return fmt.Sprintf("dload %v", l.Var)
	case *ReferenceType:
		return fmt.Sprintf("aload %v", l.Var)
	default:
		panic(fmt.Sprintf("unexpected type: %+v", l.Type))
	}
}

// ---

type GetStaticInstruction struct {
	InstructionBase
	Field string
	Type  Type
}

func NewGetStaticInstruction(field string, typ Type) *GetStaticInstruction {
	return &GetStaticInstruction{
		InstructionBase: NewInstructionBase(typ.StackSlot(), 0),
		Field:           field,
		Type:            typ,
	}
}
func (g *GetStaticInstruction) String() string {
	return fmt.Sprintf("getstatic %s %s", g.Field, g.Type)
}

// ---

type GetFieldInstruction struct {
	InstructionBase
	Field string
	Type  Type
}

func NewGetFieldInstruction(field string, typ Type) *GetFieldInstruction {
	return &GetFieldInstruction{
		InstructionBase: NewInstructionBase(typ.StackSlot(), 0),
		Field:           field,
		Type:            typ,
	}
}
func (g *GetFieldInstruction) String() string {
	return fmt.Sprintf("getfield %s %s", g.Field, g.Type)
}

// ---

type PutFieldInstruction struct {
	InstructionBase
	Field string
	Type  Type
}

func NewPutFieldInstruction(field string, typ Type) *PutFieldInstruction {
	return &PutFieldInstruction{
		InstructionBase: NewInstructionBase(0, typ.StackSlot()),
		Field:           field,
		Type:            typ,
	}
}
func (p *PutFieldInstruction) String() string {
	return fmt.Sprintf("putfield %s %s", p.Field, p.Type)
}

// ---

type PutStaticInstruction struct {
	InstructionBase
	Field string
	typ   Type
}

func NewPutStaticInstruction(field string, typ Type) *PutStaticInstruction {
	return &PutStaticInstruction{
		InstructionBase: NewInstructionBase(0, typ.StackSlot()),
		Field:           field,
		typ:             typ,
	}
}
func (p *PutStaticInstruction) String() string {
	return fmt.Sprintf("putstatic %s %s", p.Field, p.typ)
}

// ---

type InvokeSpecialInstruction struct {
	InstructionBase
	Method string
	// TODO: Implement Method type as well.
}

func NewInvokeSpecialInstruction(function string) *InvokeSpecialInstruction {
	return &InvokeSpecialInstruction{
		InstructionBase: NewInstructionBase(0, 1),
		Method:          function,
	}
}
func (i *InvokeSpecialInstruction) String() string {
	return fmt.Sprintf("invokespecial %s", i.Method)
}

// ---

type InvokeStaticInstruction struct {
	InstructionBase
	Method string
}

func NewInvokeStaticInstruction(function string) *InvokeStaticInstruction {
	return &InvokeStaticInstruction{
		InstructionBase: NewInstructionBase(0, 1),
		Method:          function,
	}
}
func (i *InvokeStaticInstruction) String() string {
	return fmt.Sprintf("invokestatic %s", i.Method)
}

// ---

type AddInstruction struct {
	InstructionBase
	Type Type
}

func NewAddInstruction(t Type) *AddInstruction {
	return &AddInstruction{
		InstructionBase: NewInstructionBase(1, 2),
		Type:            t,
	}
}
func (a *AddInstruction) String() string {
	switch a.Type.(type) {
	case *LongType:
		return "ladd"
	case *DoubleType:
		return "dadd"
	default:
		panic(fmt.Sprintf("add instruction: unexpected type: %+v", a.Type))
	}
}

// ------

type SubInstruction struct {
	InstructionBase
	Type Type
}

func NewSubInstruction(t Type) *SubInstruction {
	return &SubInstruction{
		InstructionBase: NewInstructionBase(t.StackSlot(), 2*t.StackSlot()),
		Type:            t,
	}
}
func (a *SubInstruction) String() string {
	switch a.Type.(type) {
	case *LongType:
		return "lsub"
	case *DoubleType:
		return "lsub"
	default:
		panic(fmt.Sprintf("sub instruction: unexpected type: %+v", a.Type))
	}
}

// ----

type NewInstruction struct {
	InstructionBase
	Type Type
}

func NewNewInstruction(t Type) *NewInstruction {
	return &NewInstruction{
		InstructionBase: NewInstructionBase(t.StackSlot(), 0),
		Type:            t,
	}
}
func (i *NewInstruction) String() string {
	if x, ok := i.Type.(*ReferenceType); ok {
		return fmt.Sprintf("new %s", x.Class)
	}
	panic(fmt.Sprintf("new instruction: unexpected type: %+v", i.Type))
}

// ----

type DupInstruction struct {
	InstructionBase
	Type Type
}

func NewDupInstruction(t Type) *DupInstruction {
	return &DupInstruction{
		InstructionBase: NewInstructionBase(t.StackSlot(), 0),
		Type:            t,
	}
}
func (i *DupInstruction) String() string {
	return "dup"
}
