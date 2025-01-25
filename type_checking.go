package main

import (
	"go/ast"
	"go/importer"
	"go/token"
	"go/types"
	"log"
	"strings"
)

// Uses the package name of each ast.File as first argument of the Check
// function. ("package path")
// So this function should be used for local packages.
func get_type_info(fset *token.FileSet, afs []*ast.File) (*types.Info, error) {
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
	}

	conf := types.Config{Importer: importer.Default()}

	// since we already check when parsing the files
	// that they are of the same package, we do not need to do it twice
	// and get the package name from the ast.File
	pkg_name := afs[0].Name.String()

	_, err := conf.Check(pkg_name, fset, afs, info)

	if err != nil {
		return nil, err
	}

	return info, nil
}

// get types info, but only take the Uses field
// and only take the func values
func get_func_info(fset *token.FileSet, afps []*ast.File) (map[*ast.Ident]types.Object, map[*ast.Ident]types.Object, error) {
	filtered_uses := make(map[*ast.Ident]types.Object)
	filtered_defs := make(map[*ast.Ident]types.Object)

	info, err := get_type_info(fset, afps)

	if err != nil {
		return nil, nil, err
	}

	for key, value := range info.Uses {
		tf, ok := value.(*types.Func)

		if !ok {
			continue
		}

		filtered_uses[key] = tf
	}

	for key, value := range info.Defs {
		tf, ok := value.(*types.Func)

		if !ok {
			continue
		}

		filtered_defs[key] = tf
	}

	return filtered_defs, filtered_uses, err
}

func get_tree_string(obj types.Object) string {
	tf, ok := obj.(*types.Func)

	if !ok {
		log.Fatal("obj is not of type *types.Func.")
	}

	tree_string := tf.FullName()

	// remove package name
	tree_string = strings.ReplaceAll(tree_string, obj.Pkg().Name()+".", "")

	return tree_string
}
