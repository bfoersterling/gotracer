package main

import (
	"go/ast"
	"go/token"
	"sort"
	"strings"
)

func list_uncalled_funcs(fset *token.FileSet, afps []*ast.File) string {
	uncalled_funcs := make([]string, 0)

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
			uncalled_funcs = append(uncalled_funcs, fcall.call_name)
		}
	}

	sort.Strings(uncalled_funcs)

	return strings.Join(uncalled_funcs, "\n")
}
