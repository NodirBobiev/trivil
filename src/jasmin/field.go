package jasmin

import "fmt"

// Field represents Field in Jasmin/JVM
type Field struct {
	EntityBase
}

func NewField(name string, class *Class, t Type) *Field {
	return &Field{EntityBase{
		Name:       name,
		Parent:     class,
		Type:       t,
		AccessFlag: Protected,
		Static:     false,
	}}
}

func (f *Field) String() string {
	var static string
	if f.IsStatic() {
		static = " static"
	}
	return fmt.Sprintf(".field %s%s %s %s", f.AccessFlag, static, f.Name, f.Type)
}

// Variable represents vars in Jasmin/JVM
type Variable struct {
	EntityBase
	Number int
}

func NewVariable(name string, method *Method, t Type, number int) *Variable {
	return &Variable{
		EntityBase: EntityBase{
			Name:       name,
			Parent:     method,
			Type:       t,
			AccessFlag: 0,
			Static:     false,
		},
		Number: number,
	}
}
