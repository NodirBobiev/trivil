package compiler

import (
	"fmt"
	//"path"
	"os"
	"path/filepath"
	"strings"

	"trivil/ast"
	//	"trivil/env"
)

var _ = fmt.Printf

type defContext struct {
	lines []string
}

func (dc *defContext) f(format string, args ...interface{}) {
	dc.lines = append(dc.lines, fmt.Sprintf(format, args...))
}

func makeDef(m *ast.Module, folder string) {

	var dc = &defContext{
		lines: make([]string, 0),
	}

	ast.CurHost = m

	dc.f("модуль %v\n", m.Name)

	for _, d := range m.Decls {
		x, ok := d.(*ast.TypeDecl)
		if ok && x.Exported {
			dc.typeDecl(x)
		}
	}

	for _, d := range m.Decls {
		x, ok := d.(*ast.Function)
		if ok && x.Exported {
			dc.function(x, false)
		}
	}

	dc.f("")

	//fmt.Println(dc.lines)
	writeFile(folder, m.Name, ".def", dc.lines)

	ast.CurHost = nil
}

func writeFile(folder, name, ext string, lines []string) {

	var filename = filepath.Join(folder, name+ext)

	var out = strings.Join(lines, "\n")

	var err = os.WriteFile(filename, []byte(out), 0644)

	if err != nil {
		panic("Ошибка записи файла " + filename + ": " + err.Error())
	}
	fmt.Printf("записано %v\n", filename)
}

/*
func (dc *defContext) decl(d ast.Decl) {
	switch x := d.(type) {
	case *ast.Function:
		dc.function(x)
	default:
		// ничего
	}
}
*/

func (dc *defContext) function(x *ast.Function, methods bool) {
	var recv = ""
	if x.Recv != nil {
		if !methods {
			return
		}
		recv = fmt.Sprintf("(%v: %v) ", x.Recv.Name, dc.typeName(x.Recv.Typ))
	}

	var ft = x.Typ.(*ast.FuncType)

	var resTyp = ""
	if ft.ReturnTyp != nil {
		resTyp = fmt.Sprintf(": %v", dc.typeName(ft.ReturnTyp))
	}

	var params = make([]string, 0)
	for _, p := range ft.Params {
		params = append(params, fmt.Sprintf("%v: %v", p.Name, dc.typeName(p.Typ)))
	}

	dc.f("фн %v%v(%v)%v", recv, x.Name, strings.Join(params, ", "), resTyp)
}

//==== типы

func (dc *defContext) typeName(t ast.Type) string {
	return ast.TypeName(t)
}

func (dc *defContext) typeDecl(td *ast.TypeDecl) {
	switch xt := ast.UnderType(td.Typ).(type) {
	case *ast.ClassType:
		dc.classDecl(td, xt)
	default:
		dc.f("тип %v = %v", td.Name, dc.typeName(td.Typ))
	}
}

func (dc *defContext) classDecl(td *ast.TypeDecl, x *ast.ClassType) {
	dc.f("тип %v = класс{}", td.Name)

	for _, m := range x.Methods {
		dc.function(m, true)
	}
	dc.f("")
}
