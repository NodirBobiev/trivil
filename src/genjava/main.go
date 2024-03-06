package genjava

import (
	"fmt"
	"trivil/ast"
	"trivil/jasmin"
)

const (
	moduleMainClass = "MainClass"
)

func (g *genContext) getMainClass() *jasmin.Class {
	return g.packMainClass
}

func (g *genContext) genModule(m *ast.Module, main bool) {
	g.pack = jasmin.NewPackage(m.Name, nil)
	g.packMainClass = g.pack.CreateClass(moduleMainClass, nil)
	g.java.Set(g.packMainClass)
	g.java.Set(g.pack)
	g.scope.SetScope(m.Inner)
	for _, d := range m.Decls {
		switch x := d.(type) {
		case *ast.Function:
			g.genFunction(x)
		case *ast.TypeDecl:
			g.genTypeDecl(x)
		case *ast.VarDecl:
			f := g.getMainClass().CreateField(x.GetName(), g.genType(x.GetType()))
			f.SetStatic(true)
			g.java.Set(f)
			g.scope.SetEntity(x, f)
		}
	}
	if main && m.Entry != nil {
		mainClass := jasmin.MainClass(g.genEntry(m.Entry))
		g.java.Set(mainClass)
	}
}

func (g *genContext) genTypeDecl(t *ast.TypeDecl) {
	//var accessFlag jasmin.AccessFlag
	//if t.Exported {
	//	accessFlag = jasmin.Public
	//} else {
	//	accessFlag = jasmin.Private
	//}
	baseType := t.GetType().(*ast.ClassType).BaseTyp
	var super *jasmin.Class
	if baseType != nil {
		e := g.scope.GetEntityByName(baseType.(*ast.TypeRef).TypeName)
		super = e.(*jasmin.Class)
	}
	g.class = g.pack.CreateClass(t.Name, super)
	g.java.Set(g.class)
	g.scope.SetEntity(t, g.class)
	switch x := t.GetType().(type) {
	case *ast.ClassType:
		for _, f := range x.Fields {
			g.genField(f)
		}
		for _, f := range x.Methods {
			g.genFunction(f)
		}
	}

	g.class = nil
}

func (g *genContext) genField(f *ast.Field) {
	var accessFlag jasmin.AccessFlag
	if f.Exported {
		accessFlag = jasmin.Public
	} else {
		accessFlag = jasmin.Protected
	}
	ff := g.class.CreateField(f.GetName(), g.genType(f.GetType()))
	ff.SetAccessFlag(accessFlag)
	g.java.Set(ff)
	g.scope.SetEntity(f, ff)
	// TODO: Initialize field with value
}

func (g *genContext) genEntry(f *ast.EntryFn) *jasmin.Method {
	g.class = g.getMainClass()
	g.method = jasmin.MainMethod(g.class)
	g.class.Set(g.method)
	g.java.Set(g.method)
	g.genStatementSeq(f.Seq)
	method := g.method
	g.method = nil
	return method
}

func (g *genContext) genFunction(f *ast.Function) {
	if f.Mod != nil {
		g.genExternalFunction(f)
		return
	}
	var (
		isStatic   bool
		accessFlag jasmin.AccessFlag
	)
	if f.Recv != nil {
		g.class = g.pack.GetClass(g.getClassName(f.Recv.GetType()))
	} else {
		g.class = g.getMainClass()
		isStatic = true
	}
	accessFlag = jasmin.Protected
	if f.Exported {
		accessFlag = jasmin.Public
	}
	g.method = jasmin.NewMethod(f.GetName(), g.class)
	//g.method = g.class.CreateMethod(f.GetName())
	g.method.SetType(g.genType(f.Typ))
	g.method.SetAccessFlag(accessFlag)
	g.method.SetStatic(isStatic)
	g.class.Set(g.method)
	g.java.Set(g.method)
	// if the function has a receiver, then the receiver is variable this
	if f.Recv != nil {
		varDecl := f.Seq.Inner.Names[f.Recv.GetName()]
		this := jasmin.NewVariable(f.Recv.GetName(), g.method, g.genType(varDecl.GetType()), 0)
		g.scope.SetEntity(varDecl, this)
	}
	// assign local var numbers to parameters
	t := f.Typ.(*ast.FuncType)
	for _, x := range t.Params {
		varDecl := f.Seq.Inner.Names[x.GetName()]
		paramVar := g.method.AssignNumber(jasmin.NewVariable(x.GetName(), g.method, g.genType(varDecl.GetType()), -1))
		g.scope.SetEntity(varDecl, paramVar)
	}
	g.scope.SetEntity(f, g.method)
	g.genStatementSeq(f.Seq)

	g.method = nil
}

func (g *genContext) genExternalFunction(f *ast.Function) {
	mod, ok := f.Mod.Attrs["имя"]
	if !ok {
		panic(fmt.Sprintf("имя modifier wasn't found"))
	}
	e := g.mods[mod]
	g.scope.SetEntity(f, e)
}

func (g *genContext) genStatementSeq(s *ast.StatementSeq) {

	for _, i := range s.Statements {
		g.genStatement(i)
	}
}

func (g *genContext) genStatement(s ast.Statement) {
	switch x := s.(type) {
	case *ast.DeclStatement:
		g.genLocalDecl(x.D)
	case *ast.AssignStatement:
		g.genAssignStatement(x)
	case *ast.Return:
		g.genReturn(x)
	case *ast.ExprStatement:
		g.genExprStatement(x)
	default:
		panic(fmt.Sprintf("unexpected statements: %+v", s))
	}
}

func (g *genContext) genLocalDecl(d ast.Decl) {
	switch x := d.(type) {
	case *ast.VarDecl:
		t := g.genType(x.GetType())
		e := g.method.AssignNumber(jasmin.NewVariable(x.GetName(), g.method, t, -1))
		g.scope.SetEntity(d, e)
		g.method.Append(g.genExpr(x.Init)...)
		g.method.Append(jasmin.Store(e.Number, t))
	default:
		panic("unknown local decl")
	}
}

func (g *genContext) genAssignStatement(a *ast.AssignStatement) {
	load, store := g.genAssignExprLeft(a.L)
	g.method.Append(load...)
	g.method.Append(g.genExpr(a.R)...)
	g.method.Append(store)
}

func (g *genContext) genReturn(e *ast.Return) {
	instructions := g.genExpr(e.X)
	g.method.Append(append(instructions, jasmin.Return(g.genType(e.ReturnTyp)))...)
}

func (g *genContext) genExprStatement(e *ast.ExprStatement) {
	g.method.Append(g.genExpr(e.X)...)
}
