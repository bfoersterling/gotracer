package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"slices"
)

// Problem: main is not in the uses map
//=> write constructor to artificially construct main fcall
//and add it as the first element of the fcall_slice
//but only if the ast.FuncDecl for main is found
//(needs to be tested for libraries without a main function)

// Maybe have an easy handle like a token.Pos value
// that helps when searching for a specific object.
// Or a string representation of the types.Object.
//
// including a CallExpr here might be too expensive
// but you could use a slice of structs as a lookup table
type fcall struct {
	// call_name: string to be used in the tree
	// for example (*cli_args)my_method
	// can be used as entrypoint if its not a method
	// it should be equal to FuncDecl.Name.String() then
	call_name   string
	call_lparen token.Pos
	is_method   bool
	uses_key    *ast.Ident
	uses_value  types.Object
	fd          *ast.FuncDecl
}

func get_fcall_from_slice(fcalls []fcall, name string) (fcall, error) {
	var err error
	var fcall_result fcall

	for _, v := range fcalls {
		if v.call_name == name {
			return v, err
		}
	}

	err = fmt.Errorf("No fcall with call_name %s was found.\n", name)

	return fcall_result, err
}

func (func_call fcall) get_calls() []*ast.CallExpr {
	calls := make([]*ast.CallExpr, 0)

	ast.Inspect(func_call.fd, func(n ast.Node) bool {
		call_expr, ok := n.(*ast.CallExpr)
		if ok {
			calls = append(calls, call_expr)
		}
		return true
	})

	return calls
}

func (func_call fcall) get_children(all_fcalls []fcall) []fcall {
	callexprs := make([]*ast.CallExpr, 0)
	fcall_children := make([]fcall, 0)
	builtin_funcs := get_builtin_funcs()

	ast.Inspect(func_call.fd, func(n ast.Node) bool {
		callexpr, ok := n.(*ast.CallExpr)

		if !ok {
			return true
		}

		contains_builtin := slices.Contains(builtin_funcs, fmt.Sprintf("%v", callexpr.Fun))

		if !contains_builtin {
			callexprs = append(callexprs, callexpr)
		}
		return true
	})

	// search for callexprs in all fcalls

	for _, callexpr := range callexprs {
		for _, v := range all_fcalls {
			if v.call_lparen == callexpr.Lparen {
				fcall_children = append(fcall_children, v)
			}
		}
	}

	return fcall_children
}
