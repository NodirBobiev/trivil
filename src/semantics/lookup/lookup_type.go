package lookup

import (
	"fmt"
	"trivil/ast"
	"trivil/env"
)

var _ = fmt.Printf

//==== ссылка на тип

func (lc *lookContext) lookTypeRef(t ast.Type) {

	if maybe, ok := t.(*ast.MayBeType); ok {
		lc.lookTypeRef(maybe.Typ)

		if !ast.IsReferenceType(maybe.Typ) {
			env.AddError(maybe.Typ.GetPos(), "СЕМ-МБ-ТИП-НЕ-ССЫЛКА", ast.TypeName(maybe.Typ))
		}
		return
	}

	var tr, ok = t.(*ast.TypeRef)
	if !ok {
		if t == nil {
			panic("assert")
		}
		vTyp, ok := t.(*ast.VariadicType)
		if ok {
			lc.lookTypeRef(vTyp.ElementTyp)
		}
		return
	}

	if tr.Typ != nil {
		return // уже сделано
	}

	var td *ast.TypeDecl

	if tr.ModuleName != "" {
		td = lc.lookTypeDeclInModule(tr.ModuleName, tr.TypeName, tr.Pos)
	} else {
		td = lc.lookTypeDeclInScopes(tr.TypeName, tr.Pos)
	}

	tr.TypeDecl = td
	tr.Typ = tr.TypeDecl.Typ

	if tr.Typ == nil {
		panic("not resolved")
	}

	//fmt.Printf("! %v %T\n", tr.TypeDecl, tr.Typ)
}

func (lc *lookContext) lookTypeDeclInScopes(name string, pos int) *ast.TypeDecl {

	var d = findInScopes(lc.scope, name, pos)

	if d == nil {
		env.AddError(pos, "СЕМ-НЕ-НАЙДЕНО", name)
		var td = lc.makeTypeDecl(name, pos)
		addToScope(name, td, lc.scope)
		return td
	}

	td, ok := d.(*ast.TypeDecl)
	if !ok {
		env.AddError(pos, "СЕМ-ДОЛЖЕН-БЫТЬ-ТИП", name)
		return lc.makeTypeDecl(name, pos)
	}

	return td
}

func (lc *lookContext) lookTypeDeclInModule(moduleName, name string, pos int) *ast.TypeDecl {
	var d = findInScopes(lc.scope, moduleName, pos)

	if d == nil {
		env.AddError(pos, "СЕМ-НЕ-НАЙДЕН-МОДУЛЬ", moduleName)
		return lc.makeTypeDecl(name, pos)
	}

	m, ok := d.(*ast.Module)
	if !ok {
		env.AddError(pos, "СЕМ-ДОЛЖЕН-БЫТЬ-МОДУЛЬ", moduleName)
		return lc.makeTypeDecl(name, pos)
	}

	lc.checkImported(m, pos)

	d, ok = m.Inner.Names[name]
	if !ok {
		env.AddError(pos, "СЕМ-НЕ-НАЙДЕНО-В-МОДУЛЕ", m.Name, name)
		var td = lc.makeTypeDecl(name, pos)
		addToScope(name, td, lc.scope)
		return td
	}

	td, ok := d.(*ast.TypeDecl)
	if !ok {
		env.AddError(pos, "СЕМ-ДОЛЖЕН-БЫТЬ-ТИП", name)
		return lc.makeTypeDecl(name, pos)
	}

	if !d.IsExported() {
		env.AddError(pos, "СЕМ-НЕ-ЭКСПОРТИРОВАН", name, m.Name)
	}

	return td
}

func (lc *lookContext) makeTypeDecl(name string, pos int) *ast.TypeDecl {
	var td = &ast.TypeDecl{
		DeclBase: ast.DeclBase{
			Pos:      pos,
			Name:     name,
			Typ:      ast.MakeInvalidType(pos),
			Host:     lc.module,
			Exported: true,
		},
	}
	return td
}

//==== типы

func (lc *lookContext) lookTypeDecl(v *ast.TypeDecl) {

	switch x := v.Typ.(type) {
	case *ast.VectorType:
		lc.lookTypeRef(x.ElementTyp)
		// проверяю на []мб Т - это самая длинная в Тривиле цепочка,
		// если будут анонимные типы вектора, надо переделывать
		var t = x.ElementTyp
		if maybe, ok := t.(*ast.MayBeType); ok {
			lc.lookTypeRef(maybe.Typ)
			t = maybe.Typ
		}
		if !ast.IsClassType(t) {
			lc.checkRecursion(t)
		}

	case *ast.ClassType:
		if x.BaseTyp != nil {
			lc.lookTypeRef(x.BaseTyp)
			lc.checkRecursion(x.BaseTyp)
		}
		for _, f := range x.Fields {
			if f.Typ != nil {
				lc.lookTypeRef(f.Typ)
			}
			if !f.Later {
				lc.lookExpr(f.Init)
			}
		}

	case *ast.MayBeType:
		lc.lookTypeRef(x.Typ)
		lc.checkRecursion(x.Typ)

		if !ast.IsReferenceType(ast.UnderType(x.Typ)) {
			env.AddError(x.Typ.GetPos(), "СЕМ-МБ-ТИП-НЕ-ССЫЛКА", ast.TypeName(x.Typ))
		}

	case *ast.TypeRef:
		lc.lookTypeRef(x)
		var td = x.TypeDecl

		completed, exist := lc.processed[td]
		if exist {
			if !completed {
				env.AddError(td.GetPos(), "СЕМ-РЕКУРСИВНОЕ-ОПРЕДЕЛЕНИЕ", td.GetName())
			}
		} else if td.GetHost() == lc.module {
			lc.lookDecl(td)
		}

	case *ast.InvalidType:
	default:
		panic(fmt.Sprintf("lookTypeDecl: ni %T", v.Typ))
	}
}

func (lc *lookContext) checkRecursion(t ast.Type) {

	tr, ok := t.(*ast.TypeRef)
	if !ok {
		return
	}
	tr = ast.DirectTypeRef(tr)
	var td = tr.TypeDecl
	if td.GetHost() == lc.module {
		lc.lookDecl(td)
	}
}
