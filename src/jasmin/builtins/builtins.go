package builtins

import "trivil/jasmin"

var pack *jasmin.Package

func Package() *jasmin.Package {
	if pack == nil {
		pack = jasmin.NewPackage("builtins", nil)
	}
	return pack
}
