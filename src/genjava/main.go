package genjava

import (
	"fmt"
	"trivil/ast"
	"trivil/env"
	"trivil/jasmin"
)

const (
	moduleMainClass = "MainClass"
)

func (g *genContext) getMainClass() *jasmin.Class {
	return g.packMainClass
}

func (g *genContext) genModule(m *ast.Module, main bool) {
	g.pack = jasmin.NewPackage(env.OutName(m.GetName()), nil)
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

func (g *genContext) genField(f *ast.Field) {
	var accessFlag jasmin.AccessFlag
	if f.Exported {
		accessFlag = jasmin.Public
	} else {
		accessFlag = jasmin.Protected
	}
	ff := g.class.CreateField(env.OutName(f.GetName()), g.genType(f.GetType()))
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
	g.method = jasmin.NewMethod(env.OutName(f.GetName()), g.class)
	//g.method = g.class.CreateMethod(f.GetName())
	g.method.SetType(g.genType(f.Typ))
	g.method.SetAccessFlag(accessFlag)
	g.method.SetStatic(isStatic)
	g.class.Set(g.method)
	g.java.Set(g.method)
	// if the function has a receiver, then the receiver is variable this
	if f.Recv != nil {
		varDecl := f.Seq.Inner.Names[f.Recv.GetName()]
		this := jasmin.NewVariable(env.OutName(f.Recv.GetName()), g.method, g.genType(varDecl.GetType()), 0)
		g.scope.SetEntity(varDecl, this)
	}
	// assign local var numbers to parameters
	t := f.Typ.(*ast.FuncType)
	for _, x := range t.Params {
		varDecl := f.Seq.Inner.Names[x.GetName()]
		paramVar := g.method.AssignNumber(jasmin.NewVariable(env.OutName(x.GetName()), g.method, g.genType(varDecl.GetType()), -1))
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
	case *ast.If:
		g.genIf(x)
	case *ast.While:
		g.genWhile(x)
	case *ast.StatementSeq:
		g.genStatementSeq(x)
	case *ast.Break:
		g.genBreak(x)
	default:
		panic(fmt.Sprintf("unexpected statements: %+v", s))
	}
}

func (g *genContext) genLocalDecl(d ast.Decl) {
	switch x := d.(type) {
	case *ast.VarDecl:
		t := g.genType(x.GetType())
		e := g.method.AssignNumber(jasmin.NewVariable(env.OutName(x.GetName()), g.method, t, -1))
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
	if e.X == nil {
		g.method.Append(jasmin.Return(jasmin.NewVoidType()))
	} else {
		instructions := g.genExpr(e.X)
		g.method.Append(append(instructions, jasmin.Return(g.genType(e.ReturnTyp)))...)
	}
}

func (g *genContext) genExprStatement(e *ast.ExprStatement) {
	g.method.Append(g.genExpr(e.X)...)
}

func (g *genContext) genIf(s *ast.If) {
	gen := func(i *ast.If, nextLabel, skipLabel string) {
		bin, _ := i.Cond.(*ast.BinaryExpr)
		t := g.genType(bin.X.GetType())
		eq := negate(bin.Op)
		x := append(append(g.genExpr(bin.X), g.genExpr(bin.Y)...))
		x = append(x, jasmin.Cmp(t), jasmin.If(eq, nextLabel))
		g.method.Append(x...)

		g.genStatementSeq(s.Then)
		if skipLabel != "" {
			g.method.Append(jasmin.Goto(skipLabel))
		}
	}
	var (
		endIfLabel  = g.genLabel("END_IF")
		nextIfLabel string
	)

	for {
		if s.Else != nil {
			nextIfLabel = g.genLabel("ELSE_IF")
		} else {
			nextIfLabel = endIfLabel
			endIfLabel = ""
		}
		gen(s, nextIfLabel, endIfLabel)
		g.method.Append(jasmin.NewLabel(nextIfLabel))
		if s.Else == nil {
			break
		}
		if nextIf, ok := s.Else.(*ast.If); ok {
			s = nextIf
		} else {
			g.genStatement(s.Else)
			g.method.Append(jasmin.NewLabel(endIfLabel))
			break
		}
	}
	if s.Else == nil {
		nextIfLabel = ""
	}
}

func (g *genContext) genWhile(s *ast.While) {
	var (
		whileStartLabel = g.genLabel("WHILE_START")
		whileEndLabel   = g.genLabel("WHILE_END")
	)
	g.cyclesLabels = append(g.cyclesLabels, whileEndLabel)
	defer func() { g.cyclesLabels = g.cyclesLabels[:len(g.cyclesLabels)-1] }()
	g.method.Append(jasmin.NewLabel(whileStartLabel))

	bin, _ := s.Cond.(*ast.BinaryExpr)

	g.method.Append(g.genExpr(bin.X)...)
	g.method.Append(g.genExpr(bin.Y)...)
	g.method.Append(
		jasmin.Cmp(g.genType(bin.X.GetType())),
		jasmin.If(negate(bin.Op), whileEndLabel),
	)

	g.genStatementSeq(s.Seq)

	g.method.Append(jasmin.Goto(whileStartLabel))

	g.method.Append(jasmin.NewLabel(whileEndLabel))
}

func (g *genContext) genBreak(_ *ast.Break) {
	g.method.Append(jasmin.Goto(g.cyclesLabels[len(g.cyclesLabels)-1]))
}
