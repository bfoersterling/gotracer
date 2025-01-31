package main

import (
	"errors"
	"go/ast"
	"go/token"
	"go/types"
	"log"
)

type func_center struct {
	fds []ast.FuncDecl
	// all types.Info.Defs that contain func info
	func_defs map[*ast.Ident]types.Object
	// the func_uses field contains a types.Info.Uses
	// map from which everything except function info
	// has been filtered out
	func_uses map[*ast.Ident]types.Object
}

func new_func_center(fset *token.FileSet, afps []*ast.File) (func_center, error) {
	var fc func_center
	var err error

	fc.fds = get_fds_from_afps(afps)

	fc.func_defs, fc.func_uses, err = get_func_info(fset, afps)

	return fc, err
}

// get corresponding ast.FuncDecl from an *ast.CallExpr
func (fc func_center) get_funcdecl(call *ast.CallExpr) (*ast.FuncDecl, error) {
	var value_pos token.Pos
	var err error

	// get matching uses position
	for key, value := range fc.func_uses {
		if key.End() == call.Lparen {
			value_pos = value.Pos()
		}
	}

	if !value_pos.IsValid() {
		err = errors.New("CallExpr was not found in types info.")
		return nil, err
	}

	// match with funcdecl
	for _, fd := range fc.fds {
		if fd.Name.Pos() == value_pos {
			return &fd, err
		}
	}

	err = errors.New("No matching FuncDecl was found.")

	return nil, err
}

func (fc func_center) get_fcalls() ([]fcall, error) {
	fcalls := make([]fcall, 0)
	var err error

	for ukey, uvalue := range fc.func_uses {
		for dkey, dvalue := range fc.func_defs {
			if uvalue == dvalue {
				fcall_elem := fcall{
					call_name:   get_tree_string(uvalue),
					call_lparen: ukey.End(),
					defs_key:    dkey,
					defs_value:  dvalue,
					uses_key:    ukey,
					uses_value:  uvalue,
				}

				fcalls = append(fcalls, fcall_elem)
				break
			}
		}
	}

	called := false
	for _, funcdecl := range fc.fds {
		for i, v := range fcalls {
			if funcdecl.Name == v.defs_key {
				// function is actually called
				fcalls[i].fd = &funcdecl
				fcalls[i].is_method = is_method(funcdecl)
				called = true
				// don't break because the function can be called
				// multiple times and every call needs to be appended
				continue
			}

			// last element and still not found a matching fcall
			// append a new fcall element for the uncalled function
			if i == (len(fcalls)-1) && !called {
				fcall_elem := fcall{
					fd:        &funcdecl,
					is_method: is_method(funcdecl),
				}
				fcalls = append(fcalls, fcall_elem)
			}
		}
		called = false
	}

	// uncalled functions still need to get info from Defs
	for key, value := range fc.func_defs {
		for i, fcall := range fcalls {
			if fcall.fd == nil {
				log.Fatalf("DEBUG: empty fd field! (fcall: %v)\n", fcall)
			}
			if fcall.call_name != "" {
				// called functions do not need to be refilled
				continue
			}
			if key == fcall.fd.Name {
				fcalls[i].call_name = get_tree_string(value)
				fcalls[i].defs_key = key
				fcalls[i].defs_value = value
				break
			}
		}
	}

	return fcalls, err
}
