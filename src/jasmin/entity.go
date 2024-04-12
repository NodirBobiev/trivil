package jasmin

import (
	"fmt"
	"trivil/jasmin/core/tps"
)

type Entity interface {
	GetName() string
	GetFull() string
	GetParent() Entity
	GetType() tps.T
	GetAccessFlag() AccessFlag
	IsStatic() bool
	SetName(string)
	SetParent(Entity)
	SetType(tps.T)
	SetAccessFlag(AccessFlag)
	SetStatic(bool)
}
type EntityBase struct {
	Name       string
	Parent     Entity
	Type       tps.T
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
func (e *EntityBase) GetType() tps.T {
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
func (e *EntityBase) SetType(t tps.T) {
	e.Type = t
}
func (e *EntityBase) SetAccessFlag(flag AccessFlag) {
	e.AccessFlag = flag
}
func (e *EntityBase) SetStatic(value bool) {
	e.Static = value
}
