package check

import (
	"fmt"
	"unicode"

	"trivil/ast"
	"trivil/env"
)

var _ = fmt.Printf

/*
  по целевому типу:
	Байт: Цел64, Слово64, Символ (0..255), Строковый литерал (из 1-го символа)
	Цел64: Байт, Слово64, Вещ64, Символ, Строковый литерал (из 1-го символа)
	Слово64: Байт, Символ, Цел64, Строковый литерал (из 1-го символа)
	Вещ64: Цел64
	Лог: -
	Символ: Байт, Цел64, Слово64, Строковый литерал (из 1-го символа)
	Строка: Символ, []Символ, []Байт
	[]Байт: Строка, Символ
	[]Символ: Строка
	Класс: Класс (от базового к расширенному)
*/
func (cc *checkContext) conversion(x *ast.ConversionExpr) {

	cc.expr(x.X)

	var target = ast.UnderType(x.TargetTyp)

	switch target {
	case ast.Byte:
		cc.conversionToByte(x)
		return
	case ast.Int64:
		cc.conversionToInt64(x)
		return
	case ast.Word64:
		cc.conversionToWord64(x)
		return
	case ast.Float64:
		cc.conversionToFloat64(x)
		return
	case ast.Bool:
		env.AddError(x.Pos, "СЕМ-ОШ-ПРИВЕДЕНИЯ-ТИПА", ast.TypeString(x.X.GetType()), ast.Bool.Name)
		x.Typ = ast.MakeInvalidType(x.Pos)
		return
	case ast.Symbol:
		cc.conversionToSymbol(x)
		return
	case ast.String:
		cc.conversionToString(x)
		return
	case ast.String8:
		cc.conversionToString8(x)
		return
	default:
		switch xt := target.(type) {
		case *ast.VectorType:
			cc.conversionToVector(x, xt)
			return
		case *ast.ClassType:
			cc.conversionToClass(x, xt)
			return
		}
	}

	env.AddError(x.Pos, "СЕМ-ОШ-ПРИВЕДЕНИЯ-ТИПА", ast.TypeName(x.X.GetType()), ast.TypeName(x.TargetTyp))
	x.Typ = ast.MakeInvalidType(x.Pos)
}

func (cc *checkContext) conversionToByte(x *ast.ConversionExpr) {

	var t = ast.UnderType(x.X.GetType())

	switch t {
	case ast.Byte:
		env.AddError(x.Pos, "СЕМ-ПРИВЕДЕНИЕ-ТИПА-К-СЕБЕ", ast.TypeString(x.X.GetType()))
		x.Typ = ast.Byte
		return
	case ast.Int64:
		var li = literal(x.X)
		if li != nil {
			if li.IntVal < 0 || li.IntVal > 255 {
				env.AddError(x.Pos, "СЕМ-ЗНАЧЕНИЕ-НЕ-В-ДИАПАЗОНЕ", ast.Byte.Name)
			} else {
				li.Kind = ast.Lit_Word
				li.WordVal = uint64(li.IntVal)
				li.Typ = ast.Byte
				x.Done = true
			}
		}
		x.Typ = ast.Byte
		return
	case ast.Word64:
		var li = literal(x.X)
		if li != nil {
			if li.WordVal > 255 {
				env.AddError(x.Pos, "СЕМ-ЗНАЧЕНИЕ-НЕ-В-ДИАПАЗОНЕ", ast.Byte.Name)
			} else {
				li.Typ = ast.Byte
				x.Done = true
			}
		}
		x.Typ = ast.Byte
		return
	case ast.Symbol:
		var li = literal(x.X)
		if li != nil {
			if li.WordVal > 255 {
				env.AddError(x.Pos, "СЕМ-ЗНАЧЕНИЕ-НЕ-В-ДИАПАЗОНЕ", ast.Byte.Name)
			} else {
				li.Kind = ast.Lit_Word
				li.Typ = ast.Byte
				x.Done = true
			}
		}
		x.Typ = ast.Byte
		return
	case ast.String:
		var li = literal(x.X)
		if li != nil {
			if len(li.StrVal) != 1 {
				env.AddError(x.Pos, "СЕМ-ДЛИНА-СТРОКИ-НЕ-1")
			} else {
				if li.StrVal[0] > 255 {
					env.AddError(x.Pos, "СЕМ-ЗНАЧЕНИЕ-НЕ-В-ДИАПАЗОНЕ", ast.Byte.Name)
				} else {
					li.Kind = ast.Lit_Word
					li.Typ = ast.Byte
					x.Done = true
				}
			}
			x.Typ = ast.Byte
			return
		}
	}

	env.AddError(x.Pos, "СЕМ-ОШ-ПРИВЕДЕНИЯ-ТИПА", ast.TypeString(x.X.GetType()), ast.Byte.Name)
	x.Typ = ast.MakeInvalidType(x.Pos)
}

func (cc *checkContext) conversionToInt64(x *ast.ConversionExpr) {

	var t = ast.UnderType(x.X.GetType())

	switch t {
	case ast.Int64:
		env.AddError(x.Pos, "СЕМ-ПРИВЕДЕНИЕ-ТИПА-К-СЕБЕ", ast.TypeString(x.X.GetType()))
		x.Typ = ast.Int64
		return
	case ast.Byte:
		var li = literal(x.X)
		if li != nil {
			li.Kind = ast.Lit_Int
			li.IntVal = int64(li.WordVal)
			li.Typ = ast.Int64
			x.Done = true
		}
		x.Typ = ast.Int64
		return
	case ast.Symbol:
		var li = literal(x.X)
		if li != nil {
			li.Kind = ast.Lit_Int
			li.IntVal = int64(li.WordVal)
			li.Typ = ast.Int64
			x.Done = true
		}
		x.Typ = ast.Int64
		return
	case ast.Word64:
		var li = literal(x.X)
		if li != nil {
			if li.WordVal > 1<<63-1 {
				env.AddError(x.Pos, "СЕМ-ЗНАЧЕНИЕ-НЕ-В-ДИАПАЗОНЕ", ast.Int64.Name)
			} else {
				li.Kind = ast.Lit_Int
				li.IntVal = int64(li.WordVal)
				li.Typ = ast.Int64
				x.Done = true
			}
		}
		x.Typ = ast.Int64
		return
	case ast.Float64:
		//TODO: литерал
		x.Typ = ast.Int64
		return
	case ast.String:
		var li = literal(x.X)
		if li != nil {
			if len(li.StrVal) != 1 {
				env.AddError(x.Pos, "СЕМ-ДЛИНА-СТРОКИ-НЕ-1")
			} else {
				li.Kind = ast.Lit_Int
				li.IntVal = int64(li.StrVal[0])
				x.Typ = ast.Int64
				x.Done = true
			}
			return
		}
	}

	env.AddError(x.Pos, "СЕМ-ОШ-ПРИВЕДЕНИЯ-ТИПА", ast.TypeString(x.X.GetType()), ast.Int64.Name)
	x.Typ = ast.MakeInvalidType(x.Pos)
}

func (cc *checkContext) conversionToWord64(x *ast.ConversionExpr) {

	var t = ast.UnderType(x.X.GetType())

	switch t {
	case ast.Word64:
		env.AddError(x.Pos, "СЕМ-ПРИВЕДЕНИЕ-ТИПА-К-СЕБЕ", ast.TypeString(x.X.GetType()))
		x.Typ = ast.Word64
		return
	case ast.Byte:
		var li = literal(x.X)
		if li != nil {
			li.Typ = ast.Word64
			x.Done = true
		}
		x.Typ = ast.Word64
		return
	case ast.Symbol:
		var li = literal(x.X)
		if li != nil {
			li.Typ = ast.Word64
			x.Done = true
		}
		x.Typ = ast.Word64
		return
	case ast.Int64:
		var li = literal(x.X)
		if li != nil {
			if li.IntVal < 0 {
				env.AddError(x.Pos, "СЕМ-ЗНАЧЕНИЕ-НЕ-В-ДИАПАЗОНЕ", ast.Word64.Name)
			} else {
				li.Kind = ast.Lit_Word
				li.WordVal = uint64(li.IntVal)
				li.Typ = ast.Word64
				x.Done = true
			}
		}
		x.Typ = ast.Word64
		return
	case ast.String:
		var li = literal(x.X)
		if li != nil {
			if len(li.StrVal) != 1 {
				env.AddError(x.Pos, "СЕМ-ДЛИНА-СТРОКИ-НЕ-1")
			} else {
				li.Kind = ast.Lit_Word
				li.WordVal = uint64(li.StrVal[0])
				x.Done = true
			}
			x.Typ = ast.Word64
			return
		}
	}

	env.AddError(x.Pos, "СЕМ-ОШ-ПРИВЕДЕНИЯ-ТИПА", ast.TypeString(x.X.GetType()), ast.Word64.Name)
	x.Typ = ast.MakeInvalidType(x.Pos)
}

func (cc *checkContext) conversionToFloat64(x *ast.ConversionExpr) {

	var t = ast.UnderType(x.X.GetType())

	switch t {
	case ast.Float64:
		env.AddError(x.Pos, "СЕМ-ПРИВЕДЕНИЕ-ТИПА-К-СЕБЕ", ast.TypeString(x.X.GetType()))
		x.Typ = ast.Float64
		return
	case ast.Int64:
		// пока не работаю с литералами
		x.Typ = ast.Float64
		return
	}

	env.AddError(x.Pos, "СЕМ-ОШ-ПРИВЕДЕНИЯ-ТИПА", ast.TypeString(x.X.GetType()), ast.Float64.Name)
	x.Typ = ast.MakeInvalidType(x.Pos)

}

func (cc *checkContext) conversionToSymbol(x *ast.ConversionExpr) {

	var t = ast.UnderType(x.X.GetType())

	switch t {
	case ast.Symbol:
		env.AddError(x.Pos, "СЕМ-ПРИВЕДЕНИЕ-ТИПА-К-СЕБЕ", ast.TypeString(x.X.GetType()))
		x.Typ = ast.Symbol
		return
	case ast.Byte:
		var li = literal(x.X)
		if li != nil {
			x.Done = true
			li.Typ = ast.Symbol
		}
		x.Typ = ast.Symbol
		return
	case ast.Int64:
		var li = literal(x.X)
		if li != nil {
			if li.IntVal < 0 || li.IntVal > unicode.MaxRune {
				env.AddError(x.Pos, "СЕМ-ЗНАЧЕНИЕ-НЕ-В-ДИАПАЗОНЕ", ast.Symbol.Name)
			} else {
				li.Kind = ast.Lit_Word
				li.WordVal = uint64(li.IntVal)
				li.Typ = ast.Symbol
				x.Done = true
			}
		}
		x.Typ = ast.Symbol
		return
	case ast.Word64:
		var li = literal(x.X)
		if li != nil {
			if li.WordVal > unicode.MaxRune {
				env.AddError(x.Pos, "СЕМ-ЗНАЧЕНИЕ-НЕ-В-ДИАПАЗОНЕ", ast.Symbol.Name)
			} else {
				li.Typ = ast.Symbol
				x.Done = true
			}
		}
		x.Typ = ast.Symbol
		return
	case ast.String:
		var li = literal(x.X)
		if li != nil {
			if len(li.StrVal) != 1 {
				env.AddError(x.Pos, "СЕМ-ДЛИНА-СТРОКИ-НЕ-1")
			} else {
				li.Kind = ast.Lit_Symbol
				li.WordVal = uint64(li.StrVal[0])
				x.Done = true
			}
			x.Typ = ast.Symbol
			return
		}
	}

	env.AddError(x.Pos, "СЕМ-ОШ-ПРИВЕДЕНИЯ-ТИПА", ast.TypeString(x.X.GetType()), ast.Symbol.Name)
	x.Typ = ast.MakeInvalidType(x.Pos)

}

func (cc *checkContext) conversionToString(x *ast.ConversionExpr) {

	var t = ast.UnderType(x.X.GetType())

	switch t {
	case ast.String:
		env.AddError(x.Pos, "СЕМ-ПРИВЕДЕНИЕ-ТИПА-К-СЕБЕ", ast.TypeString(x.X.GetType()))
		x.Typ = ast.String
		return
	case ast.String8:
		x.Typ = ast.String
		return
	case ast.Symbol:
		var li = literal(x.X)
		if li != nil {
			li.Typ = ast.String
			li.StrVal = make([]rune, 1)
			li.StrVal[0] = rune(li.WordVal)
			x.Done = true
		}
		x.Typ = ast.String
		return
	}

	vt, ok := t.(*ast.VectorType)
	if ok {
		var et = ast.UnderType(vt.ElementTyp)

		if et == ast.Byte || et == ast.Symbol {
			x.Typ = ast.String
			return
		}
	}

	env.AddError(x.Pos, "СЕМ-ОШ-ПРИВЕДЕНИЯ-ТИПА", ast.TypeString(x.X.GetType()), ast.String.Name)
	x.Typ = ast.MakeInvalidType(x.Pos)
}

func (cc *checkContext) conversionToString8(x *ast.ConversionExpr) {

	var t = ast.UnderType(x.X.GetType())

	switch t {
	case ast.String:
		x.Typ = ast.String8
		return
	case ast.String8:
		env.AddError(x.Pos, "СЕМ-ПРИВЕДЕНИЕ-ТИПА-К-СЕБЕ", ast.TypeString(x.X.GetType()))
		x.Typ = ast.String8
		return
	}

	env.AddError(x.Pos, "СЕМ-ОШ-ПРИВЕДЕНИЯ-ТИПА", ast.TypeString(x.X.GetType()), ast.String8.Name)
	x.Typ = ast.MakeInvalidType(x.Pos)
}

func (cc *checkContext) conversionToVector(x *ast.ConversionExpr, target *ast.VectorType) {

	var t = ast.UnderType(x.X.GetType())

	if t == ast.String {
		var et = ast.UnderType(target.ElementTyp)

		if et == ast.Byte || et == ast.Symbol {
			x.Typ = x.TargetTyp
			return
		}
	} else if t == ast.Symbol {
		x.Typ = x.TargetTyp
		return
	}

	env.AddError(x.Pos, "СЕМ-ОШ-ПРИВЕДЕНИЯ-ТИПА",
		ast.TypeString(x.X.GetType()), ast.TypeString(x.TargetTyp))
	x.Typ = ast.MakeInvalidType(x.Pos)

}

func (cc *checkContext) conversionToClass(x *ast.ConversionExpr, target *ast.ClassType) {

	var t = ast.UnderType(x.X.GetType())

	if t == target {
		env.AddError(x.Pos, "СЕМ-ПРИВЕДЕНИЕ-ТИПА-К-СЕБЕ", ast.TypeString(target))
		x.Typ = x.TargetTyp
		return
	}

	tClass, isClass := t.(*ast.ClassType)

	if !isClass {
		if tMayBe, ok := t.(*ast.MayBeType); ok {
			tClass, isClass = ast.UnderType(tMayBe.Typ).(*ast.ClassType)
		}
	}

	if isClass {
		if !isDerivedClass(tClass, target) {
			env.AddError(x.Pos, "СЕМ-ДОЛЖЕН-БЫТЬ-НАСЛЕДНИКОМ", ast.TypeName(x.TargetTyp), ast.TypeName(x.X.GetType()))
		}
		x.Typ = x.TargetTyp
		return
	}

	env.AddError(x.Pos, "СЕМ-ОШ-ПРИВЕДЕНИЯ-ТИПА",
		ast.TypeName(x.X.GetType()), ast.TypeName(x.TargetTyp))
	x.Typ = ast.MakeInvalidType(x.Pos)

}

//====

func literal(expr ast.Expr) *ast.LiteralExpr {

	switch x := expr.(type) {
	case *ast.LiteralExpr:
		return x
	case *ast.ConversionExpr:
		if x.Done {
			return literal(x.X)
		}
	}
	return nil
}

/*!
func oneSymbolString(expr ast.Expr) *ast.LiteralExpr {
	var li = literal(expr)
	if li == nil {
		return nil
	}
	if li.Kind != ast.Lit_String {
		return nil
	}

	if utf8.RuneCountInString(li.Lit) != 1 {
		return nil
	}
	return li
}
*/

func isDerivedClass(base, derived *ast.ClassType) bool {

	var c = derived

	for c.BaseTyp != nil {
		var t = ast.UnderType(c.BaseTyp)
		if t == base {
			return true
		}
		c = t.(*ast.ClassType)
	}
	return false
}
