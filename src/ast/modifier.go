package ast

import (
	"fmt"
	//	"trivil/env"
)

var _ = fmt.Printf

//====

type Modifier struct {
	Name  string
	Attrs map[string]string
}
