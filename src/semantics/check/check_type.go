package check

import (
	"fmt"

	"trivil/ast"
	"trivil/env"
)

var _ = fmt.Printf

func (cc *checkContext) isAlreadyChecked(v *ast.TypeDecl) bool {

	if v.Host != cc.module {
		return true
	}
	_, ok := cc.checkedTypes[v.Name]
	return ok
}

func (cc *checkContext) typeDecl(td *ast.TypeDecl) {

	switch x := td.Typ.(type) {
	case *ast.InvalidType:
		// nothing
	case *ast.VectorType:
		// есть ли что проверять?
	case *ast.ClassType:
		cc.classType(td, x)
	case *ast.MayBeType:
		/*nothing*/
	case *ast.TypeRef:
		/*nothing*/
	default:
		panic(fmt.Sprintf("check typeDecl: ni %T", td.Typ))
	}
}

func (cc *checkContext) classType(td *ast.TypeDecl, cl *ast.ClassType) {

	if cc.isAlreadyChecked(td) {
		return
	}
	cc.checkedTypes[td.Name] = struct{}{}

	if cl.BaseTyp != nil {
		cc.classBaseType(cl, cl.Members)
	}

	for _, f := range cl.Fields {

		if f.Later {
			if f.Typ == nil {
				env.AddError(f.Pos, "СЕМ-ДЛЯ-ПОЗЖЕ-НУЖЕН-ТИП")
			}
		} else {
			cc.expr(f.Init)

			if f.Typ != nil {
				cc.checkAssignable(f.Typ, f.Init)
			} else {
				f.Typ = f.Init.GetType()
				if f.Typ == nil {
					panic("assert - не задан тип поля")
				}
			}
		}

		prev, ok := cl.Members[f.Name]
		if ok {
			env.AddError(f.Pos, "СЕМ-ДУБЛЬ-В-КЛАССЕ", f.Name, env.PosString(prev.(ast.Node).GetPos()))
		} else {
			cl.Members[f.Name] = f
		}
	}

	for _, m := range cl.Methods {
		prev, ok := cl.Members[m.Name]
		if ok {
			prevM, ok := prev.(*ast.Function)
			if ok && ast.UnderType(prevM.Recv.Typ) != ast.UnderType(m.Recv.Typ) {
				// сигнатуры при переопределении должны совпадать
				var res = cc.compareFuncTypes(m.Typ, prevM.Typ)
				if res != "" {
					env.AddError(m.Pos, "СЕМ-РАЗНЫЕ-ТИПЫ-МЕТОДОВ", m.Name, res)
				} else {
					cl.Members[m.Name] = m
				}
			} else {
				env.AddError(m.Pos, "СЕМ-ДУБЛЬ-В-КЛАССЕ", m.Name, env.PosString(prev.(ast.Node).GetPos()))
			}
		} else {
			cl.Members[m.Name] = m
		}
	}
}

func (cc *checkContext) classBaseType(cl *ast.ClassType, members map[string]ast.Decl) {

	var tr = cl.BaseTyp.(*ast.TypeRef)

	baseClass, ok := tr.Typ.(*ast.ClassType)
	if !ok {
		env.AddError(tr.Pos, "СЕМ-БАЗА-НЕ-КЛАСС")
		return
	}

	if !cc.isAlreadyChecked(tr.TypeDecl) {
		cc.classType(tr.TypeDecl, baseClass)
	}

	if baseClass.BaseTyp != nil {
		cc.classBaseType(baseClass, members)
	}

	for _, f := range baseClass.Fields {
		members[f.Name] = f
	}
	for _, m := range baseClass.Methods {
		members[m.Name] = m
	}
}

// Возвращает "", если равны или причину, если разные
func (cc *checkContext) compareFuncTypes(t1, t2 ast.Type) string {
	ft1, ok1 := t1.(*ast.FuncType)
	ft2, ok2 := t2.(*ast.FuncType)
	if !ok1 || !ok2 {
		return "" // а вдруг где-то Invalid type
	}

	if len(ft1.Params) != len(ft2.Params) {
		return "разное число параметров"
	}

	for i, p := range ft1.Params {
		if p.Out != ft2.Params[i].Out {
			return fmt.Sprintf("не совпадает признак выходного параметра '%s'", p.Name)
		}
		if !equalTypes(p.Typ, ft2.Params[i].Typ) {
			return fmt.Sprintf("не совпадает тип у параметра '%s'", p.Name)
		}
	}

	if !equalTypes(ft1.ReturnTyp, ft2.ReturnTyp) {
		return "разные типы результата"
	}

	return ""
}

// Возвращает true, если можно присвоить
func (cc *checkContext) assignable(lt ast.Type, r ast.Expr) bool {

	cc.errorHint = ""

	if equalTypes(lt, r.GetType()) {
		return true
	}

	var t = ast.UnderType(lt)

	switch t {
	case ast.Byte:
		var li = literal(r)
		if li != nil {
			if li.Kind == ast.Lit_Int && (li.IntVal >= 0 || li.IntVal <= 255) {
				li.WordVal = uint64(li.IntVal)
				li.Typ = ast.Byte
				return true
			} else if li.Kind == ast.Lit_Word && li.WordVal <= 255 {
				li.Typ = ast.Byte
				return true
			}
		}
	case ast.Word64:
		var li = literal(r)
		if li != nil {
			if li.Kind == ast.Lit_Int && li.IntVal >= 0 {
				li.WordVal = uint64(li.IntVal)
				li.Typ = ast.Word64
				return true
			} else if li.Kind == ast.Lit_Word {
				li.Typ = ast.Word64
				return true
			}
		}
	case ast.Int64:
		var li = literal(r)
		if li != nil {
			if li.Kind == ast.Lit_Word && li.WordVal <= 1<<63-1 {
				li.Typ = ast.Int64
				return true
			}
		}
	case ast.TagPairType:
		return ast.HasTag(r.GetType())
	}

	switch xt := t.(type) {
	/*
		case *ast.VectorType:
			rvec, ok := ast.UnderType(r.GetType()).(*ast.VectorType)
			if ok && equalTypes(xt.ElementTyp, rvec.ElementTyp) {
				return true
			}
	*/
	case *ast.ClassType:
		rcl, ok := ast.UnderType(r.GetType()).(*ast.ClassType)
		if ok && isDerivedClass(xt, rcl) {
			return true
		}
	case *ast.MayBeType:
		var rt = ast.UnderType(r.GetType())
		if rt == ast.NullType {
			return true
		} else if cc.assignable(xt.Typ, r) {
			return true
		}

	}

	// TODO: function types, ...
	return false
}

func (cc *checkContext) checkAssignable(lt ast.Type, r ast.Expr) {

	if ast.IsVoidType(r.GetType()) {
		env.AddError(r.GetPos(), "СЕМ-ФН-НЕТ-ЗНАЧЕНИЯ")
		return
	}

	if cc.assignable(lt, r) {
		return
	}
	if ast.IsInvalidType(lt) || ast.IsInvalidType(r.GetType()) {
		return
	}

	env.AddError(r.GetPos(), "СЕМ-НЕСОВМЕСТИМО-ПРИСВ", cc.errorHint,
		ast.TypeName(lt), ast.TypeName(r.GetType()))
}

func equalTypes(t1, t2 ast.Type) bool {
	t1 = ast.UnderType(t1)
	t2 = ast.UnderType(t2)
	if t1 == t2 {
		return true
	}
	switch x1 := t1.(type) {
	case *ast.VectorType:
		x2, ok := t2.(*ast.VectorType)
		return ok && equalTypes(x1.ElementTyp, x2.ElementTyp)
	case *ast.VariadicType:
		x2, ok := t2.(*ast.VariadicType)
		return ok && equalTypes(x1.ElementTyp, x2.ElementTyp)
	case *ast.MayBeType:
		x2, ok := t2.(*ast.MayBeType)
		return ok && equalTypes(x1.Typ, x2.Typ)
	}
	return false
}

//=== ошибки с проверкой

func addErrorForType(t ast.Type, pos int, id string, args ...interface{}) {
	if !ast.IsInvalidType(t) {
		env.AddError(pos, id, args...)
	}
}
