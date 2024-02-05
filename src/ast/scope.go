package ast

import (
	"fmt"
)

var _ = fmt.Printf

var (
	Byte *PredefinedType
	//Int32   *PredefinedType
	Int64   *PredefinedType
	Word64  *PredefinedType
	Float64 *PredefinedType
	Bool    *PredefinedType
	Symbol  *PredefinedType
	String  *PredefinedType
	String8 *PredefinedType

	NullType *PredefinedType

	VoidType    *PredefinedType // только для вызова функции без результата
	TagPairType *PredefinedType // только в типе параметров и в типе элемента variadic
)

type Scope struct {
	Outer *Scope
	Names map[string]Decl
}

// Стандартные функции
const (
	StdLen       = "длина"
	StdTag       = "тег"
	StdSomething = "нечто"
)

// Методы
const (
	VectorAppend   = "добавить"
	VectorAllocate = "выделить"
	//VectorFill     = "заполнить"
)

var topScope *Scope // верхняя область видимости

var vectorMethods map[string]*StdFunction

func initScopes() {
	topScope = &Scope{
		Names: make(map[string]Decl),
	}

	Byte = addType("Байт")
	//	Int32 = addType("Цел32")
	Int64 = addType("Цел64")
	Float64 = addType("Вещ64")
	Word64 = addType("Слово64")
	Bool = addType("Лог")
	Symbol = addType("Символ")
	String = addType("Строка")
	String8 = addType("Строка8")
	NullType = addType("Пусто")

	VoidType = &PredefinedType{Name: "нет результата"}
	TagPairType = &PredefinedType{Name: "ТегСлово"}

	addBoolConst("истина", true)
	addBoolConst("ложь", false)

	addNullConst("пусто")

	addStdFunction(StdLen)

	addStdFunction(StdTag)
	addStdFunction(StdSomething)

	vectorMethods = make(map[string]*StdFunction)
	addVectorMethod(VectorAppend)
	//addVectorMethod(VectorFill)

	//	ShowScopes("top", topScope)
}

func addType(name string) *PredefinedType {
	var pt = &PredefinedType{
		Name: name,
	}

	var td = &TypeDecl{}
	td.Typ = pt
	td.Name = name
	topScope.Names[name] = td

	return pt
}

func addBoolConst(name string, val bool) {
	var c = &ConstDecl{
		Value: &BoolLiteral{Value: val},
	}
	c.Typ = Bool
	c.Name = name

	topScope.Names[name] = c
}

func addNullConst(name string) {
	var c = &ConstDecl{
		Value: &LiteralExpr{
			ExprBase: ExprBase{Typ: NullType},
			Kind:     Lit_Null,
		},
	}
	c.Typ = NullType
	c.Name = name

	topScope.Names[name] = c
}

func addStdFunction(name string) {
	var f = &StdFunction{}
	f.Typ = VoidType
	f.Name = name

	topScope.Names[name] = f
}

func addVectorMethod(name string) {
	var f = &StdFunction{Method: true}
	f.Typ = VoidType
	f.Name = name

	vectorMethods[name] = f
}

func NewScope(outer *Scope) *Scope {
	return &Scope{
		Outer: outer,
		Names: make(map[string]Decl),
	}
}

func ShowScopes(label string, cur *Scope) {
	if label != "" {
		fmt.Println(label)
	}

	var n = 0
	for cur != nil {
		n++
		fmt.Printf("<--- scope %d\n", n)
		for _, d := range cur.Names {
			fmt.Printf("%s ", d.GetName())
		}
		if len(cur.Names) > 0 {
			fmt.Println()
		}
		cur = cur.Outer
	}
	fmt.Printf("--- end scopes\n")
}

//=== методы для всех векторов

func VectorMethod(name string) *StdFunction {
	f, ok := vectorMethods[name]
	if ok {
		return f
	}
	return nil
}
