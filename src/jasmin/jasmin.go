package jasmin

import (
	"fmt"
	"os"
)

// ---

type Jasmin struct {
	EntityStorage
	Builtins map[Entity]struct{}
}

func (j *Jasmin) Store(f func() Entity) Entity {
	e := f()
	j.Set(e)
	return e
}

func (j *Jasmin) Save(sourceDir string) []string {
	dir := sourceDir
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		panic(fmt.Sprintf("jasmin save: make dir: %q", dir))
	}
	files := make([]string, 0)
	for _, x := range j.Entities {
		switch p := x.(type) {
		//case *Package:
		//	for _, c := range p.Entities {
		//		cc := c.(*Class)
		//		cc.Save(dir)
		//	}
		case *Class:
			files = append(files, p.Save(dir))
		}
	}
	return files
}

func NewJasmin() *Jasmin {
	return &Jasmin{
		EntityStorage: *NewEntityStorage(),
		Builtins:      make(map[Entity]struct{}),
	}
}
func (j *Jasmin) Show() {
	for _, x := range j.Entities {
		switch p := x.(type) {
		case *Package:
			p.Show()
		case *Class:
			fmt.Printf("--- Class: %s ---\n", p.GetName())
			fmt.Println(p.String())
		}
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
