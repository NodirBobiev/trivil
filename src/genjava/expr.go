package genjava

import (
	"fmt"
	"trivil/ast"
	"trivil/jasmin"
	"trivil/jasmin/core/instruction"
	"trivil/jasmin/core/tps"
	"trivil/lexer"
)

func (g *genContext) genExpr(expr ast.Expr) instruction.S {
	switch x := expr.(type) {
	case *ast.IdentExpr:
		return jasmin.NewSequence(g.genIdent(x))
	case *ast.LiteralExpr:
		return jasmin.NewSequence(g.genLiteral(x))
	case *ast.SelectorExpr:
		return g.genSelector(x)
	case *ast.BinaryExpr:
		return g.genBinaryExpr(x)
	case *ast.ClassCompositeExpr:
		return g.genClassCompositeExpr(x)
	case *ast.CallExpr:
		return g.genCallExpr(x)
	case *ast.UnaryExpr:
		return g.genUnaryExpr(x)

	default:
		panic(fmt.Sprintf("unknown ast expr: %+v", expr))
	}
}

func (g *genContext) genLiteral(li *ast.LiteralExpr) instruction.I {
	switch li.Kind {
	case ast.Lit_Int:
		g.exprType = tps.Long
		return instruction.Ldc2_w(li.IntVal)
		//return fmt.Sprintf("%d", li.IntVal)
	case ast.Lit_Word:
		g.exprType = tps.Long
		return instruction.Ldc2_w(li.WordVal)
		//return fmt.Sprintf("0x%X", li.WordVal)
	case ast.Lit_Float:
		g.exprType = tps.Double
		return instruction.Ldc2_w(li.FloatStr)
	case ast.Lit_Symbol:
		g.exprType = tps.Int
		return instruction.Iconst(li.IntVal)
		//return fmt.Sprintf("0x%X", li.WordVal)
	//case ast.Lit_String:
	//	//return genc.genStringLiteral(li)
	default:
		panic("ni")
	}
}
func (g *genContext) genIdent(id *ast.IdentExpr) instruction.I {
	varDecl := id.Obj.(*ast.VarDecl)
	return g.hints[varDecl]
}

func (g *genContext) genSelector(s *ast.SelectorExpr) instruction.S {
	field := s.Obj.(*ast.Field)
	return append(g.genExpr(s.X), g.hints[field])
}

func (g *genContext) genBinaryExpr(e *ast.BinaryExpr) instruction.S {
	x := append(g.genExpr(e.X), g.genExpr(e.Y)...)
	exprType := g.genType(e.X.GetType())
	switch e.Op {
	case lexer.ADD:
		return append(x, addInstruction(exprType))
	case lexer.SUB:
		return append(x, subInstruction(exprType))
	case lexer.MUL:
		return append(x, mulInstruction(exprType))
	default:
		panic(fmt.Sprintf("unexpected binary expr op: %+v", e.Op))

	}
}

func (g *genContext) genUnaryExpr(e *ast.UnaryExpr) instruction.S {
	x := g.genExpr(e.X)
	exprType := g.genType(e.GetType())
	switch e.Op {
	case lexer.SUB:
		return append(x, negInstruction(exprType))
	default:
		panic(fmt.Sprintf("unexpected unary expr op: %+v", e.Op))
	}
}

func (g *genContext) genClassCompositeExpr(e *ast.ClassCompositeExpr) instruction.S {
	class := g.classes[e.X.(*ast.IdentExpr).Obj.(*ast.TypeRef).TypeDecl]
	t := jasmin.NewReferenceType(class.Name)
	x := []instruction.I{
		instruction.New(class.Name),
		instruction.Dup(),
		instruction.Invokespecial(class.Name + "/<init>()V"),
	}
	for _, vp := range e.Values {
		f := g.scope.GetEntity(vp.Field)
		x = append(x, jasmin.Dup(t))
		x = append(x, g.genExpr(vp.Value)...)
		x = append(x, jasmin.PutField(f.GetFull(), f.GetType()))
	}

	return x
}

func (g *genContext) genCallExpr(e *ast.CallExpr) jasmin.Sequence {
	load, call := g.genCallExprInvoke(e.X)
	for _, arg := range e.Args {
		load = append(load, g.genExpr(arg)...)
	}
	return append(load, call)
}

func (g *genContext) genCallExprInvoke(e ast.Expr) (jasmin.Sequence, jasmin.Instruction) {
	switch x := e.(type) {
	case *ast.IdentExpr:
		method := g.scope.GetEntityByName(x.Name)
		if method.IsStatic() {
			return jasmin.NewSequence(), jasmin.InvokeStatic(method.GetFull(), method.GetType())
		} else {
			return jasmin.NewSequence(), jasmin.InvokeVirtual(method.GetFull(), method.GetType())
		}
	case *ast.SelectorExpr:
		m := g.scope.GetEntity(x.Obj.(*ast.Function)).(*jasmin.Method)
		if x.X != nil {
			return g.genExpr(x.X), jasmin.InvokeVirtual(m.GetFull(), m.GetType())
		} else {
			return jasmin.NewSequence(), jasmin.InvokeStatic(m.GetFull(), m.GetType())
		}
	}
	panic(fmt.Sprintf("genCallExprInvoke: unexpected expr: %+v", e))
}

func (g *genContext) genAssignExprLeft(e ast.Expr) (jasmin.Sequence, jasmin.Instruction) {
	switch x := e.(type) {
	case *ast.IdentExpr:
		switch d := x.Obj.(type) {
		case *ast.VarDecl:
			e := g.scope.GetEntity(d)
			if v, isVar := e.(*jasmin.Variable); isVar {
				return jasmin.NewSequence(), jasmin.Store(v.Number, v.Type)
			} else {
				return jasmin.NewSequence(), jasmin.PutStatic(e.GetFull(), e.GetType())
			}
		}
	case *ast.SelectorExpr:
		f := g.scope.GetEntity(x.Obj.(*ast.Field))
		return g.genExpr(x.X), jasmin.PutField(f.GetFull(), f.GetType())
	}
	panic(fmt.Sprintf("genAssignExprLeft: unexpeced expr: %+v", e))
}
