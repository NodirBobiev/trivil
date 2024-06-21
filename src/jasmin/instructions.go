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
		//if v, ok := c.Value.(int64); ok && v >= 0 && v <= 5 {
		//	return fmt.Sprintf("iconst_%v", c.Value)
		//} else if v, ok := c.Value.(int); ok && v >= 0 && v <= 5 {
		//	return fmt.Sprintf("iconst_%v", c.Value)
		//} else if v, ok := c.Value.(int64); ok && v == -1 {
		//	return "iconst_m1"
		//} else if v, ok := c.Value.(int); ok && v == -1 {
		//	return "iconst_m1"
		//}
		//return fmt.Sprintf("iconst %v", c.Value)
		return fmt.Sprintf("ldc %v", c.Value)
	case *LongType:
		return fmt.Sprintf("ldc2_w %v", c.Value)
	case *DoubleType:
		return fmt.Sprintf("ldc2_w %v", c.Value)
	case *ReferenceType:
		return fmt.Sprintf("ldc %v", c.Value)
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
	case *ReferenceType, *ArrayType:
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
	case *ReferenceType, *ArrayType:
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
	Type   Type
}

func NewInvokeSpecialInstruction(function string, t *MethodType) *InvokeSpecialInstruction {
	return &InvokeSpecialInstruction{
		InstructionBase: NewInstructionBase(t.Return.StackSlot(), t.Parameters.StackSlot()),
		Method:          function,
		Type:            t,
	}
}
func (i *InvokeSpecialInstruction) String() string {
	return fmt.Sprintf("invokespecial %s", i.Method)
}

// ---

type InvokeStaticInstruction struct {
	InstructionBase
	Method string
	Type   Type
}

func NewInvokeStaticInstruction(function string, t *MethodType) *InvokeStaticInstruction {
	return &InvokeStaticInstruction{
		InstructionBase: NewInstructionBase(t.Return.StackSlot(), t.Parameters.StackSlot()),
		Method:          function,
		Type:            t,
	}
}
func (i *InvokeStaticInstruction) String() string {
	return fmt.Sprintf("invokestatic %s", i.Method)
}

// ---

type InvokeVirtualInstruction struct {
	InstructionBase
	Method string
	Type   Type
}

func NewInvokeVirtualInstruction(function string, t *MethodType) *InvokeVirtualInstruction {
	return &InvokeVirtualInstruction{
		InstructionBase: NewInstructionBase(t.Return.StackSlot(), t.Parameters.StackSlot()),
		Method:          function,
		Type:            t,
	}
}
func (i *InvokeVirtualInstruction) String() string {
	return fmt.Sprintf("invokevirtual %s", i.Method)
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
		return "dsub"
	default:
		panic(fmt.Sprintf("sub instruction: unexpected type: %+v", a.Type))
	}
}

// ------

type MulInstruction struct {
	InstructionBase
	Type Type
}

func NewMulInstruction(t Type) *MulInstruction {
	return &MulInstruction{
		InstructionBase: NewInstructionBase(t.StackSlot(), 2*t.StackSlot()),
		Type:            t,
	}
}
func (i *MulInstruction) String() string {
	switch i.Type.(type) {
	case *LongType:
		return "lmul"
	case *DoubleType:
		return "dmul"
	default:
		panic(fmt.Sprintf("mul instruction: unexpected type: %+v", i.Type))
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

// ---

type ReturnInstruction struct {
	InstructionBase
	Type Type
}

func NewReturnInstruction(t Type) *ReturnInstruction {
	return &ReturnInstruction{
		InstructionBase: NewInstructionBase(0, t.StackSlot()),
		Type:            t,
	}
}
func (i *ReturnInstruction) String() string {
	switch i.Type.(type) {
	case *VoidType:
		return "return"
	case *IntType:
		return "ireturn"
	case *LongType:
		return "lreturn"
	case *DoubleType:
		return "dreturn"
	case *ReferenceType:
		return "areturn"
	default:
		panic(fmt.Sprintf("unknown type: %+v", i.Type))
	}
}

// ---

type NegInstruction struct {
	InstructionBase
	Type Type
}

func NewNegInstruction(t Type) *NegInstruction {
	return &NegInstruction{
		InstructionBase: NewInstructionBase(t.StackSlot(), t.StackSlot()),
		Type:            t,
	}
}
func (i *NegInstruction) String() string {
	switch i.Type.(type) {
	case *IntType:
		return "ineg"
	case *LongType:
		return "lneg"
	case *DoubleType:
		return "dneg"
	default:
		panic(fmt.Sprintf("unknown type: %+v", i.Type))
	}
}

// ----

type CmpInstruction struct {
	InstructionBase
	Type Type
}

func NewCmpInstruction(t Type) *CmpInstruction {
	return &CmpInstruction{
		InstructionBase: NewInstructionBase(1, 2*t.StackSlot()),
		Type:            t,
	}
}
func (i *CmpInstruction) String() string {
	switch i.Type.(type) {
	case *LongType:
		return "lcmp"
	case *DoubleType:
		return "dcmpl"
	default:
		panic(fmt.Sprintf("unknown type: %+v", i.Type))
	}
}

// ----

type IfInstruction struct {
	InstructionBase
	Code  string
	Label string
}

func NewIfInstruction(eq string, label string) *IfInstruction {
	return &IfInstruction{
		InstructionBase: NewInstructionBase(0, 1),
		Code:            "if" + eq,
		Label:           label,
	}
}
func (i *IfInstruction) String() string {
	return fmt.Sprintf("%s %s", i.Code, i.Label)
}

// ---

type IfIcmpInstruction struct {
	InstructionBase
	Code  string
	Label string
}

func NewIfIcmpInstruction(eq string, label string) *IfIcmpInstruction {
	return &IfIcmpInstruction{
		InstructionBase: NewInstructionBase(0, 2),
		Code:            "if_icmp" + eq,
		Label:           label,
	}
}
func (i *IfIcmpInstruction) String() string {
	return fmt.Sprintf("%s %s", i.Code, i.Label)
}

// ---

type GotoInstruction struct {
	InstructionBase
	Label string
}

func NewGotoInstruction(label string) *GotoInstruction {
	return &GotoInstruction{
		InstructionBase: NewInstructionBase(0, 0),
		Label:           label,
	}
}

func (i *GotoInstruction) String() string {
	return fmt.Sprintf("goto %s", i.Label)
}

// ---

type NewArrayInstruction struct {
	InstructionBase
	ElementType Type
}

func NewNewArrayInstruction(elementType Type) *NewArrayInstruction {
	return &NewArrayInstruction{
		InstructionBase: NewInstructionBase(1, 1),
		ElementType:     elementType,
	}
}

func (i *NewArrayInstruction) String() string {
	switch x := i.ElementType.(type) {
	case *LongType:
		return fmt.Sprintf("newarray long")
	case *DoubleType:
		return fmt.Sprintf("newarray double")
	case *IntType:
		return fmt.Sprintf("newarray int")
	case *ReferenceType:
		return fmt.Sprintf("anewarray %s", x.Class)
	default:
		panic(fmt.Sprintf("wrong type for newarray instruction: %+v", i.ElementType))
	}
}

// ---

type ArrayLengthInstruction struct {
	InstructionBase
}

func NewArrayLengthInstruction() *ArrayLengthInstruction {
	return &ArrayLengthInstruction{
		InstructionBase: NewInstructionBase(1, 1),
	}
}

func (i *ArrayLengthInstruction) String() string {
	return "arraylength"
}

// ---

type AstoreInstruction struct {
	InstructionBase
	ElementType Type
}

func NewAstoreInstruction(elementType Type) *AstoreInstruction {
	return &AstoreInstruction{
		InstructionBase: NewInstructionBase(0, 2+elementType.StackSlot()),
		ElementType:     elementType,
	}
}

func (i *AstoreInstruction) String() string {
	switch i.ElementType.(type) {
	case *LongType:
		return "lastore"
	case *DoubleType:
		return "dastore"
	case *IntType:
		return "iastore"
	case *ReferenceType:
		return "aastore"
	default:
		panic(fmt.Sprintf("wrong type for newarray instruction: %+v", i.ElementType))
	}
}

// ---

type AloadInstruction struct {
	InstructionBase
	ElementType Type
}

func NewAloadInstruction(elementType Type) *AloadInstruction {
	return &AloadInstruction{
		InstructionBase: NewInstructionBase(elementType.StackSlot(), 2),
		ElementType:     elementType,
	}
}

func (i *AloadInstruction) String() string {
	switch i.ElementType.(type) {
	case *LongType:
		return "laload"
	case *DoubleType:
		return "daload"
	case *IntType:
		return "iaload"
	case *ReferenceType:
		return "aaload"
	default:
		panic(fmt.Sprintf("wrong type for newarray instruction: %+v", i.ElementType))
	}
}

// --

type CastPrimitivesInstruction struct {
	InstructionBase
	FromType Type
	ToType   Type
}

func NewCastPrimitivesInstruction(fromType, toType Type) *CastPrimitivesInstruction {
	return &CastPrimitivesInstruction{
		InstructionBase: NewInstructionBase(fromType.StackSlot(), toType.StackSlot()),
		FromType:        fromType,
		ToType:          toType,
	}
}

func (i *CastPrimitivesInstruction) String() string {
	letter := func(t Type) string {
		switch t.(type) {
		case *LongType:
			return "l"
		case *DoubleType:
			return "d"
		case *IntType:
			return "i"
		}
		panic("cast instruction: wrong type")
	}

	return fmt.Sprintf("%s2%s", letter(i.FromType), letter(i.ToType))
}

// --

type IincInstruction struct {
	InstructionBase
	Var   int
	Value any
}

func NewIincInstruction(localVar int, value any) *IincInstruction {
	return &IincInstruction{
		InstructionBase: NewInstructionBase(0, 0),
		Var:             localVar,
		Value:           value,
	}
}

func (i *IincInstruction) String() string {
	return fmt.Sprintf("iinc %d %v", i.Var, i.Value)
}

// ---

type Label struct {
	InstructionBase
	Name string
}

func NewLabel(name string) *Label {
	return &Label{
		InstructionBase: NewInstructionBase(0, 0),
		Name:            name,
	}
}
func (l *Label) String() string {
	return l.Name + ":"
}
