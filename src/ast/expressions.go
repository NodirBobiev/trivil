package ast

import (
	"fmt"
	"trivil/lexer"
)

var _ = fmt.Printf

//====

type ExprBase struct {
	Pos      int
	Typ      Type
	ReadOnly bool
}

func (n *ExprBase) ExprNode() {}

func (n *ExprBase) GetPos() int {
	return n.Pos
}

func (n *ExprBase) GetType() Type {
	return n.Typ
}

func (n *ExprBase) SetType(t Type) {
	n.Typ = t
}

func (n *ExprBase) IsReadOnly() bool {
	return n.ReadOnly
}

//====

type InvalidExpr struct {
	ExprBase
}

//====

type BinaryExpr struct {
	ExprBase
	X  Expr
	Op lexer.Token
	Y  Expr
}

type UnaryExpr struct {
	ExprBase
	Op lexer.Token
	X  Expr
}

type OfTypeExpr struct {
	ExprBase
	X         Expr
	TargetTyp Type
}

type LiteralKind int

const (
	Lit_Int = iota
	Lit_Word
	Lit_Float
	Lit_Symbol
	Lit_String
	Lit_Null
)

type LiteralExpr struct {
	ExprBase
	Kind     LiteralKind
	IntVal   int64  // Цел
	WordVal  uint64 // Байт, Слово, Символ
	FloatStr string // Вещ, чтобы не терять точность
	StrVal   []rune // Строка
}

type BoolLiteral struct {
	ExprBase
	Value bool
}

type IdentExpr struct {
	ExprBase
	Name string
	Obj  Node // Decl: Var, Const or Function or TypeRef
}

type SelectorExpr struct {
	ExprBase
	X         Expr // == nil, если импортированный объект
	Name      string
	Obj       Node // импортированный объект или поле или метод
	StdMethod *StdFunction
}

type CallExpr struct {
	ExprBase
	X       Expr
	Args    []Expr
	StdFunc *StdFunction
}

type UnfoldExpr struct {
	ExprBase
	X Expr
}

type ConversionExpr struct {
	ExprBase
	X         Expr
	TargetTyp Type
	Caution   bool
	Done      bool // X уже преобразован к целевому типу
}

// Если тип передается, как параметр, например, в функции 'тег'
type TypeExpr struct {
	ExprBase
}

type NotNilExpr struct {
	ExprBase
	X Expr
}

//==== index

type GeneralBracketExpr struct {
	ExprBase
	X         Expr
	Index     Expr // indexation if != nil, otherwise composite
	Composite *ArrayCompositeExpr
}

type ArrayCompositeExpr struct {
	ExprBase
	LenExpr  Expr  // если константое выражение, то значение сохраняется в Length
	Length   int64 // если вычислено, или -1
	CapExpr  Expr
	Default  Expr // default value
	Indexes  []Expr
	MaxIndex int64 // из значение индекса, -1, если нет индексов
	Values   []Expr
}

//=== class composite

type ClassCompositeExpr struct {
	ExprBase
	X      Expr
	Values []ValuePair
}

type ValuePair struct {
	Pos   int
	Name  string
	Field *Field
	Value Expr
}
