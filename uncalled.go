package main

import (
	"go/ast"
	"go/token"
)

func list_uncalled_funcs(fset *token.FileSet, afps []*ast.File) string {
	var uncalled_funcs string

	fc, err := new_func_center(fset, afps)

	if err != nil {
		panic(err)
	}

	fcalls, err := fc.get_fcalls()

	if err != nil {
		panic(err)
	}

	for _, fcall := range fcalls {
		if fcall.uses_key == nil {
			uncalled_funcs += fcall.call_name + "\n"
		}
	}

	return uncalled_funcs
}
