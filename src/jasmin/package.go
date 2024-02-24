package jasmin

import "fmt"

// Package represents package in Jasmin/JVM
type Package struct {
	EntityBase
	EntityStorage
}

func NewPackage(name string, parent Entity) *Package {
	return &Package{
		EntityBase: EntityBase{
			Name:   name,
			Parent: parent},
		EntityStorage: *NewEntityStorage(),
	}
}

func (p *Package) CreateClass(name string, super *Class) *Class {
	c := NewClass(name, p)
	if super != nil {
		c.Super = super
		c.Constructor = DefaultConstructor(c, super)
	}
	p.Set(c)
	return c
}

func (p *Package) Show() {
	fmt.Printf("\n===== Package: %s =====\n", p.Name)
	for _, x := range p.Entities {
		if c, isClass := x.(*Class); isClass {
			fmt.Printf("--- Class: %s ---\n", c.GetName())
			fmt.Println(c.String())
		}
	}
}
