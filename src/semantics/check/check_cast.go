package check

import (
	"fmt"
	//"strconv"
	//"unicode"
	//"unicode/utf8"

	"trivil/ast"
	"trivil/env"
)

var _ = fmt.Printf

/*
  по целевому типу:
	Слово64: Цел64, Вещ64, ссылочный (вектор, класс)
	Цел64: Слово64
	Вещ64: Слово64
	ссылочный: Слово64
*/
func (cc *checkContext) cautionCast(x *ast.ConversionExpr) {

	cc.expr(x.X)

	var target = ast.UnderType(x.TargetTyp)
	var from = ast.UnderType(x.X.GetType())

	switch target {
	case ast.Int64:
		if from == ast.Word64 {
			x.Typ = x.TargetTyp
		} else {
			env.AddError(x.Pos, "СЕМ-ОШ-ОСТОРОЖНОГО-ПРИВЕДЕНИЯ", ast.TypeName(x.X.GetType()), ast.Int64.Name)
			x.Typ = ast.MakeInvalidType(x.Pos)
		}
		return
	case ast.Float64:
		if from == ast.Word64 {
			x.Typ = x.TargetTyp
		} else {
			env.AddError(x.Pos, "СЕМ-ОШ-ОСТОРОЖНОГО-ПРИВЕДЕНИЯ", ast.TypeName(x.X.GetType()), ast.Float64.Name)
			x.Typ = ast.MakeInvalidType(x.Pos)
		}
		return
	case ast.Word64:
		if from == ast.Int64 || from == ast.Float64 || ast.IsReferenceType(from) {
			x.Typ = x.TargetTyp
		} else {
			env.AddError(x.Pos, "СЕМ-ОШ-ОСТОРОЖНОГО-ПРИВЕДЕНИЯ", ast.TypeName(x.X.GetType()), ast.Word64.Name)
			x.Typ = ast.MakeInvalidType(x.Pos)
		}
	default:
		if from == ast.Word64 && ast.IsReferenceType(target) {
			x.Typ = x.TargetTyp
		} else {
			env.AddError(x.Pos, "СЕМ-ОШ-ОСТОРОЖНОГО-ПРИВЕДЕНИЯ", ast.TypeName(x.X.GetType()), ast.TypeName(x.TargetTyp))
			x.Typ = ast.MakeInvalidType(x.Pos)
		}
	}
}
