package main

import (
	"errors"
	"go/ast"
	"go/token"
	"go/types"
	"slices"
)

type func_center struct {
	fds []ast.FuncDecl
	// the func_uses field contains a types.Info.Uses
	// map from which everything except function info
	// has been filtered out
	func_uses map[*ast.Ident]types.Object
}

func new_func_center(fset *token.FileSet, afps []*ast.File) (func_center, error) {
	var fc func_center
	var err error

	fc.fds = get_fds_from_afps(afps)

	fc.func_uses, err = get_func_uses(fset, afps)

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
	main_inserted := false

	for ukey, uvalue := range fc.func_uses {
		for _, funcdecl := range fc.fds {
			// artificially insert a main call since the Uses map
			// does not contain calls to main
			if (funcdecl.Name.String() == "main") && (!main_inserted) {
				main_elem := fcall{
					call_name:   "main",
					call_lparen: 0,
					is_method:   false,
					uses_key:    nil,
					uses_value:  nil,
					fd:          &funcdecl,
				}
				fcalls = slices.Insert(fcalls, 0, main_elem)
				main_inserted = true
				continue
			}
			if uvalue.Pos() == funcdecl.Name.Pos() {
				fcall_elem := fcall{
					call_name:   get_tree_string(uvalue),
					call_lparen: ukey.End(),
					is_method:   is_method(funcdecl),
					uses_key:    ukey,
					uses_value:  uvalue,
					fd:          &funcdecl,
				}
				fcalls = append(fcalls, fcall_elem)
			}
		}
	}

	return fcalls, err
}
