package genjava

import (
	"fmt"
	"trivil/ast"
	"trivil/jasmin"
	"trivil/jasmin/builtins"
	"trivil/jasmin/core/instruction"
)

var generator *genContext

type genContext struct {
	java         *jasmin.Jasmin
	pack         *jasmin.Package
	class        *jasmin.Class
	method       instruction.S
	locals       int
	stack        int
	scope        *ast.Scope
	classes      map[ast.Decl]*jasmin.Class
	hints        map[ast.Decl]instruction.I
	mods         map[string]*jasmin.Method
	labelCounter int
}

func Generate(m *ast.Module, main bool) *jasmin.Jasmin {
	if generator == nil {
		generator = &genContext{
			java:  jasmin.NewJasmin(),
			scope: NewScope(),
			mods:  make(map[string]*jasmin.Method),
			hints: map[ast.Decl]instruction.I{},
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

func (g *genContext) genLabel(name string) string {
	g.labelCounter++
	return fmt.Sprintf("%s_%d", name, g.labelCounter)
}

func (g *genContext) resetMethod() {
	g.labelCounter = 0
	g.stack = 0
	g.locals = 0
	g.method = nil
}

func (g *genContext) decl(name string) ast.Decl {
	scope := g.scope
	for scope != nil {
		d, ok := scope.Names[name]
		if ok {
			return d
		}
		scope = scope.Outer
	}
	return nil
}
