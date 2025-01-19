package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
)

func string_to_ast(src string) (*ast.File, *token.FileSet) {
	fset := token.NewFileSet()
	ast_file, err := parser.ParseFile(fset, "", src, parser.SkipObjectResolution)
	if err != nil {
		log.Fatal(err)
	}

	return ast_file, fset
}

func strings_to_ast(srcs []string) ([]ast.File, *token.FileSet) {
	ast_files := make([]ast.File, 0)

	fset := token.NewFileSet()

	for _, src_file := range srcs {
		af, err := parser.ParseFile(fset, "", src_file, parser.SkipObjectResolution)
		if err != nil {
			log.Fatalf("parser.ParseFile failed.\n%s", err)
		}
		ast_files = append(ast_files, *af)
	}

	return ast_files, fset
}

// Returns slice of pointer to ast.File instead of raw ast.Files.
// Also returns an error instead of just doing a log.Fatal
// which makes the function more testable.
// The downside of using a map here is that the order of the map
// is random.
// So you cannot write unit tests on token.Pos values in conjunction
// with this function.
func strmap_to_ast(srcs map[string]string) ([]*ast.File, *token.FileSet, error) {
	ast_files := make([]*ast.File, 0)

	fset := token.NewFileSet()

	for file_name, file_content := range srcs {
		af, err := parser.ParseFile(fset, file_name, file_content, parser.SkipObjectResolution)
		if err != nil {
			return nil, nil, err
		}
		ast_files = append(ast_files, af)
	}

	return ast_files, fset, nil
}

// same as above but return []ast.File instead of []*ast.File to be
// compatible with parse_dir()
func strmap_to_afs(srcs map[string]string) ([]ast.File, *token.FileSet, error) {
	ast_files := make([]ast.File, 0)

	fset := token.NewFileSet()

	for file_name, file_content := range srcs {
		af, err := parser.ParseFile(fset, file_name, file_content, parser.SkipObjectResolution)
		if err != nil {
			return nil, nil, err
		}
		ast_files = append(ast_files, *af)
	}

	return ast_files, fset, nil
}

// this function might improve performance later
// get all funcdecls of all ast.Files first instead
// of looping through all nodes every time
// in get_funcdecl_from_fname_multifile
func get_funcdecls_from_afs(afs []ast.File) []*ast.FuncDecl {
	fds := make([]*ast.FuncDecl, 0)

	for _, af := range afs {
		for _, decl := range af.Decls {
			fd, ok := decl.(*ast.FuncDecl)

			if !ok {
				continue
			}

			fds = append(fds, fd)
		}
	}

	return fds
}

func unmask_var(i interface{}) {
	fmt.Printf("%v\n", reflect.TypeOf(i))
	fmt.Printf("%v\n", reflect.ValueOf(i))
}

func get_core(i interface{}) string {
	return reflect.TypeOf(i).String() + " - " + reflect.ValueOf(i).String()
}

func print_struct(any interface{}) string {
	return fmt.Sprintf("%+v", any)
}
