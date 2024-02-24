package jasmin

import "fmt"

// Entity is basically Package, Class, Method, Field and Variable
type Entity interface {
	GetName() string
	GetFull() string
	GetParent() Entity
	GetType() Type
	GetAccessFlag() AccessFlag
	IsStatic() bool
	SetName(string)
	SetParent(Entity)
	SetType(Type)
	SetAccessFlag(AccessFlag)
	SetStatic(bool)
}
type EntityBase struct {
	Name       string
	Parent     Entity
	Type       Type
	AccessFlag AccessFlag
	Static     bool
}

func (e *EntityBase) GetName() string {
	return e.Name
}
func (e *EntityBase) GetParent() Entity {
	return e.Parent
}
func (e *EntityBase) GetFull() string {
	if e.Parent != nil {
		return fmt.Sprintf("%s/%s", e.Parent.GetFull(), e.Name)
	}
	return e.Name
}
func (e *EntityBase) GetType() Type {
	return e.Type
}
func (e *EntityBase) GetAccessFlag() AccessFlag {
	return e.AccessFlag
}
func (e *EntityBase) IsStatic() bool {
	return e.Static
}
func (e *EntityBase) SetName(name string) {
	e.Name = name
}
func (e *EntityBase) SetParent(parent Entity) {
	e.Parent = parent
}
func (e *EntityBase) SetType(t Type) {
	e.Type = t
}
func (e *EntityBase) SetAccessFlag(flag AccessFlag) {
	e.AccessFlag = flag
}
func (e *EntityBase) SetStatic(value bool) {
	e.Static = value
}

// EntityStorage keeps track of Package, Class, Method and Field by their full name
type EntityStorage struct {
	Entities map[string]Entity
}

func NewEntityStorage() *EntityStorage {
	return &EntityStorage{Entities: map[string]Entity{}}
}
func (s *EntityStorage) Set(e Entity) {
	s.Entities[e.GetFull()] = e
}
func (s *EntityStorage) Get(fullName string) Entity {
	e, ok := s.Entities[fullName]
	if !ok {
		panic(fmt.Sprintf("entity storage get: %q doesn't exist", fullName))
	}
	return e
}
func (s *EntityStorage) GetPackage(fullName string) *Package {
	e := s.Get(fullName)
	c, ok := e.(*Package)
	if !ok {
		panic(fmt.Sprintf("entity storage get package: %q is not package but %+v", fullName, e))
	}
	return c
}
func (s *EntityStorage) GetClass(fullName string) *Class {
	e := s.Get(fullName)
	c, ok := e.(*Class)
	if !ok {
		panic(fmt.Sprintf("entity storage get class: %q is not class but %+v", fullName, e))
	}
	return c
}
func (s *EntityStorage) GetMethod(fullName string) *Method {
	e := s.Get(fullName)
	c, ok := e.(*Method)
	if !ok {
		panic(fmt.Sprintf("entity storage get method: %q is not method but %+v", fullName, e))
	}
	return c
}
func (s *EntityStorage) GetField(fullName string) *Field {
	e := s.Get(fullName)
	c, ok := e.(*Field)
	if !ok {
		panic(fmt.Sprintf("entity storage get field: %q is not field but %+v", fullName, e))
	}
	return c
}
