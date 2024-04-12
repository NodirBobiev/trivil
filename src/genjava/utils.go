package genjava

import (
	"fmt"
	"trivil/jasmin/core/instruction"
	"trivil/jasmin/core/tps"
	"trivil/lexer"
)

func negate(op lexer.Token) string {
	//case lexer.GTR, lexer.GEQ, lexer.LSS, lexer.LEQ, lexer.EQ, lexer.NEQ:
	switch op {
	case lexer.EQ:
		return "ne"
	case lexer.NEQ:
		return "eq"
	case lexer.GTR:
		return "le"
	case lexer.LEQ:
		return "gt"
	case lexer.LSS:
		return "ge"
	case lexer.GEQ:
		return "lt"
	default:
		panic(fmt.Sprintf("toIfIcmp doesn't support: %v lexer token", op))
	}
}

func toIfIcmp(op lexer.Token) string {
	return "if_icmp" + negate(op)
}

func toIf(op lexer.Token) string {
	return "if" + negate(op)
}

func loadInstruction(t tps.T, local int) instruction.I {
	switch t {
	case tps.Int:
		return instruction.Iload(local)
	case tps.Long:
		return instruction.Lload(local)
	case tps.Double:
		return instruction.Dload(local)
	default:
		return instruction.Aload(local)
	}
}

func storeInstruction(t tps.T, local int) instruction.I {
	switch t {
	case tps.Int:
		return instruction.Istore(local)
	case tps.Long:
		return instruction.Lstore(local)
	case tps.Double:
		return instruction.Dstore(local)
	default:
		return instruction.Astore(local)
	}
}

func returnInstruction(t tps.T) instruction.I {
	switch t {
	case tps.Void:
		return instruction.Return()
	case tps.Int:
		return instruction.Ireturn()
	case tps.Long:
		return instruction.Lreturn()
	case tps.Double:
		return instruction.Dreturn()
	default:
		return instruction.Areturn()
	}
}

func cmpInstruction(t tps.T) instruction.I {
	switch t {
	case tps.Long:
		return instruction.Lcmp()
	default:
		return instruction.Dcmpl()
	}
}

func ifInstruction(eq string, label string) instruction.I {
	switch eq {
	case "ne":
		return instruction.Ifne(label)
	case "eq":
		return instruction.Ifeq(label)
	case "le":
		return instruction.Ifle(label)
	case "gt":
		return instruction.Ifgt(label)
	case "ge":
		return instruction.Ifge(label)
	default:
		return instruction.Iflt(label)
	}
}

func addInstruction(t tps.T) instruction.I {
	switch t {
	case tps.Long:
		return instruction.Ladd()
	case tps.Double:
		return instruction.Dadd()
	default:
		return instruction.Iadd()
	}
}

func subInstruction(t tps.T) instruction.I {
	switch t {
	case tps.Long:
		return instruction.Lsub()
	case tps.Double:
		return instruction.Dsub()
	default:
		return instruction.Isub()
	}
}

func mulInstruction(t tps.T) instruction.I {
	switch t {
	case tps.Long:
		return instruction.Lmul()
	case tps.Double:
		return instruction.Dmul()
	default:
		return instruction.Imul()
	}
}

func negInstruction(t tps.T) instruction.I {
	switch t {
	case tps.Long:
		return instruction.Lneg()
	case tps.Double:
		return instruction.Dneg()
	default:
		return instruction.Ineg()
	}
}

func typeSize(t tps.T) int {
	switch t {
	case tps.Long, tps.Double:
		return 2
	default:
		return 1
	}
}
