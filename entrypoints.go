package main

import (
	"go/ast"
	"go/token"
	"log"
	"sort"
	"strings"
)

func list_all_entrypoints(fset *token.FileSet, afps []*ast.File) string {
	entrypoints := make([]string, 0)

	fc, err := new_func_center(fset, afps)

	if err != nil {
		log.Fatal(err)
	}

	for _, value := range fc.func_defs {
		entrypoints = append(entrypoints, get_tree_string(value))
	}

	sort.Strings(entrypoints)

	return strings.Join(entrypoints, "\n")
}
