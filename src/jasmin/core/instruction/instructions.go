package instruction

import (
	"trivil/jasmin/core/code"
	"trivil/jasmin/core/tps"
)

func static(isStatic bool) string {
	if isStatic {
		return "static"
	}
	return ""
}

func Label(label string) I {
	return newI(code.Label, label+":")
}

func Method(access string, isStatic bool, nameAndType string) I {
	return newI(code.Method, ".method", access, static(isStatic), nameAndType)
}
func EndMethod() I {
	return newI(code.EndMethod, ".end method")
}

func LimitStack(stack int) I {
	return newI(code.LimitStack, ".limit stack", stack)
}

func LimitLocals(locals int) I {
	return newI(code.LimitLocals, ".limit locals", locals)
}

func Class(access string, class string) I {
	return newI(code.Class, ".class", access, class)
}

func Super(class string) I {
	return newI(code.Super, ".super", class)
}

func Field(access string, isStatic bool, name string, fieldType tps.T) I {
	return newI(code.Field, ".field", access, static(isStatic), name, fieldType)
}

func Iconst(value any) I {
	return newI(code.Iconst, "iconst", value)
}

func Ldc2_w(value any) I {
	return newI(code.Ldc2_w, "ldc2_w", value)
}

func Istore(local int) I {
	return newI(code.Istore, "istore", local)
}

func Lstore(local int) I {
	return newI(code.Lstore, "lstore", local)
}

func Dstore(local int) I {
	return newI(code.Dstore, "dstore", local)
}

func Astore(local int) I {
	return newI(code.Astore, "astore", local)
}

func Iload(local int) I {
	return newI(code.Iload, "iload", local)
}

func Lload(local int) I {
	return newI(code.Lload, "lload", local)
}

func Dload(local int) I {
	return newI(code.Dload, "dload", local)
}

func Aload(local int) I {
	return newI(code.Aload, "aload", local)
}

func Iadd() I {
	return newI(code.Iadd, "iadd")
}

func Ladd() I {
	return newI(code.Ladd, "ladd")
}

func Dadd() I {
	return newI(code.Dadd, "dadd")
}

func Isub() I { return newI(code.Isub, "isub") }

func Lsub() I {
	return newI(code.Lsub, "lsub")
}

func Dsub() I {
	return newI(code.Dsub, "dsub")
}

func Imul() I {
	return newI(code.Imul, "imul")
}

func Lmul() I {
	return newI(code.Lmul, "lmul")
}

func Dmul() I {
	return newI(code.Dmul, "dmul")
}

func Ineg() I {
	return newI(code.Ineg, "ineg")
}

func Lneg() I {
	return newI(code.Lneg, "lneg")
}

func Dneg() I {
	return newI(code.Dneg, "dneg")
}

func Lcmp() I {
	return newI(code.Lcmp, "lcmp")
}

func Dcmpl() I {
	return newI(code.Dcmpl, "dcmpl")
}

func New(class string) I {
	return newI(code.New, "new", class)
}

func Dup() I {
	return newI(code.Dup, "dup")
}

func Getstatic(field, fieldType string) I {
	return newI(code.Getstatic, "getstatic", field, fieldType)
}

func Getfield(field, fieldType string) I {
	return newI(code.Getfield, "getfield", field, fieldType)
}

func Putfield(field, fieldType string) I {
	return newI(code.Putfield, "putfield", field, fieldType)
}

func Putstatic(field, fieldType string) I {
	return newI(code.Putstatic, "putstatic", field, fieldType)
}

func Invokespecial(fn string) I {
	return newI(code.Invokespecial, "invokespecial", fn)
}

func Invokestatic(fn string) I {
	return newI(code.Invokestatic, "invokestatic", fn)
}

func Invokevirtual(fn string) I {
	return newI(code.Invokevirtual, "invokevirtual", fn)
}

func Return() I {
	return newI(code.Return, "return")
}

func Ireturn() I {
	return newI(code.Ireturn, "ireturn")
}

func Lreturn() I {
	return newI(code.Lreturn, "lreturn")
}

func Dreturn() I {
	return newI(code.Dreturn, "dreturn")
}

func Areturn() I {
	return newI(code.Areturn, "areturn")
}

func Goto(label string) I { return newI(code.Goto, "goto", label) }

func Ifne(label string) I { return newI(code.Ifne, "ifne", label) }

func Ifeq(label string) I { return newI(code.Ifeq, "ifeq", label) }

func Ifle(label string) I { return newI(code.Ifle, "ifle", label) }

func Ifgt(label string) I { return newI(code.Ifgt, "ifgt", label) }

func Ifge(label string) I { return newI(code.Ifge, "ifge", label) }

func Iflt(label string) I { return newI(code.Iflt, "iflt", label) }
