package jasmin

import (
	"fmt"
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

func NewJasmin() *Jasmin {
	return &Jasmin{
		EntityStorage: *NewEntityStorage(),
		Builtins:      make(map[Entity]struct{}),
	}
}
func (j *Jasmin) Show() {
	for _, x := range j.Entities {
		if p, isPackage := x.(*Package); isPackage {
			p.Show()
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
