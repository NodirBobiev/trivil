package jasmin

import (
	"fmt"
	"strings"
)

// Method represents method in Jasmin/JVM
// Don't set type and static after an instruction is appended
type Method struct {
	EntityBase
	EntityStorage
	Instructions []Instruction
	Locals       int
	Stack        int
	curStack     int
}

func NewMethod(name string, class *Class) *Method {
	return &Method{
		EntityBase: EntityBase{
			Name:       name,
			Type:       VoidMethodType(),
			Static:     true,
			Parent:     class,
			AccessFlag: Protected,
		},
		Instructions:  make([]Instruction, 0),
		EntityStorage: *NewEntityStorage(),
		Locals:        0,
		Stack:         0,
		curStack:      0,
	}
}

func (m *Method) Append(i ...Instruction) {
	for _, x := range i {
		m.curStack += x.In() - x.Out()
		if m.curStack > m.Stack {
			m.Stack = m.curStack
		}
	}
	m.Instructions = append(m.Instructions, i...)
}
func (m *Method) AssignNumber(v *Variable) *Variable {
	v.Number = m.Locals
	m.Locals += v.GetType().StackSlot()
	return v
}
func (m *Method) String() string {
	var (
		static string
		d      = 1
	)
	result := strings.Builder{}
	if m.IsStatic() {
		static = " static"
	}
	result.WriteString(fmt.Sprintf(".method %s%s %s%s\n", m.AccessFlag, static, m.Name, m.Type))
	result.WriteString(fmt.Sprintf("%s.limit stack %d\n", tab(d), m.Stack))
	result.WriteString(fmt.Sprintf("%s.limit locals %d\n", tab(d), m.Locals))
	d++
	for _, i := range m.Instructions {
		result.WriteString(fmt.Sprintf("%s%s\n", tab(d), i))
	}
	if _, ok := m.Type.(*MethodType).Return.(*VoidType); ok {
		result.WriteString(fmt.Sprintf("%sreturn\n", tab(d)))
	}
	result.WriteString(".end method\n")
	return result.String()
}

func (m *Method) SetType(t Type) {
	if _, ok := t.(*MethodType); !ok {
		panic(fmt.Sprintf("method set type: expected method type but got %+v", t))
	}
	m.Type = t
}
func (m *Method) SetStatic(value bool) {
	if m.Static && !value {
		m.Locals++
	} else if !m.Static && value {
		m.Locals--
	}
	m.Static = value
}
func (c *Method) GetFull() string {
	return c.EntityBase.GetFull() + c.Type.String()
}
func DefaultConstructor(class *Class, super *Class) *Method {
	m := NewMethod("<init>", class)
	m.SetStatic(false)
	m.SetAccessFlag(Public)
	if class != nil && super != nil {
		m.Append(Load(0, NewReferenceType(class.GetFull())),
			InvokeSpecial(super.Constructor.GetFull(), super.Constructor.GetType()))
	}
	return m
}

func MainMethod(class *Class) *Method {
	m := NewMethod("main", class)
	m.SetType(MainMethodType())
	m.Locals = 1
	m.SetAccessFlag(Public)
	return m
}
