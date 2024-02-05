package check

import (
	"fmt"
	"trivil/ast"
	"trivil/env"
	"trivil/lexer"
)

var _ = fmt.Printf

func (cc *checkContext) expr(expr ast.Expr) {
	switch x := expr.(type) {
	case *ast.IdentExpr:
		if _, ok := x.Obj.(*ast.TypeRef); ok {
			env.AddError(x.Pos, "СЕМ-ТИП-В-ВЫРАЖЕНИИ")
			x.Typ = ast.MakeInvalidType(x.Pos)
			return
		}

		x.Typ = x.Obj.(ast.Decl).GetType()

		if _, isVar := x.Obj.(*ast.VarDecl); !isVar {
			x.ReadOnly = true
		}
		//fmt.Printf("ident %v %v %v\n", x.Name, v.ReadOnly, x.ReadOnly)

	case *ast.UnaryExpr:
		cc.expr(x.X)
		cc.unaryExpr(x)

	case *ast.BinaryExpr:
		cc.expr(x.X)
		cc.expr(x.Y)
		cc.binaryExpr(x)

	case *ast.OfTypeExpr:
		cc.expr(x.X)
		cc.ofTypeExpr(x)

	case *ast.SelectorExpr:
		cc.selector(x)

	case *ast.CallExpr:
		if x.StdFunc != nil {
			cc.callStdFunction(x)
			return
		}
		cc.call(x)

	case *ast.UnfoldExpr:
		env.AddError(x.Pos, "СЕМ-РАЗВОРАЧИВАНИЕ-ТОЛЬКО-ВАРИАДИК")
		x.Typ = ast.MakeInvalidType(x.Pos)

	case *ast.ConversionExpr:
		if x.Caution {
			cc.cautionCast(x)
		} else {
			cc.conversion(x)
		}

	case *ast.NotNilExpr:
		cc.notNil(x)

	case *ast.GeneralBracketExpr:
		cc.generalBracketExpr(x)

	case *ast.ClassCompositeExpr:
		cc.classComposite(x)

	case *ast.LiteralExpr:
		switch x.Kind {
		case ast.Lit_Int:
			x.Typ = ast.Int64
		case ast.Lit_Word:
			x.Typ = ast.Word64
		case ast.Lit_Float:
			x.Typ = ast.Float64
		case ast.Lit_String:
			x.Typ = ast.String
		case ast.Lit_Symbol:
			x.Typ = ast.Symbol
		default:
			panic(fmt.Sprintf("LiteralExpr - bad kind: ni %v", x))
		}
		x.ReadOnly = true
	case *ast.BoolLiteral:
		x.Typ = ast.Bool
		x.ReadOnly = true
	default:
		panic(fmt.Sprintf("expression: ni %T", expr))
	}

}

func (cc *checkContext) selector(x *ast.SelectorExpr) {
	if x.Obj != nil {
		// imported object
		if _, ok := x.Obj.(*ast.TypeRef); ok {
			env.AddError(x.Pos, "СЕМ-ТИП-В-ВЫРАЖЕНИИ")
			x.Typ = ast.MakeInvalidType(x.Pos)
		} else {
			x.Typ = x.Obj.(ast.Decl).GetType()
		}
		return
	}
	cc.expr(x.X)

	var t = x.X.GetType()

	switch xt := ast.UnderType(t).(type) {
	case *ast.ClassType:
		d, ok := xt.Members[x.Name]
		if !ok {
			env.AddError(x.Pos, "СЕМ-ОЖИДАЛОСЬ-ПОЛЕ-ИЛИ-МЕТОД", x.Name)
		} else {
			if d.GetHost() != cc.module && !d.IsExported() {
				env.AddError(x.Pos, "СЕМ-НЕ-ЭКСПОРТИРОВАН", d.GetName(), d.GetHost().Name)
			}
			x.Typ = d.GetType()
			x.Obj = d
		}
	case *ast.VectorType:
		var m = ast.VectorMethod(x.Name)
		if m == nil {
			env.AddError(x.GetPos(), "СЕМ-НЕ-НАЙДЕН-МЕТОД-ВЕКТОРА", x.Name)
			x.StdMethod = &ast.StdFunction{Method: true}
			x.StdMethod.Name = "" // отметить ошибку
		} else {
			x.StdMethod = m
			// x.Typ - не задан
			return
		}
	default:
		// TODO: выдать отдельную ошибку, если пропущен "^"
		env.AddError(x.GetPos(), "СЕМ-ОЖИДАЛСЯ-ТИП-КЛАССА", ast.TypeName(t))
		x.Typ = ast.MakeInvalidType(x.X.GetPos())
		return
	}

	if x.Typ == nil {
		x.Typ = ast.MakeInvalidType(x.X.GetPos())
	}
}

func (cc *checkContext) notNil(x *ast.NotNilExpr) {
	cc.expr(x.X)

	var t = x.X.GetType()
	maybe, ok := ast.UnderType(t).(*ast.MayBeType)
	if !ok {
		env.AddError(x.Pos, "СЕМ-ОЖИДАЛСЯ-МБ-ТИП", ast.TypeName(t))
		x.Typ = ast.MakeInvalidType(x.Pos)
		return
	}
	x.Typ = maybe.Typ
}

//=== индексация

func looksLikeComposite(c *ast.ArrayCompositeExpr) bool {
	return c.LenExpr != nil || c.CapExpr != nil || c.Default != nil ||
		len(c.Indexes) > 0 ||
		len(c.Values) != 1
}

func (cc *checkContext) generalBracketExpr(x *ast.GeneralBracketExpr) {

	var t = cc.typeExpr(x.X)

	if t != nil || looksLikeComposite(x.Composite) {
		cc.arrayComposite(x.Composite, t)

		if t == nil {
			t = ast.MakeInvalidType(x.X.GetPos())
		}
		x.Typ = t
		x.X = nil
		return
	}

	// это индексация
	cc.expr(x.X)

	t = x.X.GetType()

	if !ast.IsIndexableType(t) {
		env.AddError(x.X.GetPos(), "СЕМ-ОЖИДАЛСЯ-ТИП-МАССИВА", ast.TypeString(t))
		x.Typ = ast.MakeInvalidType(x.Pos)
	} else {
		x.Index = x.Composite.Values[0]
		cc.expr(x.Index)
		if !ast.IsIntegerType(x.Index.GetType()) {
			env.AddError(x.Index.GetPos(), "СЕМ-ОШ-ТИП-ИНДЕКСА", ast.TypeString(x.Index.GetType()))
		}
		x.Typ = ast.ElementType(t)
		if ast.IsTagPairType(x.Typ) {
			x.Typ = ast.TagPairType
		}

		if ast.IsVariadicType(t) {
			x.ReadOnly = true
		}
	}
	x.Composite = nil

	if x.X.IsReadOnly() || ast.UnderType(t) == ast.String8 {
		x.ReadOnly = true
	}
}

func (cc *checkContext) unaryExpr(x *ast.UnaryExpr) {
	switch x.Op {
	case lexer.SUB:
		var t = x.X.GetType()
		if ast.IsInt64(t) || ast.IsWord64(t) || ast.IsFloatType(t) {
			// ok
		} else {
			env.AddError(x.X.GetPos(), "СЕМ-ОШ-УНАРНАЯ-ТИП",
				ast.TypeString(x.X.GetType()), x.Op.String())
		}
		x.Typ = t
	case lexer.BITNOT:
		var t = x.X.GetType()
		if ast.IsIntegerType(t) {
			// ok
		} else {
			env.AddError(x.X.GetPos(), "СЕМ-ОШ-УНАРНАЯ-ТИП",
				ast.TypeString(x.X.GetType()), x.Op.String())
		}
		x.Typ = t
	case lexer.NOT:
		if !ast.IsBoolType(x.X.GetType()) {
			env.AddError(x.X.GetPos(), "СЕМ-ОШ-УНАРНАЯ-ТИП",
				ast.TypeString(x.X.GetType()), x.Op.String())
		}
		x.Typ = ast.Bool

	default:
		panic(fmt.Sprintf("unary expr ni: %T op=%s", x, x.Op.String()))
	}
}

func (cc *checkContext) binaryExpr(x *ast.BinaryExpr) {

	switch x.Op {
	case lexer.ADD, lexer.SUB, lexer.MUL, lexer.REM, lexer.QUO:
		var t = x.X.GetType()
		if ast.IsInt64(t) || ast.IsWord64(t) || ast.IsFloatType(t) {
			checkOperandTypes(x)
		} else {
			addErrorForType(t, x.X.GetPos(), "СЕМ-ОШ-ТИП-ОПЕРАНДА",
				ast.TypeString(t), x.Op.String())
		}
		x.Typ = t
	case lexer.AND, lexer.OR:
		if !ast.IsBoolType(x.X.GetType()) {
			addErrorForType(x.X.GetType(), x.X.GetPos(), "СЕМ-ОШ-ТИП-ОПЕРАНДА",
				ast.TypeString(x.X.GetType()), x.Op.String())
		} else if !ast.IsBoolType(x.Y.GetType()) {
			addErrorForType(x.Y.GetType(), x.Y.GetPos(), "СЕМ-ОШ-ТИП-ОПЕРАНДА",
				ast.TypeString(x.Y.GetType()), x.Op.String())
		}
		x.Typ = ast.Bool

	case lexer.BITAND, lexer.BITOR, lexer.BITXOR:
		var t = x.X.GetType()
		if ast.IsInt64(t) || ast.IsWord64(t) || ast.IsByte(t) {
			checkOperandTypes(x)
		} else {
			addErrorForType(t, x.X.GetPos(), "СЕМ-ОШ-ТИП-ОПЕРАНДА",
				ast.TypeString(t), x.Op.String())
		}
		x.Typ = t

	case lexer.SHL, lexer.SHR:
		var t = x.X.GetType()
		if !ast.IsIntegerType(t) {
			addErrorForType(t, x.X.GetPos(), "СЕМ-ОШ-ТИП-ОПЕРАНДА",
				ast.TypeString(t), x.Op.String())
		}
		var t2 = x.Y.GetType()
		if !ast.IsIntegerType(t2) {
			addErrorForType(t2, x.Y.GetPos(), "СЕМ-ОШ-ТИП-ОПЕРАНДА",
				ast.TypeString(t2), x.Op.String())
		}
		x.Typ = t

	case lexer.EQ, lexer.NEQ:
		var t = ast.UnderType(x.X.GetType())
		if t == ast.Byte || t == ast.Int64 || t == ast.Float64 || t == ast.Word64 ||
			t == ast.Symbol || t == ast.String || t == ast.Bool {
			checkOperandTypes(x)
		} else if ast.IsClassType(t) {
			checkClassOperands(x)
		} else if ast.IsMayBeType(t) {
			checkMayBeOperands(x)
		} else {
			addErrorForType(t, x.Pos, "СЕМ-ОШ-ТИП-ОПЕРАНДА",
				ast.TypeString(x.X.GetType()), x.Op.String())
		}

		x.Typ = ast.Bool
	case lexer.LSS, lexer.LEQ, lexer.GTR, lexer.GEQ:
		var t = ast.UnderType(x.X.GetType())
		if t == ast.Byte || t == ast.Int64 || t == ast.Float64 || t == ast.Word64 || t == ast.Symbol {
			checkOperandTypes(x)
		} else {
			addErrorForType(t, x.Pos, "СЕМ-ОШ-ТИП-ОПЕРАНДА",
				ast.TypeString(x.X.GetType()), x.Op.String())
		}
		x.Typ = ast.Bool

	default:
		panic(fmt.Sprintf("binary expr ni: %T op=%s", x, x.Op.String()))
	}
}

func checkOperandTypes(x *ast.BinaryExpr) {
	if equalTypes(x.X.GetType(), x.Y.GetType()) {
		return
	}
	env.AddError(x.Pos, "СЕМ-ОПЕРАНДЫ-НЕ-СОВМЕСТИМЫ",
		ast.TypeName(x.X.GetType()), x.Op.String(), ast.TypeName(x.Y.GetType()))

}

// Считаю, что "пусто" может быть только вторым операндом
func checkClassOperands(x *ast.BinaryExpr) {

	var l = ast.UnderType(x.X.GetType()).(*ast.ClassType)
	r, ok := ast.UnderType(x.Y.GetType()).(*ast.ClassType)

	if ok {
		if l == r || isDerivedClass(l, r) || isDerivedClass(r, l) {
			return
		}
	}
	env.AddError(x.Pos, "СЕМ-ОПЕРАНДЫ-НЕ-СОВМЕСТИМЫ",
		ast.TypeName(x.X.GetType()), x.Op.String(), ast.TypeName(x.Y.GetType()))
}

// Считаю, что "пусто" может быть только вторым операндом
func checkMayBeOperands(x *ast.BinaryExpr) {

	var l = ast.UnderType(x.X.GetType()).(*ast.MayBeType)
	var r = ast.UnderType(x.Y.GetType())

	if r == ast.NullType {
		return
	} else if rmb, ok := r.(*ast.MayBeType); ok && equalTypes(l.Typ, rmb.Typ) {
		return
	}
	env.AddError(x.Pos, "СЕМ-ОПЕРАНДЫ-НЕ-СОВМЕСТИМЫ",
		ast.TypeName(x.X.GetType()), x.Op.String(), ast.TypeName(x.Y.GetType()))

}

func (cc *checkContext) ofTypeExpr(x *ast.OfTypeExpr) {

	x.Typ = ast.Bool

	var t = x.X.GetType()
	maybe, ok := ast.UnderType(t).(*ast.MayBeType)
	if ok {
		t = maybe.Typ
	}

	cl, ok := ast.UnderType(t).(*ast.ClassType)
	if !ok {
		env.AddError(x.X.GetPos(), "СЕМ-ОПЕРАЦИЯ-ТИПА", ast.TypeName(x.X.GetType()))
		return
	}

	target, ok := ast.UnderType(x.TargetTyp).(*ast.ClassType)
	if !ok {
		env.AddError(x.TargetTyp.GetPos(), "СЕМ-ОПЕРАЦИЯ-ТИПА", ast.TypeName(x.TargetTyp))
		return
	}
	/*
		if t == target {
			env.AddError(x.Pos, "СЕМ-ПРИВЕДЕНИЕ-ТИПА-К-СЕБЕ", ast.TypeString(target))
			return
		}
	*/

	if !isDerivedClass(cl, target) {
		env.AddError(x.Pos, "СЕМ-ДОЛЖЕН-БЫТЬ-НАСЛЕДНИКОМ", ast.TypeName(x.TargetTyp), ast.TypeName(t))
	}

}

func isLValue(expr ast.Expr) bool {

	if expr.IsReadOnly() {
		return false
	}

	switch x := expr.(type) {
	case *ast.IdentExpr:
		if v, isVar := x.Obj.(*ast.VarDecl); isVar {
			return !v.AssignOnce
		} else {
			return !x.ReadOnly
		}
	case *ast.GeneralBracketExpr:
		return x.Index != nil
	case *ast.SelectorExpr:
		if f, ok := x.Obj.(*ast.Field); ok {
			return !f.AssignOnce
		}
		return true
	case *ast.ConversionExpr:
		return isLValue(x.X)
	default:
		return false
	}
}

func (cc *checkContext) checkLValue(expr ast.Expr) {
	if isLValue(expr) {
		return
	}
	env.AddError(expr.GetPos(), "СЕМ-НЕ-ПРИСВОИТЬ")
}
