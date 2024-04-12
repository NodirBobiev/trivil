package jasmin

import (
	"fmt"
	"os"
)

// ---

// Package represents package in Jasmin/JVM
type Package struct {
	Name    string
	Classes map[string]*Class
}

func NewPackage(name string) *Package {
	pack := &Package{
		Name:    name,
		Classes: make(map[string]*Class),
	}
	mainClass := MainClass(name)
	pack.Classes[mainClass.Name] = mainClass
	return pack
}

func (p *Package) Show() {
	fmt.Printf("\n===== Package: %s =====\n", p.Name)
	for _, c := range p.Classes {
		fmt.Printf("--- Class: %s ---\n", c.Name)
		fmt.Println(c.String())
	}
}

// ---

type Jasmin struct {
	Packages map[string]*Package
}

func (j *Jasmin) Save(sourceDir string) []string {
	dir := sourceDir
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		panic(fmt.Sprintf("jasmin save: make dir: %q", dir))
	}
	files := make([]string, 0)
	for _, p := range j.Packages {
		for _, c := range p.Classes {
			files = append(files, c.Save(dir))
		}
	}
	return files
}

func (j *Jasmin) Show() {
	for _, p := range j.Packages {
		p.Show()
	}
}

type AccessFlag int

const (
	Public AccessFlag = iota + 1
	Protected
	Private
)

func (a AccessFlag) String() string {
	switch a {
	case Public:
		return "public"
	case Protected:
		return "protected"
	case Private:
		return "private"
	default:
		panic(fmt.Sprintf("unknown access flag: %+v", a))
	}
}
