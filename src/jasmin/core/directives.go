package core

import (
	"trivil/jasmin/codes"
	"trivil/jasmin/types"
)

func Method(access string, isStatic bool, name string, methodType types.Type) Directive {
	return NewDirective(codes.Method, ".method", access, static(isStatic), name, methodType)
}

func EndMethod() Directive {
	return NewDirective(codes.EndMethod, ".end method")
}

func LimitStack(stack int) Directive {
	return NewDirective(codes.LimitStack, ".limit stack", stack)
}

func LimitLocals(locals int) Directive {
	return NewDirective(codes.LimitLocals, ".limit locals", locals)
}

func Class(access string, class string) Directive {
	return NewDirective(codes.Class, ".class", access, class)
}

func Super(class string) Directive {
	return NewDirective(codes.Super, ".super", class)
}

func Field(access string, isStatic bool, name string, fieldType types.Type) Directive {
	return NewDirective(codes.Field, ".field", access, static(isStatic), name, fieldType)
}
