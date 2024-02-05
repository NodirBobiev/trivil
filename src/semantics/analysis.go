package semantics

import (
	"fmt"

	"trivil/ast"
	"trivil/env"
	"trivil/semantics/check"
	"trivil/semantics/lookup"
)

var _ = fmt.Printf

func Analyse(m *ast.Module) {
	lookup.Process(m)

	if env.ErrorCount() > 0 {
		return
	}

	check.Process(m)

}
