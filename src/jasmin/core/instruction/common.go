package instruction

import (
	"fmt"
	"strings"
	"trivil/jasmin/core/code"
)

type I interface {
	String() string
	OpCode() code.Code
}

type S []I

// Append adds to the back of the sequence and returns the result.
// The original sequence is changed
func (s *S) Append(other ...I) *S {
	*s = append(*s, other...)
	return s
}

// Prepend adds to the frond of the sequence and returns the result.
// Similar to Append the original sequence is changed.
func (s *S) Prepend(other ...I) *S {
	*s = append(other, *s...)
	return s
}

type instruction struct {
	desc string
	code code.Code
}

func (i *instruction) String() string {
	return i.desc
}
func (i *instruction) OpCode() code.Code {
	return i.code
}

func newI(code code.Code, desc ...interface{}) I {
	return &instruction{
		desc: joinBySpace(desc),
		code: code,
	}
}

func joinBySpace(args ...any) string {
	var builder strings.Builder

	for i, arg := range args {
		str := fmt.Sprint(arg)
		if len(str) > 0 && i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(str)
	}

	return builder.String()
}
