package stack

import (
	"trivil/jasmin/core/instruction"
)

type slots struct {
	pop  int
	push int
}

var (
	instructions = map[instruction.I]slots{}
)

func Set(i instruction.I, pop, push int) {
	instructions[i] = slots{
		pop:  pop,
		push: push,
	}
}

func Get(i instruction.I) (pop, push int) {
	var (
		s  slots
		ok bool
	)
	if s, ok = instructions[i]; !ok {
		s = opCodes[i.OpCode()]
	}
	return s.pop, s.push
}
