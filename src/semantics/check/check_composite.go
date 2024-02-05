package check

import (
	"fmt"
	"trivil/ast"
	"trivil/env"
)

var _ = fmt.Printf

func (cc *checkContext) typeExpr(expr ast.Expr) ast.Type {

	switch x := expr.(type) {
	case *ast.IdentExpr:
		if tr, ok := x.Obj.(*ast.TypeRef); ok {
			return tr
		} else {
			return nil
		}
	case *ast.SelectorExpr:
		if tr, ok := x.Obj.(*ast.TypeRef); ok {
			return tr
		} else {
			return nil
		}
	}

	return nil
}

//==== конструктор вектора

func (cc *checkContext) arrayComposite(c *ast.ArrayCompositeExpr, t ast.Type) {

	var elemT ast.Type = nil

	if t == nil {
		env.AddError(c.Pos, "СЕМ-КОНСТРУКТОР-НЕТ-ТИПА")
	} else if !ast.IsIndexableType(t) {
		env.AddError(c.Pos, "СЕМ-КОН-ВЕКТОРА-ОШ-ТИП")
	} else {
		c.Typ = t
		elemT = ast.ElementType(t)
	}

	if c.LenExpr != nil {
		cc.expr(c.LenExpr)
		cc.checkAssignable(ast.Int64, c.LenExpr)
	}

	if c.CapExpr != nil {
		cc.expr(c.CapExpr)
		cc.checkAssignable(ast.Int64, c.CapExpr)
	}

	if c.Default != nil {
		cc.expr(c.Default)
		if elemT != nil {
			cc.checkAssignable(elemT, c.Default)
		}
	}

	for _, inx := range c.Indexes {
		cc.expr(inx)
		cc.checkAssignable(ast.Int64, inx)
		cc.checkConstExpr(inx)
	}

	for _, val := range c.Values {
		cc.expr(val)
		if elemT != nil {
			cc.checkAssignable(elemT, val)
		}
	}

	if c.Default == nil && (c.LenExpr != nil || len(c.Indexes) > 0) {
		// TODO: если длина не задана явно, можно проверить, что индексы без дырок
		env.AddError(c.Pos, "СЕМ-КОН-ВЕКТОРА-НЕТ-УМОЛЧАНИЯ")
	}

	cc.arrayCompositeIndexes(c)
}

func (cc *checkContext) arrayCompositeIndexes(c *ast.ArrayCompositeExpr) {

	// если были ошибки, не пытаюсь проверить индексы и длину
	if env.ErrorCount() > 0 {
		return
	}

	if len(c.Indexes) == 0 {
		return
	}

	if c.LenExpr != nil && cc.isConstExpr(c.LenExpr) {
		c.Length = cc.calculateIntConstExpr(c.LenExpr)
		if c.Length < 0 {
			env.AddError(c.LenExpr.GetPos(), "СЕМ-КОН-ВЕКТОРА-ОШ-ДЛИНА")
		}
	}

	var inxMap = make(map[int64]int)

	var max int64 = -1
	for _, inx := range c.Indexes {
		index := cc.calculateIntConstExpr(inx)

		pos, ok := inxMap[index]
		if ok {
			env.AddError(inx.GetPos(), "СЕМ-КОН-ВЕКТОРА-ИНДЕКС-ДУБЛЬ", index, env.PosString(pos))
		} else {
			inxMap[index] = inx.GetPos()
		}

		if index > max {
			max = index
		}
		if index < 0 || c.Length >= 0 && index >= c.Length {
			env.AddError(inx.GetPos(), "СЕМ-КОН-ВЕКТОРА-ИНДЕКС-ДИАП", c.Length, index)
		}

	}

	c.MaxIndex = max
}

//==== конструктор класса

func (cc *checkContext) classComposite(c *ast.ClassCompositeExpr) {

	var t = cc.typeExpr(c.X)

	if t == nil {
		env.AddError(c.Pos, "СЕМ-КОНСТРУКТОР-НЕТ-ТИПА")
		c.Typ = ast.MakeInvalidType(c.X.GetPos())
		return
	}

	cl, isClass := ast.UnderType(t).(*ast.ClassType)
	if !isClass {
		env.AddError(c.Pos, "СЕМ-КЛАСС-КОМПОЗИТ-ОШ-ТИП")
		c.Typ = ast.MakeInvalidType(c.X.GetPos())
	} else {
		c.Typ = t
	}

	for _, vp := range c.Values {
		cc.expr(vp.Value)
	}

	if !isClass {
		return
	}

	// проверяю поля и типы
	var vals = make(map[string]bool)
	for i, vp := range c.Values {
		d, ok := cl.Members[vp.Name]
		if !ok {
			env.AddError(vp.Pos, "СЕМ-КЛАСС-КОМПОЗИТ-НЕТ-ПОЛЯ", vp.Name)
		} else {
			f, ok := d.(*ast.Field)
			if !ok {
				env.AddError(vp.Pos, "СЕМ-КЛАСС-КОМПОЗИТ-НЕ-ПОЛE")
			} else if f.Host != cc.module && !f.Exported {
				env.AddError(vp.Pos, "СЕМ-НЕ-ЭКСПОРТИРОВАН", f.Name, f.Host.Name)
			} else {
				vals[vp.Name] = true
				c.Values[i].Field = f
				cc.checkAssignable(f.Typ, vp.Value)
			}
		}
	}
	// проверяю позднюю инициализацию
	for name, d := range cl.Members {
		if f, ok := d.(*ast.Field); ok && f.Later {
			_, ok := vals[name]
			if !ok {
				env.AddError(c.Pos, "СЕМ-НЕТ-ПОЗЖЕ-ПОЛЯ", name)
			}
		}
	}
}
