package lexer

import (
	"fmt"
	"testing"
	"trivil/env"
)

type pair struct {
	tok Token
	lit string
}

type one struct {
	text  string
	pairs []pair
}

var tests = []one{
	{"#", []pair{{NEQ, ""}}},
	{"~", []pair{{NOT, ""}}},

	{"(:", []pair{{LCONV, ""}}},
	{"++", []pair{{INC, ""}}},
	{"<=", []pair{{LEQ, ""}}},
	{">=", []pair{{GEQ, ""}}},
	{":=", []pair{{ASSIGN, ""}}},

	{":&", []pair{{BITAND, ""}}},
	{":|", []pair{{BITOR, ""}}},
	{":\\", []pair{{BITXOR, ""}}},
	{":~", []pair{{BITNOT, ""}}},

	{"<<", []pair{{SHL, ""}}},
	{">>", []pair{{SHR, ""}}},

	{"1", []pair{{INT, "1"}}},
	{"0x1", []pair{{INT, "0x1"}}},
	{"0xA", []pair{{INT, "0xA"}}},
	{"0xa", []pair{{INT, "0xa"}}},
	{"1.", []pair{{FLOAT, "1."}}},
	{"1.0", []pair{{FLOAT, "1.0"}}},

	{"'a'", []pair{{SYMBOL, "a"}}},
	{"'\t'", []pair{{SYMBOL, "\t"}}},
	{"'\\''", []pair{{SYMBOL, "\\'"}}},
	{"'\uFFFD'", []pair{{SYMBOL, "\uFFFD"}}},

	{"\"a\"", []pair{{STRING, "a"}}},
	{"\"\t\"", []pair{{STRING, "\t"}}},
	{"\"'\"", []pair{{STRING, "'"}}},

	{"модуль", []pair{{MODULE, "модуль"}}},
	{"модуль ", []pair{{MODULE, "модуль"}}},
	{"модуль\n", []pair{{MODULE, "модуль"}, {NL, ""}}},
	{"сложное имя", []pair{{IDENT, "сложное имя"}}},
	{"№ сложное имя", []pair{{IDENT, "№ сложное имя"}}},
	{"фн сложное имя(", []pair{{FN, "фн"}, {IDENT, "сложное имя"}, {LPAR, ""}}},
	{"имя153", []pair{{IDENT, "имя153"}}},
	{"имя 153", []pair{{IDENT, "имя"}, {INT, "153"}}},
	{"как дела?", []pair{{IDENT, "как дела?"}}},
	//{"как дела ?", []pair{{IDENT, "как дела"}, {NNQUERY, ""}}},
	//{"Паниковать !", []pair{{IDENT, "Паниковать"}, {NNCHECK, ""}}},
	{"если-нет", []pair{{IDENT, "если-нет"}}},
	{"ц--", []pair{{IDENT, "ц"}, {DEC, ""}}},
	{"надо ложь иначе авария", []pair{{GUARD, "надо"}, {IDENT, "ложь"}, {ELSE, "иначе"}, {CRASH, "авария"}}},

	{"типа", []pair{{OFTYPE, "типа"}}},
}

//===

func TestValid(t *testing.T) {
	fmt.Printf("--- valid tests: %d ---\n", len(tests))
	t.Run("valid tests", func(t *testing.T) {
		for _, test := range tests {
			check(t, test)
		}
	})
}

func check(t *testing.T, test one) {

	var src = env.AddImmSource(test.text)
	var lex = new(Lexer)
	lex.Init(src)

	var actual = make([]pair, 0)
	for true {
		_, tok, lit := lex.Scan()
		if tok == EOF {
			break
		}
		actual = append(actual, pair{tok, lit})
	}

	for i := 0; i < len(actual) && i < len(test.pairs); i++ {
		if actual[i].tok != test.pairs[i].tok {
			t.Errorf("Лексема %s вместо %s в тесте:\n'%s'", actual[i].tok, test.pairs[i].tok, test.text)
		} else if actual[i].lit != test.pairs[i].lit {
			t.Errorf("Текст лексемы '%s' вместо '%s' в тесте:\n'%s'", actual[i].lit, test.pairs[i].lit, test.text)
		}
	}

	if len(actual) != len(test.pairs) {
		t.Errorf("Получено %d лексем вместо %d в тексте\n'%s'", len(actual), len(test.pairs), test.text)
	}
}
