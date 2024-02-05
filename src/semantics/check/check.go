package check

import (
	"fmt"
	"trivil/ast"
	"trivil/env"
	"trivil/lexer"
)

var _ = fmt.Printf

type checkContext struct {
	module       *ast.Module
	checkedTypes map[string]struct{}
	returnTyp    ast.Type
	errorHint    string // не используется
	loopCount    int
}

func Process(m *ast.Module) {
	var cc = &checkContext{
		module:       m,
		checkedTypes: make(map[string]struct{}),
	}

	for _, d := range m.Decls {
		switch x := d.(type) {
		case *ast.TypeDecl:
			cc.typeDecl(x)
		case *ast.VarDecl:
			cc.varDecl(x)
		case *ast.ConstDecl:
			cc.constDecl(x)
		case *ast.Function:
			// дальше
		default:
			panic(fmt.Sprintf("check: ni %T", d))
		}
	}

	// обойти функции
	for _, d := range m.Decls {
		f, ok := d.(*ast.Function)
		if ok {
			cc.function(f)
		}
	}

	if m.Entry != nil {
		cc.entry(m.Entry)
	}
}

//==== константы и переменные

func (cc *checkContext) varDecl(v *ast.VarDecl) {

	if v.Later {
		if v.Typ == nil {
			env.AddError(v.Pos, "СЕМ-ДЛЯ-ПОЗЖЕ-НУЖЕН-ТИП")
		}
		if v.Host == nil {
			env.AddError(v.Pos, "СЕМ-ПОЗЖЕ-ЛОК-ПЕРЕМЕННАЯ")
		}
	} else {
		cc.expr(v.Init)

		if v.Typ != nil {
			cc.checkAssignable(v.Typ, v.Init)
		} else {
			v.Typ = v.Init.GetType()
			if v.Typ == nil {
				panic("assert - не задан тип переменной " + env.PosString(v.Pos))
			}

			if ast.IsVoidType(v.Typ) {
				env.AddError(v.Pos, "СЕМ-ФН-НЕТ-ЗНАЧЕНИЯ")
				return
			}
			if ast.IsTagPairType(v.Typ) {
				env.AddError(v.Pos, "СЕМ-СОХРАНЕНИЕ-ПОЛИМОРФНОГО")
				return
			}
		}
	}
}

func (cc *checkContext) constDecl(v *ast.ConstDecl) {
	cc.expr(v.Value)

	if v.Typ != nil {
		cc.checkAssignable(v.Typ, v.Value)
	} else {
		v.Typ = v.Value.GetType()
		if v.Typ == nil {
			panic("assert - не задан тип константы")
		}
		if ast.IsVoidType(v.Typ) {
			env.AddError(v.Pos, "СЕМ-ФН-НЕТ-ЗНАЧЕНИЯ")
			return
		}
	}
	cc.checkConstExpr(v.Value)
}

//==== functions

func (cc *checkContext) function(f *ast.Function) {

	var ft = f.Typ.(*ast.FuncType)

	var vPar = ast.VariadicParam(ft)
	if vPar != nil && vPar.Out {
		env.AddError(vPar.Pos, "СЕМ-ВАРИАТИВНЫЙ-ВЫХОДНОЙ")
	}

	if f.Seq != nil {
		cc.returnTyp = ft.ReturnTyp

		cc.loopCount = 0
		cc.statements(f.Seq)

		if cc.returnTyp != nil {
			if !cc.seqHasReturn(f.Seq) {
				env.AddError(f.Pos, "СЕМ-НЕТ-ВЕРНУТЬ")
			}
		}

		cc.returnTyp = nil
	}
}

//==== проверка наличия "вернуть"

func (cc *checkContext) seqHasReturn(s *ast.StatementSeq) bool {
	if len(s.Statements) == 0 {
		return false
	}
	var last = s.Statements[len(s.Statements)-1]
	switch x := last.(type) {
	case *ast.Return:
		return true
	case *ast.Guard:
		return true
	case *ast.Crash:
		return true
	case *ast.If:
		return cc.ifHasReturn(x)
	case *ast.Select:
		return cc.selectHasReturn(x)
	case *ast.SelectType:
		return cc.selectTypeHasReturn(x)
	default:
		return false
	}
}

func (cc *checkContext) ifHasReturn(x *ast.If) bool {
	if !cc.seqHasReturn(x.Then) {
		return false
	}
	if x.Else == nil {
		return false
	}
	if elsif, ok := x.Else.(*ast.If); ok {
		return cc.ifHasReturn(elsif)
	}
	return cc.seqHasReturn(x.Else.(*ast.StatementSeq))
}

func (cc *checkContext) selectHasReturn(x *ast.Select) bool {
	if x.Else == nil || !cc.seqHasReturn(x.Else) {
		return false
	}

	for _, c := range x.Cases {
		if !cc.seqHasReturn(c.Seq) {
			return false
		}
	}

	return true
}

func (cc *checkContext) selectTypeHasReturn(x *ast.SelectType) bool {
	if x.Else == nil || !cc.seqHasReturn(x.Else) {
		return false
	}

	for _, c := range x.Cases {
		if !cc.seqHasReturn(c.Seq) {
			return false
		}
	}

	return true
}

//==== statements

func (cc *checkContext) entry(e *ast.EntryFn) {
	cc.loopCount = 0
	cc.statements(e.Seq)
}

func (cc *checkContext) statements(seq *ast.StatementSeq) {

	for _, s := range seq.Statements {
		cc.statement(s)
	}
}

func (cc *checkContext) statement(s ast.Statement) {
	switch x := s.(type) {
	case *ast.StatementSeq:
		cc.statements(x) // из else
	case *ast.ExprStatement:
		if b, isBinary := x.X.(*ast.BinaryExpr); isBinary && b.Op == lexer.EQ {
			env.AddError(x.X.GetPos(), "СЕМ-ОЖИДАЛОСЬ-ПРИСВАИВАНИЕ")
		} else {
			cc.expr(x.X)
			if _, isCall := x.X.(*ast.CallExpr); !isCall {
				env.AddError(x.Pos, "СЕМ-ЗНАЧЕНИЕ-НЕ-ИСПОЛЬЗУЕТСЯ")
			}
		}
	case *ast.DeclStatement:
		cc.localDecl(x.D)
	case *ast.AssignStatement:
		cc.expr(x.L)
		cc.expr(x.R)
		cc.checkAssignable(x.L.GetType(), x.R)
		cc.checkLValue(x.L)
	case *ast.IncStatement:
		cc.expr(x.L)
		if !ast.IsIntegerType(x.L.GetType()) {
			env.AddError(x.GetPos(), "СЕМ-ОШ-УНАРНАЯ-ТИП",
				ast.TypeString(x.L.GetType()), lexer.INC.String())
		}
		cc.checkLValue(x.L)
	case *ast.DecStatement:
		cc.expr(x.L)
		if !ast.IsIntegerType(x.L.GetType()) {
			env.AddError(x.GetPos(), "СЕМ-ОШ-УНАРНАЯ-ТИП",
				ast.TypeString(x.L.GetType()), lexer.DEC.String())
		}
		cc.checkLValue(x.L)
	case *ast.If:
		cc.expr(x.Cond)
		if !ast.IsBoolType(x.Cond.GetType()) {
			env.AddError(x.Cond.GetPos(), "СЕМ-ТИП-ВЫРАЖЕНИЯ", ast.Bool.Name)
		}

		cc.statements(x.Then)
		if x.Else != nil {
			cc.statement(x.Else)
		}
	case *ast.While:
		cc.checkWhile(x)
	case *ast.Cycle:
		cc.checkCycle(x)
	case *ast.Guard:
		cc.expr(x.Cond)
		if !ast.IsBoolType(x.Cond.GetType()) {
			env.AddError(x.Cond.GetPos(), "СЕМ-ТИП-ВЫРАЖЕНИЯ", ast.Bool.Name)
		}
		cc.statement(x.Else)
		if seq, ok := x.Else.(*ast.StatementSeq); ok {
			if !isTerminating(seq) {
				env.AddError(x.Else.GetPos(), "СЕМ-НЕ-ЗАВЕРШАЮЩИЙ")
			}
		}
	case *ast.Select:
		cc.checkSelect(x)
	case *ast.SelectType:
		cc.checkSelectType(x)

	case *ast.Return:
		if x.X != nil {
			cc.expr(x.X)

			if cc.returnTyp == nil {
				env.AddError(x.Pos, "СЕМ-ОШ-ВЕРНУТЬ-ЛИШНЕЕ")
			} else {
				cc.checkAssignable(cc.returnTyp, x.X)
				x.ReturnTyp = cc.returnTyp
			}
		} else if cc.returnTyp != nil {
			env.AddError(x.Pos, "СЕМ-ОШ-ВЕРНУТЬ-НУЖНО")
		}
	case *ast.Break:
		if cc.loopCount == 0 {
			env.AddError(x.Pos, "СЕМ-ПРЕРВАТЬ-ВНЕ-ЦИКЛА")
		}
	case *ast.Crash:
		cc.expr(x.X)
		if !ast.IsStringType(x.X.GetType()) {
			env.AddError(x.X.GetPos(), "СЕМ-ТИП-ВЫРАЖЕНИЯ", ast.String.Name)
		}

	default:
		panic(fmt.Sprintf("statement: ni %T", s))
	}
}

func (cc *checkContext) localDecl(decl ast.Decl) {

	switch x := decl.(type) {
	case *ast.VarDecl:
		cc.varDecl(x)
	default:
		panic(fmt.Sprintf("local decl: ni %T", decl))
	}
}

//==== циклы

func (cc *checkContext) checkWhile(x *ast.While) {
	cc.expr(x.Cond)
	if !ast.IsBoolType(x.Cond.GetType()) {
		env.AddError(x.Cond.GetPos(), "СЕМ-ТИП-ВЫРАЖЕНИЯ", ast.Bool.Name)
	}
	cc.loopCount++
	cc.statements(x.Seq)
	cc.loopCount--
}

func (cc *checkContext) checkCycle(x *ast.Cycle) {
	cc.expr(x.Expr)
	var t = x.Expr.GetType()

	if !ast.IsIndexableType(t) {
		env.AddError(x.Expr.GetPos(), "СЕМ-ОЖИДАЛСЯ-ТИП-МАССИВА", ast.TypeString(t))
		t = ast.MakeInvalidType(x.Expr.GetPos())
	}

	if x.IndexVar != nil {
		if x.IndexVar.Typ != nil {
			panic("ni - явные типы для итерируемых переменных")
		}
		x.IndexVar.Typ = ast.Int64
	}
	if x.ElementVar != nil {
		if x.ElementVar.Typ != nil {
			panic("ni - явные типы для итерируемых переменных")
		}
		x.ElementVar.Typ = ast.ElementType(t)
	}

	cc.loopCount++
	cc.statements(x.Seq)
	cc.loopCount--
}

//==== оператор выбора по выражению

func (cc *checkContext) checkSelect(x *ast.Select) {

	if x.X == nil {
		cc.checkPredicateSelect(x)
		return
	}

	cc.expr(x.X)
	checkSelectExpr(x.X)

	for _, c := range x.Cases {
		for _, e := range c.Exprs {
			cc.expr(e)
			if !equalTypes(x.X.GetType(), e.GetType()) {
				env.AddError(e.GetPos(), "СЕМ-ВЫБОР-ОШ-ТИП-ВАРИАНТА", ast.TypeName(e.GetType()), ast.TypeName(x.X.GetType()))
			}
		}
		cc.statements(c.Seq)
	}
	if x.Else != nil {
		cc.statements(x.Else)
	}
}

func (cc *checkContext) checkPredicateSelect(x *ast.Select) {

	for _, c := range x.Cases {
		for _, e := range c.Exprs {
			cc.expr(e)
			if !equalTypes(e.GetType(), ast.Bool) {
				env.AddError(e.GetPos(), "СЕМ-ВЫБОР-ОШ-ТИП-ПРЕДИКАТА", ast.TypeName(ast.Bool), ast.TypeName(e.GetType()))
			}
		}
		cc.statements(c.Seq)
	}
	if x.Else != nil {
		cc.statements(x.Else)
	}
}

func checkSelectExpr(x ast.Expr) {
	var t = ast.UnderType(x.GetType())
	switch t {
	case ast.Byte, ast.Int64, ast.Word64, ast.Symbol, ast.String:
		return
	default:
		if ast.IsClassType(t) {
			return
		}
	}
	env.AddError(x.GetPos(), "СЕМ-ВЫБОР-ОШ-ТИП", ast.TypeName(x.GetType()))
}

//==== оператор выбора по типу

func (cc *checkContext) checkSelectType(x *ast.SelectType) {

	cc.expr(x.X)

	tClass, isClass := ast.UnderType(x.X.GetType()).(*ast.ClassType)

	if !isClass {
		env.AddError(x.GetPos(), "СЕМ-ВЫБОР-ТИП-КЛАССА", ast.TypeName(x.X.GetType()))
	}

	for _, c := range x.Cases {
		for _, t := range c.Types {
			caseClass, ok := ast.UnderType(t).(*ast.ClassType)
			if !ok {
				env.AddError(t.GetPos(), "СЕМ-ОЖИДАЛСЯ-ТИП-КЛАССА", ast.TypeName(t))
			} else if isClass {
				if caseClass != tClass && !isDerivedClass(tClass, caseClass) {
					env.AddError(t.GetPos(), "СЕМ-ДОЛЖЕН-БЫТЬ-НАСЛЕДНИКОМ", ast.TypeName(t), ast.TypeName(x.X.GetType()))
				}
			}
		}
		if c.Var != nil {
			if len(c.Types) > 1 {
				env.AddError(c.Pos, "СЕМ-ВЫБОР-ОДИН-ТИП")
			}

			c.Var.Typ = c.Types[0]

		}
		cc.statements(c.Seq)
	}
	if x.Else != nil {
		cc.statements(x.Else)
	}
}

//====

func isTerminating(seq *ast.StatementSeq) bool {
	if len(seq.Statements) == 0 {
		return false
	}
	var st = seq.Statements[len(seq.Statements)-1]
	switch st.(type) {
	case *ast.Return, *ast.Break, *ast.Crash:
		return true
	default:
		return false
	}

}
