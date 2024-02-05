package lookup

import (
	"fmt"
	"strconv"

	"trivil/ast"
	"trivil/env"
)

var _ = fmt.Printf

func addToScope(name string, d ast.Decl, scope *ast.Scope) {
	x, ok := scope.Names[name]
	if ok {
		if _, inv := x.(*ast.InvalidDecl); !inv {
			env.AddError(d.GetPos(), "СЕМ-УЖЕ-ОПИСАНО", name)
		}
		return
	}
	scope.Names[name] = d

	//fmt.Printf("scope: %v\n", scope.Names)
}

func findInScopes(scope *ast.Scope, name string, pos int) ast.Decl {

	var cur = scope

	for {
		if cur == nil {
			//ast.ShowScopes("not found "+name, scope)
			return nil
		}

		d, ok := cur.Names[name]
		if ok {
			return d
		}

		cur = cur.Outer
	}
}

// Всегда возвращает объект, возможно InvalidDesc
func lookInScopes(scope *ast.Scope, name string, pos int) ast.Decl {

	var d = findInScopes(scope, name, pos)
	if d != nil {
		return d
	}
	env.AddError(pos, "СЕМ-НЕ-НАЙДЕНО", name)
	var inv = &ast.InvalidDecl{
		DeclBase: ast.DeclBase{Pos: pos, Name: name},
	}
	addToScope(name, inv, scope)
	return inv
}

//==== проверить импорт

func nameForCheckImported(mod *ast.Module, no int) string {
	return strconv.Itoa(no) + mod.Name
}

// Проверяет, что использованный модуль импортирован в исходном файле
func (lc *lookContext) checkImported(mod *ast.Module, pos int) {
	var no = env.SourceNo(pos)
	_, ok := lc.module.Inner.Names[nameForCheckImported(mod, no)]
	if !ok {
		source, _, _ := env.SourcePos(pos)
		env.AddError(pos, "СЕМ-МОДУЛЬ-НЕ-ИМПОРТИРОВАН", mod.Name, source.FileName)
		// не повторять ошибку
		lc.module.Inner.Names[nameForCheckImported(mod, no)] = mod
	}
}
