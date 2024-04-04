package core

import (
	"fmt"
	"trivil/jasmin/codes"
)

// Statement is the core of Jasmin source file. According to the Jasmin documentation
// https://jasmin.sourceforge.net/guide.html, source file consists of three types of
// statements: Directive, Instruction and Label.
type Statement interface {
	String() string
	Code() codes.Code
}

type Label interface {
	Statement
	Label() string
}

type Instruction interface {
	Statement
	Instruction()
	In() int  // how many entries are added onto the stack
	Out() int // how many entries are removed from the stack
}

type Directive interface {
	Statement
	Directive()
}

// -------------------------

type statement struct {
	code codes.Code
	desc string
}

func (s *statement) String() string {
	return s.desc
}
func (s *statement) Code() codes.Code {
	return s.code
}

// -----------------------------------------------------

type instruction struct {
	statement
	in  int
	out int
}

func (i *instruction) Instruction() {
}
func (i *instruction) In() int {
	return i.in
}
func (i *instruction) Out() int {
	return i.out
}

func NewInstruction(code codes.Code, in, out int, desc ...interface{}) Instruction {
	return &instruction{
		statement: statement{
			code: code,
			desc: joinBySpace(desc...),
		},
		in:  in,
		out: out,
	}
}

// --------------------------------------------------

type label struct {
	statement
	name string
}

func (l *label) Label() string {
	return l.name
}

func NewLabel(code codes.Code, name string) Label {
	return &label{
		statement: statement{
			code: code,
			desc: fmt.Sprintf("%:", name),
		},
		name: name,
	}
}

// ----------------------------------------------------

type directive struct {
	statement
}

func (d *directive) Directive() {
}

func NewDirective(code codes.Code, desc ...interface{}) Directive {
	return &directive{
		statement: statement{
			code: code,
			desc: joinBySpace(desc...),
		},
	}
}

// -------------------------------------------------

type Statements []Statement

// Append adds to the back of the Statements and returns the result.
// The original Statements is changed
func (s *Statements) Append(other ...Statement) *Statements {
	*s = append(*s, other...)
	return s
}

// Prepend adds to the frond of the Statements and returns the result.
// Similar to Append the original Statements is changed.
func (s *Statements) Prepend(other ...Statement) *Statements {
	*s = append(other, *s...)
	return s
}
