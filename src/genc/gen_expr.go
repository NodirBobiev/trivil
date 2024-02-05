package genc

import (
	"fmt"
	"strings"
	"unicode"

	"trivil/ast"
	"trivil/lexer"
)

var _ = fmt.Printf

func (genc *genContext) genExpr(expr ast.Expr) string {

	switch x := expr.(type) {
	case *ast.IdentExpr:
		return genc.genIdent(x)
	case *ast.LiteralExpr:
		return genc.genLiteral(x)
	case *ast.UnaryExpr:
		return fmt.Sprintf("%s(%s)", unaryOp(x.Op), genc.genExpr(x.X))
	case *ast.BinaryExpr:
		return genc.genBinaryExpr(x)
	case *ast.OfTypeExpr:
		return genc.genOfTypeExpr(x)
	case *ast.SelectorExpr:
		return genc.genSelector(x)
	case *ast.CallExpr:
		return genc.genCall(x)
	case *ast.ConversionExpr:
		if x.Caution {
			return genc.genCautionCast(x)
		} else {
			return genc.genConversion(x)
		}
	case *ast.NotNilExpr:
		return genc.genNotNil(x)

	case *ast.GeneralBracketExpr:
		return genc.genBracketExpr(x)

	case *ast.ClassCompositeExpr:
		return genc.genClassComposite(x)

	default:
		panic(fmt.Sprintf("gen expression: ni %T", expr))
	}
}

func (genc *genContext) genIdent(id *ast.IdentExpr) string {

	var d = id.Obj.(ast.Decl)
	var s = genConstAsLiteral(d)
	if s != "" {
		return s
	}

	if isVarOutParam(d) {
		return fmt.Sprintf("(*%s)", genc.declName(d))
	}

	return genc.declName(d)
}

func isVarOutParam(d ast.Decl) bool {
	v, ok := d.(*ast.VarDecl)
	if !ok {
		return false
	}
	return v.OutParam
}

func genConstAsLiteral(d ast.Decl) string {

	c, ok := d.(*ast.ConstDecl)
	if !ok {
		return ""
	}
	switch x := c.Value.(type) {
	case *ast.BoolLiteral:
		if x.Value {
			return "true"
		} else {
			return "false"
		}
	case *ast.LiteralExpr:
		if x.Typ == ast.NullType {
			return "NULL"
		}
	}
	return ""
}

//==== literals

func (genc *genContext) genLiteral(li *ast.LiteralExpr) string {
	switch li.Kind {
	case ast.Lit_Int:
		return fmt.Sprintf("%d", li.IntVal)
	case ast.Lit_Word:
		return fmt.Sprintf("0x%X", li.WordVal)
	case ast.Lit_Float:
		return li.FloatStr
	case ast.Lit_Symbol:
		return fmt.Sprintf("0x%X", li.WordVal)
	case ast.Lit_String:
		return genc.genStringLiteral(li)
	default:
		panic("ni")
	}
}

func (genc *genContext) genStringLiteral(li *ast.LiteralExpr) string {

	if len(li.StrVal) == 0 {
		return fmt.Sprintf("%s()", rt_emptyString)
	}

	var name = genc.localName(nm_stringLiteral)
	genc.g("static TString %s = NULL;", name)

	var outs = encodeLiteralString(li.StrVal)
	//fmt.Printf("! байты=%d  символы=%d\n", len(outs), len(li.StrVal))

	// передаю -1 в число байтов, чтобы не учитывать Си эскейп последовательности
	return fmt.Sprintf("%s(&%s, %d, %d, \"%s\")", rt_newLiteralString, name, -1, len(li.StrVal), outs)
}

func encodeLiteralString(runes []rune) string {
	var b = strings.Builder{}
	for _, r := range runes {
		switch r {
		case '\a':
			b.WriteString("\\a")
		case '\b':
			b.WriteString("\\b")
		case '\f':
			b.WriteString("\\f")
		case '\n':
			b.WriteString("\\n")
		case '\r':
			b.WriteString("\\r")
		case '\t':
			b.WriteString("\\t")
		case '\v':
			b.WriteString("\\v")
		case '\\':
			b.WriteString("\\\\")
		case '"':
			b.WriteString("\\\"")
		default:
			if unicode.IsControl(r) {
				panic("ni") // можно \0xddd
			}
			b.WriteRune(r)
		}
	}
	return b.String()
}

//==== унарные операции

func unaryOp(op lexer.Token) string {
	switch op {
	case lexer.SUB:
		return "-"
	case lexer.NOT:
		return "!"
	case lexer.BITNOT:
		return "~"

	default:
		panic("ni unary" + op.String())
	}
}

//==== бинарные операции

func binaryOp(op lexer.Token) string {
	switch op {
	case lexer.OR:
		return "||"
	case lexer.AND:
		return "&&"
	case lexer.EQ:
		return "=="
	case lexer.NEQ:
		return "!="
	case lexer.LSS:
		return "<"
	case lexer.LEQ:
		return "<="
	case lexer.GTR:
		return ">"
	case lexer.GEQ:
		return ">="
	case lexer.ADD:
		return "+"
	case lexer.SUB:
		return "-"
	case lexer.BITOR:
		return "|"
	case lexer.MUL:
		return "*"
	case lexer.QUO:
		return "/"
	case lexer.REM:
		return "%"
	case lexer.BITAND:
		return "&"
	case lexer.BITXOR:
		return "^"
	case lexer.SHL:
		return "<<"
	case lexer.SHR:
		return ">>"
	default:
		panic("ni binary" + op.String())
	}
}

func (genc *genContext) genBinaryExpr(x *ast.BinaryExpr) string {

	if ast.IsStringType(x.X.GetType()) {
		var not = ""
		if x.Op == lexer.NEQ {
			not = "!"
		}

		return fmt.Sprintf("%s%s(%s, %s)", not, rt_equalStrings, genc.genExpr(x.X), genc.genExpr(x.Y))
	}

	if ast.IsClassType(x.X.GetType()) {
		return fmt.Sprintf("((void*)%s %s (void*)%s)", genc.genExpr(x.X), binaryOp(x.Op), genc.genExpr(x.Y))
	} else {
		return fmt.Sprintf("(%s %s %s)", genc.genExpr(x.X), binaryOp(x.Op), genc.genExpr(x.Y))
	}
}

func (genc *genContext) genOfTypeExpr(x *ast.OfTypeExpr) string {
	var tname = genc.typeRef(x.TargetTyp)

	return fmt.Sprintf("%s(%s, %s)",
		rt_isClassType,
		genc.genExpr(x.X),
		tname+nm_class_info_ptr_suffix)

}

//==== selector

func (genc *genContext) genSelector(x *ast.SelectorExpr) string {
	if x.X == nil {
		return genc.declName(x.Obj.(ast.Decl))
	}

	var cl = ast.UnderType(x.X.GetType()).(*ast.ClassType)
	return fmt.Sprintf("(%s)->%s.%s%s",
		genc.genExpr(x.X),
		nm_class_fields,
		pathToField(cl, x.Name),
		genc.outName(x.Name))
}

func pathToField(cl *ast.ClassType, name string) string {
	var path = ""
	for {
		if cl.BaseTyp == nil {
			break
		}
		cl = ast.UnderType(cl.BaseTyp).(*ast.ClassType)

		_, ok := cl.Members[name]
		if !ok {
			break
		}
		path += nm_base_fields + "."
	}
	return path
}

//==== проверка на nil

func (genc *genContext) genNotNil(x *ast.NotNilExpr) string {
	return fmt.Sprintf("(%s)%s(%s,%s)",
		genc.typeRef(x.Typ),
		rt_nilcheck,
		genc.genExpr(x.X),
		genPos(x.Pos))
}

//==== индексация и композит массива

func (genc *genContext) genBracketExpr(x *ast.GeneralBracketExpr) string {

	if x.Index != nil {

		switch xt := ast.UnderType(x.X.GetType()).(type) {
		case *ast.VectorType:
			return genc.genVectorIndex(x.X, x.Index, "len")
		case *ast.VariadicType:
			return genc.genVariadicIndex(xt, x.X, x.Index)
		default:
			if xt == ast.String8 {
				return genc.genVectorIndex(x.X, x.Index, "bytes")
			}
		}
		panic("assert")
	}

	return genc.genArrayComposite(x.Composite)
}

func (genc *genContext) genVectorIndex(x, inx ast.Expr, lenName string) string {
	var name string
	if id, ok := x.(*ast.IdentExpr); ok {
		name = genc.genIdent(id)
	} else {
		name = genc.localName("")

		genc.c("%s %s = %s;",
			genc.typeRef(x.GetType()),
			name,
			genc.genExpr(x))
	}
	return fmt.Sprintf("%s->body[%s(%s, %s->%s,%s)]",
		name,
		rt_indexcheck,
		genc.genExpr(inx),
		name,
		lenName,
		genPos(inx.GetPos()))

}

func (genc *genContext) genVariadicIndex(vt *ast.VariadicType, x, inx ast.Expr) string {

	var vPar = genc.genExpr(x)

	if ast.IsTagPairType(vt.ElementTyp) {
		/*
			return fmt.Sprintf("((%s*)(%s + sizeof(TInt64)))[%s(%s, *(TInt64 *)%s) << 1]",
				predefinedTypeName(ast.Int64.Name),
				vPar,
				rt_indexcheck,
				genc.genExpr(inx),
				vPar)
		*/
		panic("assert")
	} else {

		return fmt.Sprintf("((%s*)%s)[%s(%s, %s%s, %s)]",
			genc.typeRef(vt.ElementTyp),
			vPar,
			rt_indexcheck,
			genc.genExpr(inx),
			vPar,
			nm_variadic_len_suffic,
			genPos(inx.GetPos()))
	}
}

//==== конструктор вектора

func (genc *genContext) genArrayComposite(x *ast.ArrayCompositeExpr) string {
	var vt = ast.UnderType(x.Typ).(*ast.VectorType)

	var name = genc.localName("")
	var lenExpr = genc.genArrayCompositeLen(x)

	// нужно ли проверять макс индекс на < длина?
	if len(x.Indexes) > 0 && x.LenExpr != nil {
		// c.Length = maxIndex + 1
		var lenName = genc.localName("len")
		genc.c("%s %s = %s;", predefinedTypeName(ast.Int64.Name), lenName, lenExpr)
		lenExpr = lenName

		genc.c("%s(%d, %s, %s);", rt_indexcheck, x.MaxIndex, lenName, genPos(x.Pos))
	}

	var s = ""

	if x.Default != nil {
		s = fmt.Sprintf("%s %s = %s(sizeof(%s), %s, %s, %s);",
			genc.typeRef(x.Typ),
			name,
			rt_newVectorFill,
			genc.typeRef(vt.ElementTyp),
			lenExpr,
			genc.genArrayCompositeCap(x),
			genc.genFiller(x.Default))
	} else {
		s = fmt.Sprintf("%s %s = %s(sizeof(%s), %s, %s);",
			genc.typeRef(x.Typ),
			name,
			rt_newVector,
			genc.typeRef(vt.ElementTyp),
			lenExpr,
			genc.genArrayCompositeCap(x))
	}

	var list = make([]string, len(x.Values))
	for i, val := range x.Values {
		var inx string
		if len(x.Indexes) == 0 {
			inx = fmt.Sprintf("%d", i)
		} else {
			inx = genc.genExpr(x.Indexes[i])
		}
		list[i] = fmt.Sprintf("%s->body[%s] = %s;", name, inx, genc.genExpr(val))
	}
	s += strings.Join(list, " ")

	genc.c("%s", s)

	return name
}

func (genc *genContext) genArrayCompositeLen(x *ast.ArrayCompositeExpr) string {
	// если длина задана явно и это константое выражение
	if x.Length >= 0 {
		return fmt.Sprintf("%d", x.Length)
	}

	// если длина задана явно
	if x.LenExpr != nil {
		return genc.genExpr(x.LenExpr)
	}

	// если длина задана неявно по индексам
	if x.MaxIndex >= 0 {
		return fmt.Sprintf("%d", x.MaxIndex+1)
	}

	// если последовательность значений (без индекса)
	return fmt.Sprintf("%d", len(x.Values))
}

func (genc *genContext) genArrayCompositeCap(x *ast.ArrayCompositeExpr) string {
	if x.CapExpr != nil {
		return genc.genExpr(x.CapExpr)
	}
	return "0"
}

func (genc *genContext) genFiller(x ast.Expr) string {
	var expr = genc.genExpr(x)

	return genc.genCastToWord64(expr, x.GetType())
}

//==== конструктор класса

func (genc *genContext) genClassComposite(x *ast.ClassCompositeExpr) string {
	var name = genc.localName("")

	var tname = genc.typeRef(x.Typ)
	var s = fmt.Sprintf("%s %s = %s(%s);",
		tname,
		name,
		rt_newObject,
		tname+nm_class_info_ptr_suffix)

	var cl = ast.UnderType(x.Typ).(*ast.ClassType)

	var list = make([]string, len(x.Values))
	for i, v := range x.Values {
		var cast = genc.assignCast(v.Field.Typ, v.Value.GetType())
		list[i] = fmt.Sprintf("%s->%s.%s%s = %s%s;",
			name, nm_class_fields, pathToField(cl, v.Name), genc.outName(v.Name),
			cast, genc.genExpr(v.Value))
	}
	s += strings.Join(list, " ")

	genc.c("%s", s)
	return name
}
