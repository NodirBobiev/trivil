package jasmin

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"trivil/jasmin/core/instruction"
	"trivil/jasmin/core/tps"
)

// Class represents class in Jasmin/JVM
type Class struct {
	Name       string
	Directives instruction.S
	Fields     instruction.S
	Methods    []instruction.S
}

func (c *Class) String() string {

	result := strings.Builder{}
	for _, i := range c.Directives {
		result.WriteString(fmt.Sprintf("%s\n", i))
	}
	for _, i := range c.Fields {
		result.WriteString(fmt.Sprintf("%s\n", i))
	}
	for _, seq := range c.Methods {
		for _, i := range seq {
			result.WriteString(fmt.Sprintf("%s\n", i))
		}
	}
	return result.String()
}
func (c *Class) Save(dir string) string {
	dest := filepath.Join(dir, strings.ReplaceAll(c.Name, "/", "_")+".j")
	file, err := os.Create(dest)
	if err != nil {
		panic(fmt.Sprintf("save to %q: create: %s", dest, err))
	}
	defer file.Close()
	_, err = file.WriteString(c.String())
	if err != nil {
		panic(fmt.Sprintf("save to %q: write: %s", dest, err))
	}
	return dest
}

func MainClass(path ...string) *Class {
	return &Class{
		Name: strings.Join(append(path, "Main"), "/"),
		Directives: []instruction.I{
			instruction.Class("public", "Main"),
			instruction.Super("java/lang/Object"),
		},
		Fields:  make(instruction.S, 0),
		Methods: make([]instruction.S, 0),
	}
}

func MainMethod(fnInvoke string) instruction.S {
	return []instruction.I{
		instruction.Method("public", true, "main"+tps.MainMethod.String()),
		instruction.LimitStack(1),
		instruction.LimitLocals(1),
		instruction.Aload(0),
		instruction.Invokestatic(fnInvoke),
		instruction.EndMethod(),
	}
}
