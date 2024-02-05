package lexer

import (
	"strconv"
	//"unicode"
)

type Token int

const (
	// Special tokens
	Invalid Token = iota
	EOF
	NL

	LINE_COMMENT
	BLOCK_COMMENT
	MODIFIER

	// literals
	IDENT
	INT
	FLOAT
	STRING
	SYMBOL

	// operators
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	AND // &
	OR  // |
	NOT // ~

	BITAND // :&
	BITOR  // :|
	BITXOR // :\
	BITNOT // :~

	SHL // <<
	SHR // >>

	INC // ++
	DEC // --

	EQ  // =
	LSS // <
	GTR // >

	NEQ // #
	LEQ // <=
	GEQ // >=

	NOTNIL // ^

	ELLIPSIS // ...
	ASSIGN   // :=

	LPAR   // (
	RPAR   // )
	LBRACK // [
	RBRACK // ]
	LBRACE // {
	RBRACE // }
	LCONV  // (:
	//RCONV  // ›

	COMMA // ,
	DOT   // .
	SEMI  // ;
	COLON // :

	// keywords
	keyword_beg
	AMONG
	BREAK
	CAUTION
	CLASS
	CONST
	CRASH
	CYCLE
	ELSE
	ENTRY
	FN
	GUARD
	IF
	IMPORT
	LATER
	MAYBE
	MODULE
	OFTYPE
	OTHER
	RETURN
	SELECT
	TYPE
	VAR
	WHEN
	WHILE
	keyword_end
)

var tokens = [...]string{
	Invalid: "Invalid",

	EOF: "EOF",
	NL:  "NL",

	LINE_COMMENT:  "LINE_COMMENT",
	BLOCK_COMMENT: "BLOCK_COMMENT",

	MODIFIER: "@",

	IDENT:  "идентификатор",
	INT:    "целый литерал",
	FLOAT:  "вещественный литерал",
	STRING: "строковый литерал",
	SYMBOL: "символьный литерал",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",

	AND: "&",
	OR:  "|",
	NOT: "~",
	//	XOR:     "^",

	BITAND: ":&",
	BITOR:  ":|",
	BITXOR: ":\\",
	BITNOT: ":~",

	SHL: "<<",
	SHR: ">>",

	INC: "++",
	DEC: "--",

	EQ:  "=",
	LSS: "<",
	GTR: ">",

	NEQ: "#",
	LEQ: "<=",
	GEQ: ">=",

	NOTNIL: "^",

	ELLIPSIS: "...",
	ASSIGN:   ":=",

	LPAR: "(",
	RPAR: ")",

	LBRACK: "[",
	RBRACK: "]",

	LBRACE: "{",
	RBRACE: "}",

	LCONV: "(:", // or .<
	//RCONV: "›",

	COMMA: ",",
	DOT:   ".",
	SEMI:  ";",
	COLON: ":",

	CRASH:   "авария",
	RETURN:  "вернуть",
	ENTRY:   "вход",
	SELECT:  "выбор",
	OTHER:   "другое",
	IF:      "если",
	ELSE:    "иначе",
	IMPORT:  "импорт",
	CLASS:   "класс",
	WHEN:    "когда",
	CONST:   "конст",
	MAYBE:   "мб",
	MODULE:  "модуль",
	GUARD:   "надо",
	CAUTION: "осторожно",
	WHILE:   "пока",
	LATER:   "позже",
	BREAK:   "прервать",
	VAR:     "пусть",
	AMONG:   "среди",
	TYPE:    "тип",
	OFTYPE:  "типа",
	FN:      "фн",
	CYCLE:   "цикл",
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token, keyword_end-(keyword_beg+1))
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
func Lookup(ident string) Token {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}
	return IDENT
}
