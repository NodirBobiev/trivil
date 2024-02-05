package genc

import (
	"fmt"
	"strings"

	"trivil/ast"
)

var _ = fmt.Printf

//==== call

func (genc *genContext) genCall(call *ast.CallExpr) string {

	if call.StdFunc != nil {
		return genc.genStdFuncCall(call)
	}

	if isMethodCall(call.X) {
		return genc.genMethodCall(call)
	}

	var left = genc.genExpr(call.X)

	var cargs = genc.genArgs(call)

	return left + "(" + cargs + ")"
}

func (genc *genContext) genArgs(call *ast.CallExpr) string {

	var ft = call.X.GetType().(*ast.FuncType)

	var cargs = make([]string, len(ft.Params))
	var normLen = len(ft.Params)

	var vPar = ast.VariadicParam(ft)
	if vPar != nil {
		normLen--
	}

	// не вариативные параметры
	for i := 0; i < normLen; i++ {
		var p = ft.Params[i]
		var arg = call.Args[i]
		var expr = genc.genExpr(arg)
		if p.Out {
			cargs[i] = fmt.Sprintf("&(%s)", expr)
		} else {
			var cast = genc.assignCast(p.Typ, arg.GetType())
			cargs[i] = fmt.Sprintf("%s%s", cast, expr)
		}
	}

	if vPar != nil {
		var vTyp = vPar.Typ.(*ast.VariadicType)

		if ast.IsTagPairType(vTyp.ElementTyp) {
			cargs[normLen] = genc.genVariadicTaggedArgs(call, vPar, normLen)
		} else {
			cargs[normLen] = genc.genVariadicArgs(call, vPar, vTyp, normLen)
		}
	}
	return strings.Join(cargs, ", ")
}

func getUnfold(call *ast.CallExpr) *ast.UnfoldExpr {
	if len(call.Args) == 0 {
		return nil
	}
	var last = call.Args[len(call.Args)-1]
	if u, ok := last.(*ast.UnfoldExpr); ok {
		return u
	}
	return nil
}

//TODO: нужно ли выдержать какой-то порядок вычисления аргументов?
func (genc *genContext) genVariadicArgs(call *ast.CallExpr, vPar *ast.Param, vTyp *ast.VariadicType, normCount int) string {

	var unfold = getUnfold(call)
	if unfold != nil {

		if ast.IsVariadicType(unfold.X.GetType()) {
			var vr = genc.genExpr(unfold.X)
			return fmt.Sprintf("%s%s, %s", vr, nm_variadic_len_suffic, vr)
		} else {
			var loc = genc.localName("")
			genc.c("%s %s = %s;", genc.typeRef(unfold.X.GetType()), loc, genc.genExpr(unfold.X))
			return fmt.Sprintf("%s->len, %s->body", loc, loc)
		}
	} else {
		var loc = genc.localName("")
		var et = genc.typeRef(vTyp.ElementTyp)
		var vLen = len(call.Args) - normCount

		var cargs = make([]string, vLen)
		var n = 0
		for i := normCount; i < len(call.Args); i++ {
			cargs[n] = genc.genExpr(call.Args[i])
			n++
		}
		genc.c("%s %s[%d] ={%s};", et, loc, vLen, strings.Join(cargs, ", "))

		return fmt.Sprintf("%d, &%s", vLen, loc)
	}
}

func (genc *genContext) genVariadicTaggedArgs(call *ast.CallExpr, vPar *ast.Param, normCount int) string {

	var unfold = getUnfold(call)
	if unfold != nil {
		var vr = genc.genExpr(unfold.X)
		return fmt.Sprintf("%s%s, %s", vr, nm_variadic_len_suffic, vr)
	} else {
		var loc = genc.localName("")
		var vLen = len(call.Args) - normCount

		//genc.c("struct { TInt64 len; TInt64 body[%d]; } %s;", vLen*2, loc)

		var cargs = make([]string, vLen*2)
		var n = 0
		for i := normCount; i < len(call.Args); i++ {
			cargs[n] = genc.genTypeTag(call.Args[i].GetType())
			cargs[n+1] = genc.castToWord64(call.Args[i])
			n += 2
		}
		genc.c("TWord64 %s[%d] = {%s};", loc, vLen*2, strings.Join(cargs, ", "))

		return fmt.Sprintf("%d, &%s", vLen, loc)
	}
}

func (genc *genContext) castToWord64(e ast.Expr) string {
	return genc.genCastToWord64(genc.genExpr(e), e.GetType())
	/*
		var s = genc.genExpr(e)
		switch ast.UnderType(e.GetType()) {
		case ast.Float64:
			return fmt.Sprintf("((%s)%s).w", rt_cast_union, s)
		case ast.Word64:
			return s
		default:
			return fmt.Sprintf("(%s)%s", predefinedTypeName(ast.Word64.Name), s)
		}
	*/
}

func isMethodCall(left ast.Expr) bool {
	sel, ok := left.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	f, ok := sel.Obj.(*ast.Function)
	if !ok {
		return false
	}

	return f.Recv != nil
}

func (genc *genContext) genMethodCall(call *ast.CallExpr) string {

	sel := call.X.(*ast.SelectorExpr)
	f := sel.Obj.(*ast.Function)

	var name string
	if id, ok := sel.X.(*ast.IdentExpr); ok {
		name = genc.genIdent(id)
	} else {
		name = genc.localName("")

		genc.c("%s %s = %s;",
			genc.typeRef(sel.X.GetType()),
			name,
			genc.genExpr(sel.X))
	}

	//TODO - можно убрать каст, если лишний
	var args = fmt.Sprintf("(%s)%s", genc.typeRef(f.Recv.Typ), name)

	if len(call.Args) > 0 {
		args += ", " + genc.genArgs(call)
	}

	return fmt.Sprintf("%s->%s->%s(%s)", name, nm_VT_field, genc.outName(f.Name), args)
}

func (genc *genContext) genStdFuncCall(call *ast.CallExpr) string {

	switch call.StdFunc.Name {
	case ast.StdLen:
		return genc.genStdLen(call)
	case ast.StdTag:
		return genc.genStdTag(call)
	case ast.StdSomething:
		return genc.genStdSomething(call)

	case ast.VectorAppend:
		return genc.genVectorAppend(call)

	default:
		panic("assert: не реализована стандартная функция " + call.StdFunc.Name)
	}
}

func (genc *genContext) genStdLen(call *ast.CallExpr) string {
	var a = call.Args[0]

	var t = ast.UnderType(a.GetType())
	if t == ast.String {
		return fmt.Sprintf("%s(%s)", rt_lenString, genc.genExpr(a))
	} else if t == ast.String8 {
		return fmt.Sprintf("(%s)->bytes", genc.genExpr(a))
	}

	switch t.(type) {
	case *ast.VectorType:
		return fmt.Sprintf("(%s)->len", genc.genExpr(a))
	case *ast.VariadicType:
		return fmt.Sprintf("%s%s", genc.genExpr(a), nm_variadic_len_suffic)
	default:
		panic("ni")
	}
}

// Запрос длины по выражению expr тип typ
func (genc *genContext) genLen(expr string, typ ast.Type) string {

	var t = ast.UnderType(typ)
	if t == ast.String {
		return fmt.Sprintf("%s(%s)", rt_lenString, expr)
	} else if t == ast.String8 {
		return fmt.Sprintf("(%s)->bytes", expr)
	}

	switch t.(type) {
	case *ast.VectorType:
		return fmt.Sprintf("(%s)->len", expr)
	case *ast.VariadicType:
		return fmt.Sprintf("%s%s", expr, nm_variadic_len_suffic)
	default:
		panic("ni")
	}
}

func (genc *genContext) genStdTag(call *ast.CallExpr) string {

	var a = call.Args[0]

	if tExpr, ok := a.(*ast.TypeExpr); ok {
		return genc.genTypeTag(tExpr.Typ)
	} else {
		return genc.genTagPairTag(a)
	}
}

// Выдает тег по статическому типу
func (genc *genContext) genTypeTag(typ ast.Type) string {
	var t = ast.UnderType(typ)
	switch x := t.(type) {
	case *ast.PredefinedType:
		return fmt.Sprintf("%s%s()", rt_tag, predefinedTypeName(x.Name))
	case *ast.ClassType:
		var tr = ast.DirectTypeRef(typ)
		return fmt.Sprintf("((%s)%s).w",
			rt_cast_union,
			genc.declName(tr.TypeDecl)+nm_class_info_ptr_suffix)
	case *ast.VectorType:
		panic("ni")
	case *ast.MayBeType:
		return genc.genTypeTag(x.Typ)
	default:
		panic(fmt.Sprintf("ni type tag %T", t))
	}
}

// Тег для TagPair выражения
func (genc *genContext) genTagPairTag(e ast.Expr) string {
	switch x := e.(type) {
	case *ast.GeneralBracketExpr:
		if x.Index == nil {
			panic("assert - не может быть композита")
		}

		var left = genc.genExpr(x.X)

		return fmt.Sprintf("((%s*)(%s))[%s(%s, %s%s, %s) << 1]",
			predefinedTypeName(ast.Word64.Name),
			left,
			rt_indexcheck,
			genc.genExpr(x.Index),
			left,
			nm_variadic_len_suffic,
			genPos(x.Pos))

	case *ast.IdentExpr:
		panic("ni - не вариадик параметр '*'")
	}
	panic("assert")
}

func (genc *genContext) genStdSomething(call *ast.CallExpr) string {

	var a = call.Args[0]

	var t = a.GetType()
	if !ast.IsTagPairType(t) {
		panic("assert")
	}

	switch x := a.(type) {
	case *ast.GeneralBracketExpr:
		if x.Index == nil {
			panic("assert")
		}

		var left = genc.genExpr(x.X)

		return fmt.Sprintf("((%s*)(%s))[(%s(%s, %s%s, %s) << 1)+1]",
			predefinedTypeName(ast.Word64.Name),
			left,
			rt_indexcheck,
			genc.genExpr(x.Index),
			left,
			nm_variadic_len_suffic,
			genPos(x.Pos))

	case *ast.IdentExpr:
		panic("ni")
	}
	panic("assert")
}

//==== векторные

func (genc *genContext) genVectorAppend(call *ast.CallExpr) string {

	var vt = ast.UnderType(call.X.GetType()).(*ast.VectorType)
	var et = genc.typeRef(vt.ElementTyp)

	var unfold = getUnfold(call)
	if unfold != nil {

		if ast.IsVariadicType(unfold.X.GetType()) {
			var vr = genc.genExpr(unfold.X)

			return fmt.Sprintf("%s(%s, sizeof(%s), %s%s, %s)",
				rt_vectorAppend,
				genc.genExpr(call.X),
				et,
				vr, nm_variadic_len_suffic,
				vr)

		} else {
			var loc = genc.localName("")
			genc.c("%s %s = %s;", genc.typeRef(unfold.X.GetType()), loc, genc.genExpr(unfold.X))

			var lenName = "len"
			if ast.UnderType(unfold.X.GetType()) == ast.String8 {
				lenName = "bytes"
			}

			return fmt.Sprintf("%s(%s, sizeof(%s), %s->%s, %s->body)",
				rt_vectorAppend,
				genc.genExpr(call.X),
				et,
				loc,
				lenName,
				loc)
		}
	} else {
		var loc = genc.localName("")

		var cargs = make([]string, len(call.Args))
		for i, a := range call.Args {
			var cast = genc.assignCast(vt.ElementTyp, a.GetType())
			cargs[i] = fmt.Sprintf("%s%s", cast, genc.genExpr(a))
		}

		genc.c("%s %s[%d] = {%s};", et, loc, len(call.Args), strings.Join(cargs, ", "))

		return fmt.Sprintf("%s(%s, sizeof(%s), %d, %s)",
			rt_vectorAppend,
			genc.genExpr(call.X),
			et,
			len(call.Args),
			loc)
	}
}
