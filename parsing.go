package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
)

func parse_dir(go_dir string) ([]ast.File, *token.FileSet, error) {
	files := get_gofiles(go_dir)

	ast_files := make([]ast.File, 0)

	fset := token.NewFileSet()

	var err error

	var package_name string

	for i, file := range files {
		ast_file, err := parser.ParseFile(fset, file, nil, parser.SkipObjectResolution)

		if err != nil {
			return nil, nil, err
		}

		if i == 0 {
			package_name = ast_file.Name.String()
		}

		if (i > 0) && (ast_file.Name.String() != package_name) {
			err = errors.New("All files need to be of the same package.")
			return nil, nil, err
		}

		ast_files = append(ast_files, *ast_file)
	}

	return ast_files, fset, err
}

func parse_dir_afps(go_dir string) ([]*ast.File, *token.FileSet, error) {
	files := get_gofiles(go_dir)

	ast_files := make([]*ast.File, 0)

	fset := token.NewFileSet()

	var err error

	var package_name string

	for i, file := range files {
		ast_file, err := parser.ParseFile(fset, file, nil, parser.SkipObjectResolution)

		if err != nil {
			return nil, nil, err
		}

		if i == 0 {
			package_name = ast_file.Name.String()
		}

		if (i > 0) && (ast_file.Name.String() != package_name) {
			err = errors.New("All files need to be of the same package.")
			return nil, nil, err
		}

		ast_files = append(ast_files, ast_file)
	}

	return ast_files, fset, err
}
