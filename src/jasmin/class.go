package jasmin

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Class represents class in Jasmin/JVM
type Class struct {
	EntityBase
	EntityStorage
	Super       *Class
	Constructor *Method
}

func NewClass(name string, parent Entity) *Class {
	c := &Class{
		EntityBase: EntityBase{
			Name:       name,
			Parent:     parent,
			AccessFlag: Public,
		},
		Super:         JavaLangObjectClass(),
		Constructor:   nil,
		EntityStorage: *NewEntityStorage(),
	}
	c.Constructor = DefaultConstructor(c, JavaLangObjectClass())
	return c
}

func (c *Class) CreateMethod(name string) *Method {
	m := NewMethod(name, c)
	c.Set(m)
	return m
}
func (c *Class) CreateField(name string, typ Type) *Field {
	f := NewField(name, c, typ)
	c.Set(f)
	return f
}
func (c *Class) String() string {
	result := strings.Builder{}
	result.WriteString(fmt.Sprintf(".class %s %s\n", c.AccessFlag, c.GetFull()))
	result.WriteString(fmt.Sprintf(".super %s\n", c.Super.GetFull()))
	for _, x := range c.Entities {
		if f, isField := x.(*Field); isField {
			result.WriteString(fmt.Sprintf("%s\n", f))
		}
	}
	result.WriteString(c.Constructor.String())
	for _, x := range c.Entities {
		if m, isMethod := x.(*Method); isMethod {
			result.WriteString(m.String())
		}
	}

	return result.String()
}
func (c *Class) Save(dir string) string {
	dest := filepath.Join(dir, strings.ReplaceAll(c.GetFull(), "/", "_")+".j")
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

func MainClass(methodToCall *Method) *Class {
	class := NewClass("Main", nil)
	method := MainMethod(class)
	class.Set(method)
	method.Append(InvokeStatic("builtins/Scan/initScanner()V", VoidMethodType()))
	method.Append(Load(0, NewStringType()))
	method.Append(InvokeStatic(methodToCall.GetFull(), methodToCall.GetType()))
	return class
}

var (
	javaLangObjectClass *Class
)

func JavaLangObjectClass() *Class {
	if javaLangObjectClass == nil {
		javaLangObjectClass = &Class{
			EntityBase: EntityBase{Name: "java/lang/Object"},
		}
		javaLangObjectClass.Constructor = &Method{
			EntityBase: EntityBase{
				Name:       "<init>",
				Type:       VoidMethodType(),
				AccessFlag: Public,
				Parent:     javaLangObjectClass,
			},
		}
	}
	return javaLangObjectClass
}
