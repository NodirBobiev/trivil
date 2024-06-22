package genjava

import (
	"fmt"
	"trivil/lexer"
)

func stringify(op lexer.Token) string {
	switch op {
	case lexer.EQ:
		return "eq"
	case lexer.NEQ:
		return "ne"
	case lexer.GTR:
		return "gt"
	case lexer.LEQ:
		return "le"
	case lexer.LSS:
		return "lt"
	case lexer.GEQ:
		return "ge"
	default:
		panic(fmt.Sprintf("stringify doesn't support: %v lexer token", op))
	}
}
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
		panic(fmt.Sprintf("negate doesn't support: %v lexer token", op))
	}
}

func toIfIcmp(op lexer.Token) string {
	return "if_icmp" + negate(op)
}
func toIf(op lexer.Token) string {
	return "if" + negate(op)
}
