package genjava

import (
	"fmt"
	"trivil/ast"
	"trivil/jasmin"
	"trivil/lexer"
)

func (g *genContext) genExpr(expr ast.Expr) jasmin.Sequence {
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
	case *ast.GeneralBracketExpr:
		return g.genGeneralBracketExpr(x)
	default:
		panic(fmt.Sprintf("unknown ast expr: %+v", expr))
	}
}

func (g *genContext) genLiteral(li *ast.LiteralExpr) jasmin.Instruction {
	switch li.Kind {
	case ast.Lit_Int:
		g.exprType = jasmin.NewLongType()
		return jasmin.Const(li.IntVal, jasmin.NewLongType())
		//return fmt.Sprintf("%d", li.IntVal)
	case ast.Lit_Word:
		g.exprType = jasmin.NewLongType()
		return jasmin.Const(li.WordVal, jasmin.NewLongType())
		//return fmt.Sprintf("0x%X", li.WordVal)
	case ast.Lit_Float:
		g.exprType = jasmin.NewDoubleType()
		return jasmin.Const(li.FloatStr, jasmin.NewDoubleType())
		//return li.FloatStr
	case ast.Lit_Symbol:
		g.exprType = jasmin.NewIntType()
		return jasmin.Const(li.WordVal, jasmin.NewIntType())
		//return fmt.Sprintf("0x%X", li.WordVal)
	case ast.Lit_String:
		outs := string(li.StrVal) //genc.EncodeLiteralString(li.StrVal)
		g.exprType = jasmin.NewStringType()
		return jasmin.Const(fmt.Sprintf("%q", outs), g.exprType)
		//return genc.genStringLiteral(li)
	default:
		panic(fmt.Sprintf("unexpected literal:%v", li.Kind))
	}
}
func (g *genContext) genIdent(id *ast.IdentExpr) jasmin.Instruction {
	switch x := id.Obj.(type) {
	case *ast.VarDecl:
		e := g.scope.GetEntity(x)
		g.exprType = e.GetType()
		if v, isVar := e.(*jasmin.Variable); isVar {
			return jasmin.Load(v.Number, v.GetType())
		} else {
			return jasmin.GetStatic(e.GetFull(), e.GetType())
		}
	default:
		panic(fmt.Sprintf("genIdent: unexpected obj: %+v", id.Obj))
	}
}

func (g *genContext) genSelector(s *ast.SelectorExpr) jasmin.Sequence {
	switch x := s.Obj.(type) {
	case *ast.Field:
		field := g.scope.GetEntity(x)
		g.exprType = field.GetType()
		return append(g.genExpr(s.X), jasmin.GetField(field.GetFull(), field.GetType()))
	}
	panic(fmt.Sprintf("gen selector: unexpected expr: %+v", s.Obj))
}

func (g *genContext) genBinaryExpr(e *ast.BinaryExpr) jasmin.Sequence {
	x := append(g.genExpr(e.X), g.genExpr(e.Y)...)
	t := g.genType(e.X.GetType())
	switch e.Op {
	case lexer.ADD:
		g.exprType = t
		return append(x, jasmin.Add(t))
	case lexer.SUB:
		g.exprType = t
		return append(x, jasmin.Sub(t))
	case lexer.MUL:
		g.exprType = t
		return append(x, jasmin.Mul(t))
	default:
		panic(fmt.Sprintf("unexpected binary expr op: %+v", e.Op))

	}
}

func (g *genContext) genUnaryExpr(e *ast.UnaryExpr) jasmin.Sequence {
	x := g.genExpr(e.X)
	switch e.Op {
	case lexer.SUB:
		return append(x, jasmin.Neg(g.exprType))
	default:
		panic(fmt.Sprintf("unexpected unary expr op: %+v", e.Op))
	}
}

func (g *genContext) genClassCompositeExpr(e *ast.ClassCompositeExpr) jasmin.Sequence {
	c := g.scope.GetEntity(e.X.(*ast.IdentExpr).Obj.(*ast.TypeRef).TypeDecl).(*jasmin.Class)
	t := jasmin.NewReferenceType(c.GetFull())
	g.exprType = t
	x := jasmin.NewSequence(
		jasmin.New(t),
		jasmin.Dup(t),
		jasmin.InvokeSpecial(c.Constructor.GetFull(), c.Constructor.GetType()))
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
	case *ast.GeneralBracketExpr:
		arrayType := g.genType(x.X.GetType()).(*jasmin.ArrayType)
		indexType := g.genType(x.Index.GetType())
		return append(append(g.genExpr(x.X), g.genExpr(x.Index)...),
				jasmin.CastPrimitives(indexType, jasmin.NewIntType())),
			jasmin.Astore(arrayType.ElementType)
	}

	panic(fmt.Sprintf("genAssignExprLeft: unexpeced expr: %+v", e))
}

func (g *genContext) genGeneralBracketExpr(e *ast.GeneralBracketExpr) jasmin.Sequence {
	if e.Index != nil {
		arrayType := g.genType(e.X.GetType()).(*jasmin.ArrayType)
		result := g.genExpr(e.X)
		// calculate index expression
		result = append(result, g.genExpr(e.Index)...)
		// cast index value to int
		result = append(result, jasmin.CastPrimitives(g.genType(e.Index.GetType()), jasmin.NewIntType()))
		// load the element at given index onto the stack
		result = append(result, jasmin.Aload(arrayType.ElementType))
		return result
	}
	return g.genArrayCompositeExpr(e.Composite)
}

func (g *genContext) genArrayCompositeExpr(e *ast.ArrayCompositeExpr) jasmin.Sequence {
	arrayType := g.genType(e.GetType()).(*jasmin.ArrayType)

	var (
		lenType = jasmin.NewIntType()
		lenExpr = jasmin.NewSequence(jasmin.Const(0, lenType))
	)
	if e.LenExpr != nil {
		lenExpr = append(
			g.genExpr(e.LenExpr),
			jasmin.CastPrimitives(g.genType(e.LenExpr.GetType()), lenType))
	} else if len(e.Values) > 0 {
		lenExpr = jasmin.NewSequence(jasmin.Const(len(e.Values), lenType))
	}
	result := append(lenExpr, jasmin.NewArray(arrayType.ElementType))

	if e.Default != nil {
		defaultType := g.genType(e.Default.GetType())
		defaultLocalNumber := g.method.GetLocalNumber(defaultType)
		defaultExpr := append(g.genExpr(e.Default), jasmin.Store(defaultLocalNumber, defaultType))
		result = append(result, defaultExpr...)
		result = append(result, g.genArrayFiller(defaultLocalNumber, arrayType)...)
	}
	indexType := jasmin.NewIntType()
	for i := 0; i < len(e.Values); i++ {
		valueExpr := e.Values[i]
		// load array
		result = append(result, jasmin.Dup(arrayType))
		if len(e.Indexes) > 0 {
			indexExpr := e.Indexes[i]
			// calculate index expression
			result = append(result, g.genExpr(indexExpr)...)
			// cast index value to int
			result = append(result, jasmin.CastPrimitives(g.genType(indexExpr.GetType()), indexType))
		} else {
			if i == 0 {
				result = append(result, jasmin.Const(i, indexType))
			} else {
				result = append(result, jasmin.Const(i, indexType))
			}
		}
		// calculate value expression
		result = append(result, g.genExpr(valueExpr)...)
		// store
		result = append(result, jasmin.Astore(arrayType.ElementType))
	}
	return result
}

func (g *genContext) genSimpleLocalVariable(value any, typ jasmin.Type) (jasmin.Sequence, int) {
	localNumber := g.method.GetLocalNumber(typ)
	return jasmin.NewSequence(
			jasmin.Const(value, typ),
			jasmin.Store(localNumber, typ)),
		localNumber
}

func (g *genContext) genArrayFiller(filler int, arrayType *jasmin.ArrayType) jasmin.Sequence {
	var (
		startLabel = g.genLabel("FILLER_START")
		endLabel   = g.genLabel("FILLER_END")
	)
	indexType := jasmin.NewIntType()
	// initialize an indexing variable
	result, index := g.genSimpleLocalVariable(0, indexType)

	// start the loop
	result = append(result, jasmin.NewLabel(startLabel))

	// check if index variable is greater or equal than length and jump to endLabel
	result = append(result, jasmin.Dup(arrayType), jasmin.ArrayLength(), jasmin.Load(index, indexType))
	result = append(result, jasmin.IfIcmp("le", endLabel))

	// load array, load index, load filler and store
	result = append(result,
		jasmin.Dup(arrayType),
		jasmin.Load(index, indexType),
		jasmin.Load(filler, arrayType.ElementType),
		jasmin.Astore(arrayType.ElementType))

	// increment index
	result = append(result, jasmin.Iinc(index, 1))

	// jump to the start of the loop
	result = append(result, jasmin.Goto(startLabel))

	// end the loop
	result = append(result, jasmin.NewLabel(endLabel))

	return result
}
