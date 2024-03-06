package builtins

import "trivil/jasmin"

var (
	printClass        *jasmin.Class
	printLongMethod   *jasmin.Method
	printDoubleMethod *jasmin.Method
	printlnMethod     *jasmin.Method
)

func PrintClass() *jasmin.Class {
	if printClass == nil {
		printClass = Package().CreateClass("Print", jasmin.JavaLangObjectClass())
	}
	return printClass
}

func unaryPrintMethod(name string, t jasmin.Type) *jasmin.Method {
	tt := jasmin.UnaryVoidMethodType(t)
	m := jasmin.NewFullMethod(name, true, jasmin.Public, tt, PrintClass())
	m.Locals = t.StackSlot()
	m.Append(
		jasmin.GetStatic("java/lang/System/out", jasmin.NewReferenceType("java/io/PrintStream")),
		jasmin.Load(0, t),
		jasmin.InvokeVirtual("java/io/PrintStream/print"+tt.String(), tt))
	return m
}

func PrintLongMethod() *jasmin.Method {
	if printLongMethod == nil {
		printLongMethod = unaryPrintMethod("print_long", jasmin.NewLongType())
	}
	return printLongMethod
}

func PrintDoubleMethod() *jasmin.Method {
	if printDoubleMethod == nil {
		printDoubleMethod = unaryPrintMethod("print_double", jasmin.NewDoubleType())
	}
	return printDoubleMethod
}

func PrintlnMethod() *jasmin.Method {
	if printlnMethod == nil {
		t := jasmin.VoidMethodType()
		printlnMethod = jasmin.NewFullMethod("println", true, jasmin.Public, t, PrintClass())

		printlnMethod.Append(
			jasmin.GetStatic("java/lang/System/out", jasmin.NewReferenceType("java/io/PrintStream")),
			jasmin.InvokeVirtual("java/io/PrintStream/println"+t.String(), t))
	}
	return printlnMethod
}
