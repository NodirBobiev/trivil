package common

import "errors"

type Test interface {
	GetName() string
	GetGroup() string
	GetID() string
	Run() error
}

type TestBase struct {
	Name  string
	Group string
	ID    string
}

func (t *TestBase) GetName() string {
	return t.Name
}
func (t *TestBase) GetID() string {
	return t.ID
}
func (t *TestBase) GetGroup() string {
	return t.Group
}

func (t *TestBase) Run() error {
	return errors.New("unimplemented")
}
