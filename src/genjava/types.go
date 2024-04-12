package genjava

import (
	"fmt"
	"trivil/ast"
	"trivil/jasmin/core/tps"
)

//func (g *genContext) isPredefinedType(t ast.Type) bool {
//
//}

func (g *genContext) genType(t ast.Type) tps.T {
	switch x := t.(type) {
	case *ast.FuncType:
		return g.genFunctionType(x)
	case *ast.PredefinedType:
		return predefinedTypeName(x.Name)
	case *ast.TypeRef:
		//fmt.Printf("!TR1 %v %T\n", x.TypeName, x.Typ)
		// Пропускаю type ref до последнего
		var tr = ast.DirectTypeRef(x)
		switch y := tr.Typ.(type) {
		//case *ast.MayBeType:
		//	return g.genType(y.Typ)
		case *ast.PredefinedType:
			return predefinedTypeName(y.Name)
			//default:
			//	panic(fmt.Sprintf("typeref unknown type: %+v", x))
			//	//if g.genTypes {
			//	//	return genc.forwardTypeName(tr.TypeDecl)
			//	//}
			//	//return genc.declName(tr.TypeDecl)
		}
		return tps.NewReference(g.getClassName(x))
	//case *ast.MayBeType:
	//	return genc.typeRef(x.Typ)
	default:
		// TODO: To Be Implemented
		panic(fmt.Sprintf("unknown type: %+v", t))
	}
}

func (g *genContext) genFunctionType(t *ast.FuncType) tps.T {
	var (
		parameters tps.Parameters
		returnType tps.T = tps.Void
	)
	for _, p := range t.Params {
		parameters = append(parameters, g.genType(p.GetType()))
	}

	if t.ReturnTyp != nil {
		returnType = g.genType(t.ReturnTyp)
	}
	return tps.NewMethod(returnType, parameters...)
}

func (g *genContext) getClassName(t ast.Type) string {
	switch x := t.(type) {
	case *ast.TypeRef:
		class := g.scope.GetEntity(x.TypeDecl)
		return class.GetFull()
	}
	panic(fmt.Sprintf("get class name: unexpected type: %+v", t))
}

func predefinedTypeName(name string) tps.T {
	switch name {
	//case "Байт":
	//	return jasmin.NewPrimitiveType(jasmin.BytePrimitive)
	case "Цел64":
		return tps.Long
	case "Вещ64":
		return tps.Double
	case "Слово64":
		return tps.Long
	//case "Лог":
	//	return jasmin.NewPrimitiveType(jasmin.BooleanPrimitive)
	case "Символ":
		return tps.Int
	//case "Строка":
	//	return jasmin.StringType
	//case "Строка8":
	//	return jasmin.StringType
	//case "Пусто":
	//	return jasmin.EmptyType
	default:
		panic(fmt.Sprintf("predefinedTypeName: ni %s", name))
	}
}
