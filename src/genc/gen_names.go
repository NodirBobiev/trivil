package genc

import (
	"fmt"
	"trivil/ast"
	"trivil/env"
)

var _ = fmt.Printf

const typeNamePrefix = "T"

// класс струкура и мета информация
const (
	nm_class_fields        = "f"
	nm_class_fields_suffix = "_F"

	nm_base_fields = "_BASE"
	nm_VT_field    = "vtable"

	nm_VT_suffix             = "_VT"
	nm_meta_suffix           = "_Meta"
	nm_meta_field            = "_meta_"
	nm_class_info_suffix     = "_class_info"
	nm_class_info_ptr_suffix = "_class_info_ptr"
	nm_object_init_suffux    = "__init__"

	nm_base_class_info_struct = "_BaseClassInfo"

	nm_variadic_len_suffic = "_len"
)

// prefixes for generated names
const (
	nm_stringLiteral = "strlit"
)

// run-time API
const (
	rt_prefix = "tri_"

	rt_cast_union = "TUnion64"

	rt_init = rt_prefix + "init"

	rt_newLiteralString = rt_prefix + "newLiteralString" // сохраняет в переменную
	rt_newString        = rt_prefix + "newString"
	rt_lenString        = rt_prefix + "lenString"
	rt_emptyString      = rt_prefix + "emptyString"
	rt_equalStrings     = rt_prefix + "equalStrings"

	rt_newVector     = rt_prefix + "newVector"
	rt_newVectorFill = rt_prefix + "newVectorFill"
	//rt_lenVector     = rt_prefix + "lenVector"

	rt_indexcheck = rt_prefix + "indexcheck"
	rt_nilcheck   = rt_prefix + "nilcheck"

	rt_newObject      = rt_prefix + "newObject"
	rt_checkClassType = rt_prefix + "checkClassType"
	rt_isClassType    = rt_prefix + "isClassType"

	rt_convert = rt_prefix

	rt_crash = rt_prefix + "crash"

	rt_tag = rt_prefix + "tag"

	rt_vectorAppend = rt_prefix + "vectorAppend"
)

func (genc *genContext) localName(prefix string) string {
	if prefix == "" {
		prefix = "loc"
	}

	genc.autoNo++
	return fmt.Sprintf("%s%d", prefix, genc.autoNo)
}

//====

func (genc *genContext) declName(d ast.Decl) string {

	out, ok := genc.declNames[d]
	if ok {
		return out
	}

	f, is_fn := d.(*ast.Function)

	if is_fn && f.External {
		name, ok := f.Mod.Attrs["имя"]
		if !ok {
			name = env.OutName(f.Name)
		}
		genc.declNames[d] = name

		return name
	}

	out = ""
	var host = d.GetHost()
	if host != nil {
		out = genc.declName(host) + "__"
	}

	var prefix = ""
	if _, ok := d.(*ast.TypeDecl); ok {
		prefix = typeNamePrefix
	}

	out += prefix + env.OutName(d.GetName())

	genc.declNames[d] = out

	return out
}

func (genc *genContext) outName(name string) string {
	return env.OutName(name)
}

func (genc *genContext) functionName(f *ast.Function) string {

	if f.Recv != nil {
		return genc.typeRef(f.Recv.Typ) + "_" + genc.outName(f.Name)
	}
	return genc.declName(f)
}
