package stack

import "trivil/jasmin/core/code"

var (
	opCodes = map[code.Code]slots{
		code.Iconst: {
			pop:  0,
			push: 1,
		},
		code.Ldc2_w: {
			pop:  0,
			push: 2,
		},
	}
)
