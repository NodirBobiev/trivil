package genjava

import (
	"trivil/ast"
	"trivil/env"
	"trivil/jasmin"
)

func (g *genContext) genTypeDecl(td *ast.TypeDecl) {
	switch x := td.GetType().(type) {
	case *ast.ClassType:
		g.genClassType(td, x)
	case *ast.VectorType:
		g.genVectorType(td, x)
	}
}

func (g *genContext) genClassType(td *ast.TypeDecl, t *ast.ClassType) {
	baseType := t.BaseTyp
	var super *jasmin.Class
	if baseType != nil {
		e := g.scope.GetEntityByName(baseType.(*ast.TypeRef).TypeName)
		super = e.(*jasmin.Class)
	}
	g.class = g.pack.CreateClass(env.OutName(td.GetName()), super)
	g.java.Set(g.class)
	g.scope.SetEntity(td, g.class)
	for _, f := range t.Fields {
		g.genField(f)
	}
	g.class = nil
}

func (g *genContext) genVectorType(td *ast.TypeDecl, t *ast.VectorType) {

}
