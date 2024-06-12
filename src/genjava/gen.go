package genjava

import (
	"fmt"
	"trivil/ast"
	"trivil/jasmin"
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
	labelCounter  int
	cyclesLabels  []string
}

func Generate(m *ast.Module, main bool) *jasmin.Jasmin {
	if generator == nil {
		generator = &genContext{
			java:         jasmin.NewJasmin(),
			scope:        NewScope(),
			mods:         make(map[string]*jasmin.Method),
			cyclesLabels: make([]string, 0),
		}
		generator.initMods()
	}
	generator.genModule(m, main)
	return generator.java
}

func (g *genContext) initMods() {
	g.mods["print_int64"] = jasmin.NewVoidPublicStaticMethod("builtins/Print/print_int64", jasmin.NewLongType())
	g.mods["print_float64"] = jasmin.NewVoidPublicStaticMethod("builtins/Print/print_float64", jasmin.NewDoubleType())
	g.mods["print_string"] = jasmin.NewVoidPublicStaticMethod("builtins/Print/print_string", jasmin.NewStringType())
	g.mods["println"] = jasmin.NewPublicStaticMethod("builtins/Print/println", jasmin.VoidMethodType())

	g.mods["scan_int64"] = jasmin.NewNullaryPublicStaticMethod("builtins/Scan/scan_int64", jasmin.NewLongType())
	g.mods["scan_float64"] = jasmin.NewNullaryPublicStaticMethod("builtins/Scan/scan_float64", jasmin.NewDoubleType())
	g.mods["scan_string"] = jasmin.NewNullaryPublicStaticMethod("builtins/Scan/scan_string", jasmin.NewStringType())
	g.mods["scanln"] = jasmin.NewNullaryPublicStaticMethod("builtins/Scan/scanln", jasmin.NewStringType())
}

func (g *genContext) genLabel(name string) string {
	g.labelCounter++
	return fmt.Sprintf("%s_%d", name, g.labelCounter)
}

func (g *genContext) resetMethod() {
	g.labelCounter = 0
}
