package genjava

import (
	"trivil/ast"
	"trivil/jasmin"
	"trivil/jasmin/builtins"
)

var generator *genContext

type genContext struct {
	java          *jasmin.Jasmin
	pack          *jasmin.Package
	packMainClass *jasmin.Class
	class         *jasmin.Class
	method        *jasmin.Method
	scope         *Scope
	exprType      jasmin.Type
	mods          map[string]*jasmin.Method
}

func Generate(m *ast.Module, main bool) *jasmin.Jasmin {
	if generator == nil {
		generator = &genContext{
			java:  jasmin.NewJasmin(),
			scope: NewScope(),
			mods:  make(map[string]*jasmin.Method),
		}
		generator.init()
	}
	generator.genModule(m, main)
	return generator.java
}

func (g *genContext) init() {
	g.mods["print_int64"] = builtins.PrintLongMethod()
	g.mods["print_float64"] = builtins.PrintDoubleMethod()
	g.mods["println"] = builtins.PrintlnMethod()
	g.java.Set(builtins.PrintClass())
}
