package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"slices"
)

func get_builtin_funcs() []string {
	return []string{
		"append",
		"cap",
		"clear",
		"close",
		"complex",
		"copy",
		"delete",
		"imag",
		"len",
		"make",
		"max",
		"min",
		"new",
		"panic",
		"print",
		"println",
		"real",
		"recover",
	}
}

func get_fd_from_pos(fds []ast.FuncDecl, namepos token.Pos) (*ast.FuncDecl, error) {
	var err error

	for _, fd := range fds {
		if fd.Name.Pos() == namepos {
			return &fd, nil
		}
	}

	err = errors.New("No matching FuncDecl found.")

	return nil, err
}

// WARNING: untested function
// use slice of ast.FuncDecls instead of searching through
// all ast.Files every time
func get_fd_from_fname(fds []ast.FuncDecl, fname string) *ast.FuncDecl {
	for _, fd := range fds {
		if fd.Name.String() == fname {
			return &fd
		}
	}

	return nil
}

// search for a function name in one source file
// get the function definition, not the function calls
func get_funcdecl_from_fname(src ast.Node, fname string) *ast.FuncDecl {
	ast_file := src.(*ast.File)

	for _, v := range ast_file.Decls {
		fn := v.(*ast.FuncDecl)

		if fn.Name.Name == fname {
			return fn
		}
	}

	return nil
}

// search for a function name in multiple source files
// get the function definition, not the function calls
func get_funcdecl_from_fname_multifile(ast_files []ast.File, fname string) *ast.FuncDecl {
	selector := make([]string, 0)
	if is_selectorstr(fname) {
		// split method name and selector
		selector = split_selectorstr(fname)
	}
	for _, ast_file := range ast_files {
		for _, v := range ast_file.Decls {
			fn, ok := v.(*ast.FuncDecl)

			if !ok {
				continue
			}

			if fn.Recv != nil {
				// is method

				if len(selector) == 0 {
					continue
				}
				// compare selector (fmt) and ident (Println)
				// .Recv.List[0].Type may contain a number (maybe a pos?) in front of the type
				// maybe for pointer receivers?
				first_recv_type := fmt.Sprintf("%v", fn.Recv.List[0].Type)
				first_recv_type = get_receiver(first_recv_type)

				if selector[0] != first_recv_type {
					fmt.Printf("DEBUG: %v and %v mismatch!\n", selector[0], first_recv_type)
					continue
				}
				if selector[1] == fn.Name.Name {
					return fn
				}
			}

			if fn.Name.Name == fname {
				return fn
			}
		}
	}

	return nil
}

func is_funcdecl(ast_files []ast.File, fname string) bool {
	for _, ast_file := range ast_files {
		for _, v := range ast_file.Decls {
			fn, ok := v.(*ast.FuncDecl)

			if !ok {
				continue
			}

			if fn.Name.Name == fname {
				return true
			}
		}
	}
	return false
}

func get_string_from_fake_string(any interface{}) string {
	return fmt.Sprintf("%s", any)
}

// return all function calls inside a function
func get_calls(fn *ast.FuncDecl) []*ast.CallExpr {
	calls := make([]*ast.CallExpr, 0)

	ast.Inspect(fn, func(n ast.Node) bool {
		// find function call statements
		// this is a "type assertion"
		call_expr, ok := n.(*ast.CallExpr)
		if ok {
			calls = append(calls, call_expr)
		}
		return true
	})

	return calls
}

// return all function calls inside a node
// but exclude builtin functions
func get_calls_from_node(node ast.Node) []*ast.CallExpr {
	calls := make([]*ast.CallExpr, 0)
	builtin_funcs := get_builtin_funcs()

	if node == nil {
		log.Fatal("You passed a nil pointer to this func.")
	}

	_, ok := node.(*ast.FuncDecl)
	if !ok {
		log.Fatal("Node should be a FuncDecl...")
	}

	ast.Inspect(node, func(n ast.Node) bool {
		call_expr, ok := n.(*ast.CallExpr)
		if ok {
			callexpr_fun := call_expr.Fun
			_, ok := callexpr_fun.(*ast.Ident)
			if ok && slices.Contains(builtin_funcs, fmt.Sprintf("%v", callexpr_fun)) {
				return true
			}

			calls = append(calls, call_expr)
		}
		return true
	})

	return calls
}

// return all function calls inside a node that are calls to function
// within the same package
func get_local_calls_from_node(node ast.Node) []*ast.CallExpr {
	calls := make([]*ast.CallExpr, 0)
	builtin_funcs := get_builtin_funcs()

	ast.Inspect(node, func(n ast.Node) bool {
		call_expr, ok := n.(*ast.CallExpr)
		if ok {
			callexpr_fun := call_expr.Fun
			_, ok := callexpr_fun.(*ast.Ident)
			if ok {
				if slices.Contains(builtin_funcs, fmt.Sprintf("%v", callexpr_fun)) {
					return true
				}
				calls = append(calls, call_expr)
			}
		}
		return true
	})

	return calls
}

func callexprs_to_strings(calls []*ast.CallExpr) []string {
	func_names := make([]string, 0)

	for _, call := range calls {
		fname := fmt.Sprintf("%v", call.Fun)
		func_names = append(func_names, fname)
	}

	return func_names
}

func get_fname_from_call(node ast.Node) string {
	fd, ok := node.(*ast.CallExpr)
	if !ok {
		return ""
	}

	return fmt.Sprintf("%s", fd.Fun)
}

// Goes through CallExprs and
// then uses is_funcdecl to determine if the CallExpr
// is inside a FuncDecl in ast.Files
// this function should be replaced or rewritten
func filter_calls(calls []*ast.CallExpr, afs []ast.File) []*ast.CallExpr {
	filtered_calls := make([]*ast.CallExpr, 0)

	for _, call := range calls {
		call_fname := get_fname_from_call(call)
		if is_funcdecl(afs, call_fname) {
			filtered_calls = append(filtered_calls, call)
		}
	}

	return filtered_calls
}
