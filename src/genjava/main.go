package genjava

import (
	"fmt"
	"trivil/ast"
	"trivil/env"
	"trivil/jasmin"
	"trivil/jasmin/core/instruction"
	"trivil/jasmin/core/tps"
)

const (
	moduleMainClass = "MainClass"
)

func (g *genContext) mainClass() *jasmin.Class {
	return g.pack.Classes["Main"]
}

func (g *genContext) genModule(m *ast.Module, main bool) {
	g.pack = jasmin.NewPackage(env.OutName(m.GetName()))
	g.java.Packages[g.pack.Name] = g.pack
	g.scope.SetScope(m.Inner)
	for _, d := range m.Decls {
		switch x := d.(type) {
		case *ast.Function:
			g.genFunction(x)
		case *ast.TypeDecl:
			g.genTypeDecl(x)
		case *ast.VarDecl:
			g.genGlobalVarDecl(x)
		}
	}
	if main && m.Entry != nil {
		g.genEntry(m.Entry)
	}
}

func (g *genContext) genGlobalVarDecl(v *ast.VarDecl) {
	var (
		access = "protected"
	)
	if v.Exported {
		access = "public"
	}
	g.mainClass().Fields.Append(
		instruction.Field(access, true, env.OutName(v.GetName()), g.genType(v.GetType())),
	)
}

func (g *genContext) genTypeDecl(t *ast.TypeDecl) {
	var (
		access    = "public"
		classDesc = g.pack.Name + "/" + env.OutName(t.GetName())
		super     = "java/lang/Object"
	)

	baseType := t.GetType().(*ast.ClassType).BaseTyp
	if baseType != nil {
		d := g.decl(baseType.(*ast.TypeRef).TypeName)
		super = g.classes[d].Name
	}
	g.class = &jasmin.Class{
		Name: classDesc,
		Directives: []instruction.I{
			instruction.Class(access, classDesc),
			instruction.Super(super),
		},
		Fields:  make(instruction.S, 0),
		Methods: make([]instruction.S, 0),
	}
	g.classes[t] = g.class
	g.pack.Classes[t.GetName()] = g.class

	for _, f := range t.GetType().(*ast.ClassType).Fields {
		g.genField(f)
	}

	g.class = nil
}

func (g *genContext) genField(f *ast.Field) {
	var (
		access    = "protected"
		fieldName = env.OutName(f.GetName())
		fieldType = g.genType(f.GetType())
		field     = g.class.Name + "/" + fieldName
	)
	if f.Exported {
		access = "public"
	}
	g.class.Fields.Append(
		instruction.Field(access, false, fieldName, fieldType),
	)
	g.hints[f] = instruction.Getfield(field, fieldType.String())
	// TODO: Initialize field with value
}

func (g *genContext) genEntry(f *ast.EntryFn) {
	g.class = g.mainClass()
	g.genMethod(f.Seq, "public", true, "main"+tps.MainMethod.String())
}

func (g *genContext) genMethod(seq *ast.StatementSeq, access string, isStatic bool, fnNameAndType string) {
	g.method = instruction.S{}
	g.genStatementSeq(seq)
	g.method.Prepend(
		instruction.Method(access, isStatic, fnNameAndType),
		instruction.LimitStack(g.stack),
		instruction.LimitLocals(g.locals),
	).Append(instruction.EndMethod())
	g.class.Methods = append(g.class.Methods, g.method)
	g.resetMethod()
}

func (g *genContext) genFunction(f *ast.Function) {
	if f.Mod != nil {
		g.genExternalFunction(f)
		return
	}
	var (
		isStatic bool
		access   = "protected"
		fnName   = env.OutName(f.GetName())
		fnType   = g.genType(f.GetType())
		fnDesc   = g.class.Name + "/" + fnName + fnType.String()
	)
	if f.Recv != nil {
		g.class = g.pack.Classes[g.getClassName(f.Recv.GetType())]
	} else {
		g.class = g.mainClass()
		isStatic = true
	}
	if f.Exported {
		access = "public"
	}

	// if the function has a receiver, then the receiver is "this" (0 as local number).
	if f.Recv != nil {
		varDecl := f.Seq.Inner.Names[f.Recv.GetName()]
		g.hints[varDecl] = instruction.Aload(0)
		g.locals++
	}
	// assign local var numbers to parameters
	t := f.Typ.(*ast.FuncType)
	for _, x := range t.Params {
		varDecl := f.Seq.Inner.Names[x.GetName()]
		varType := g.genType(varDecl.GetType())
		g.hints[varDecl] = loadInstruction(varType, g.locals)
		g.locals += typeSize(varType)
	}

	if isStatic {
		g.hints[f] = instruction.Invokestatic(fnDesc)
	} else {
		g.hints[f] = instruction.Invokevirtual(fnDesc)
	}

	g.genMethod(f.Seq, access, isStatic, fnName+fnType.String())
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
	case *ast.StatementSeq:
		g.genStatementSeq(x)
	default:
		panic(fmt.Sprintf("unexpected statements: %+v", s))
	}
}

func (g *genContext) genLocalDecl(d ast.Decl) {
	varDecl := d.(*ast.VarDecl)
	varType := g.genType(varDecl.GetType())
	g.hints[varDecl] = loadInstruction(varType, g.locals)
	local := g.locals
	g.locals += typeSize(varType)
	g.method.Append(g.genExpr(varDecl.Init)...)
	g.method.Append(storeInstruction(varType, local))

}

func (g *genContext) genAssignStatement(a *ast.AssignStatement) {
	load, store := g.genAssignExprLeft(a.L)
	g.method.Append(load...)
	g.method.Append(g.genExpr(a.R)...)
	g.method.Append(store)
}

func (g *genContext) genReturn(e *ast.Return) {
	if e.X == nil {
		g.method.Append(instruction.Return())
	} else {
		instructions := g.genExpr(e.X)
		g.method.Append(append(instructions, returnInstruction(g.genType(e.ReturnTyp)))...)
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
