package main

import "go/ast"

func is_method(fd ast.FuncDecl) bool {
	if fd.Recv == nil {
		return false
	}

	return true
}
