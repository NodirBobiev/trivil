package genjava

import (
	"fmt"
	"trivil/ast"
	"trivil/jasmin"
)

type Scope struct {
	Decls        map[ast.Decl]jasmin.Entity
	CurrentScope *ast.Scope
}

func NewScope() *Scope {
	return &Scope{
		Decls: map[ast.Decl]jasmin.Entity{},
	}
}
func (s *Scope) SetScope(scope *ast.Scope) {
	s.CurrentScope = scope

	//for scope != nil {
	//	for _, d := range scope.Names {
	//		s.Generate(d)
	//	}
	//	scope = scope.Outer
	//}
}
func (s *Scope) GetDecl(name string) ast.Decl {
	scope := s.CurrentScope
	for scope != nil {
		d, ok := scope.Names[name]
		if ok {
			return d
		}
		scope = scope.Outer
	}
	return nil
}

func (s *Scope) Exists(d ast.Decl) bool {
	_, ok := s.Decls[d]
	return ok
}

func (s *Scope) SetEntity(d ast.Decl, e jasmin.Entity) {
	s.Decls[d] = e
}

func (s *Scope) GetEntity(d ast.Decl) jasmin.Entity {
	return s.Decls[d]
}

func (s *Scope) GetEntityByName(name string) jasmin.Entity {
	return s.GetEntity(s.GetDecl(name))
}

func GetTypeName(t ast.Type) string {
	switch x := t.(type) {
	case *ast.PredefinedType:
		return x.Name
	case *ast.TypeRef:
		//fmt.Printf("!TR1 %v %T\n", x.TypeName, x.Typ)
		// Пропускаю type ref до последнего
		var tr = ast.DirectTypeRef(x)
		switch y := tr.Typ.(type) {
		case *ast.MayBeType:
			return GetTypeName(y.Typ)
		case *ast.PredefinedType:
			return y.Name
		default:
			return tr.TypeName
		}
	case *ast.MayBeType:
		return GetTypeName(x.Typ)
	default:
		panic(fmt.Sprintf("assert: %T", t))
	}
}
