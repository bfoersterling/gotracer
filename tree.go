package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"log"
	"os"
	"strings"
)

func filetree(fset *token.FileSet, af ast.File) {
	//	main_decl := get_funcdecl_from_fname(&af, "main")

	// temporary: construct slice consisting of one element to make it still work for one file
	af_slice := []ast.File{af}

	tree(os.Stdout, "main", af_slice, "")
}

func silent_filetree(fset *token.FileSet, af ast.File) string {
	//	main_decl := get_funcdecl_from_fname(&af, "main")

	var tree_buffer bytes.Buffer

	// temporary: construct slice consisting of one element to make it still work for one file
	af_slice := []ast.File{af}

	tree(&tree_buffer, "main", af_slice, "")

	return tree_buffer.String()
}

func verbose_dirtree(fset *token.FileSet, afs []ast.File) {
	fmt.Printf("main\n")

	tree(os.Stdout, "main", afs, "")
}

func verbose_calltree(fset *token.FileSet, afps []*ast.File, entrypoint string) {
	func_center, err := new_func_center(fset, afps)

	if err != nil {
		log.Fatalf("new_func_center() failed with err:\n%v\n", err)
	}

	fcalls, err := func_center.get_fcalls()

	if err != nil {
		log.Fatalf("func_center.get_fcalls() failed with err:\n%v\n", err)
	}

	_, err = get_fcall_from_slice(fcalls, entrypoint)

	if err != nil {
		if strings.Contains(err.Error(), "No fcall with call_name") {
			fmt.Printf("%s: Function %q was not found.\n", os.Args[0], entrypoint)
			os.Exit(10)
		} else {
			panic(err)
		}
	}

	fmt.Printf(entrypoint + "\n")

	calltree(os.Stdout, fcalls, entrypoint, "")
}

func silent_calltree(fset *token.FileSet, afps []*ast.File) (string, error) {
	entry_point := "main"
	var tree_buffer bytes.Buffer

	tree_buffer.WriteString(entry_point + "\n")

	func_center, err := new_func_center(fset, afps)

	if err != nil {
		return "", err
	}

	fcalls, err := func_center.get_fcalls()

	if err != nil {
		return "", err
	}

	calltree(&tree_buffer, fcalls, entry_point, "")

	return tree_buffer.String(), nil
}

func silent_dirtree(fset *token.FileSet, afs []ast.File) string {
	var tree_buffer bytes.Buffer

	tree_buffer.WriteString("main\n")

	tree(&tree_buffer, "main", afs, "")

	return tree_buffer.String()
}

// it does not seem to make sense to return errors in recursive functions
func calltree(writer io.Writer, fcalls []fcall, parent_func string, prefix string) {
	// use fcalls as lookup table to get fcall from fname
	parent_fcall, err := get_fcall_from_slice(fcalls, parent_func)

	if err != nil {
		panic(err)
	}

	fcall_children := parent_fcall.get_children(fcalls)

	for i, v := range fcall_children {
		if i == (len(fcall_children) - 1) {
			fmt.Fprintf(writer, prefix+"`--")
		} else {
			fmt.Fprintf(writer, prefix+"|--")
		}
		fmt.Fprintf(writer, "%s", v.call_name)

		if parent_func == v.call_name {
			fmt.Fprintf(writer, "->%s (recursive)\n", v.call_name)
			continue
		}

		fmt.Fprintf(writer, "\n")

		if i == (len(fcall_children) - 1) {
			calltree(writer, fcalls, v.call_name, prefix+"   ")
		} else {
			calltree(writer, fcalls, v.call_name, prefix+"|  ")
		}
	}
}

func tree(writer io.Writer, parent_func string, afs []ast.File, prefix string) {
	parent_fd := get_funcdecl_from_fname_multifile(afs, parent_func)

	if parent_fd == nil {
		log.Fatal("parent_func does not have a FuncDecl.")
	}

	callexprs := get_calls_from_node(parent_fd)

	callexprs = filter_calls(callexprs, afs)

	calls := callexprs_to_strings(callexprs)

	for i, v := range calls {
		if i == (len(calls) - 1) {
			fmt.Fprintf(writer, prefix+"`--")
		} else {
			fmt.Fprintf(writer, prefix+"|--")
		}
		fmt.Fprintf(writer, "%s", v)

		if parent_func == v {
			fmt.Fprintf(writer, "->%s (recursive)\n", v)
			continue
		}

		fmt.Fprintf(writer, "\n")

		if i == (len(calls) - 1) {
			tree(writer, v, afs, prefix+"   ")
		} else {
			tree(writer, v, afs, prefix+"|  ")
		}
	}
}
