package main

import (
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"regexp"
	"strings"
)

func get_all_fds(afs []ast.File) []ast.FuncDecl {
	fds := make([]ast.FuncDecl, 0)

	for _, af := range afs {
		for _, decl := range af.Decls {
			fd, ok := decl.(*ast.FuncDecl)

			if ok {
				fds = append(fds, *fd)
			}
		}
	}

	return fds
}

func get_fds_from_afps(afs []*ast.File) []ast.FuncDecl {
	fds := make([]ast.FuncDecl, 0)

	for _, af := range afs {
		for _, decl := range af.Decls {
			fd, ok := decl.(*ast.FuncDecl)

			if ok {
				fds = append(fds, *fd)
			}
		}
	}

	return fds
}

func node_to_string(node ast.Node, fset *token.FileSet) string {
	var sbuilder strings.Builder

	printer.Fprint(&sbuilder, fset, node)

	return sbuilder.String()
}

func is_selectorstr(call string) bool {
	selectorexpr := regexp.MustCompile(`&{.* .*}`)

	return selectorexpr.MatchString(call)
}

func split_selectorstr(call string) []string {
	call = strings.TrimLeft(call, "&{")
	call = strings.TrimRight(call, "}")

	return strings.Split(call, " ")
}

func get_receiver(recv_types string) string {
	if !is_selectorstr(recv_types) {
		return recv_types
	}

	recv_types = strings.TrimLeft(recv_types, "&{")
	recv_types = strings.TrimRight(recv_types, "}")

	type_slice := strings.Split(recv_types, " ")

	if len(type_slice) == 1 {
		log.Fatal("type_slice should contain two elements.")
	}

	return type_slice[1]
}

func has_nonstd_import(afps []*ast.File) bool {
	for _, af := range afps {
		for _, importspec := range af.Imports {
			// Import paths for std pkgs do not seem to contain
			// a dot. Imports from URLs need a dot.
			if strings.Contains(importspec.Path.Value, ".") {
				return true
			}
		}
	}

	return false
}
