package tps

import (
	"fmt"
)

type T interface {
	String() string
}

type Primitive interface {
	T
	Primitive()
}

type Reference interface {
	T
	Class() string
}

type Method interface {
	T
	Parameters() Parameters
	Return() T
}

// ---

type primitive struct {
	desc string
}

func (t *primitive) String() string {
	return t.desc
}

func (t *primitive) Primitive() {

}
func NewPrimitive(desc string) Primitive {
	return &primitive{desc: desc}
}

// ---

type reference struct {
	class string
}

func (r *reference) String() string {
	return fmt.Sprintf("L%s;", r.class)
}
func (r *reference) Class() string {
	return r.class
}

func NewReference(class string) Reference {
	return &reference{class: class}
}

//

type method struct {
	parameters Parameters
	retturn    T
}

func (m *method) String() string {
	return fmt.Sprintf("%s%s", m.parameters, m.retturn)
}
func (m *method) Parameters() Parameters {
	return m.parameters
}
func (m *method) Return() T {
	return m.retturn
}
func NewMethod(retturn T, parameters ...T) Method {
	return &method{
		parameters: parameters,
		retturn:    retturn,
	}
}

// ---

type Parameters []T

func (p *Parameters) String() (result string) {
	for _, p := range *p {
		result += p.String()
	}
	return "(" + result + ")"
}
