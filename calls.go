package main

import (
	"go/ast"
)

// convenience function for tests
func get_calls_from_afps(afps []*ast.File) ([]*ast.CallExpr, error) {
	var err error
	callexprs := make([]*ast.CallExpr, 0)

	for _, afp := range afps {
		ast.Inspect(afp, func(n ast.Node) bool {
			callexpr, ok := n.(*ast.CallExpr)
			if ok {
				callexprs = append(callexprs, callexpr)
			}
			return true
		})
	}

	return callexprs, err
}
