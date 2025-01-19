package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"slices"
	"testing"
)

func Test_get_fd_from_pos(t *testing.T) {
	// 1
	srcs := make(map[string]string, 0)

	srcs["main.go"] = `package main
// a comment
func main() {
	cli_args := cli_args {
		true,
	}
	cli_args.parse()
	parse()
}`

	afs, _, err := strmap_to_ast(srcs)

	if err != nil {
		t.Fatalf("strmap_to_ast() should not return an error in this test case.\n")
	}

	fds := get_fds_from_afps(afs)

	fd, err := get_fd_from_pos(fds, 32)

	if err != nil {
		t.Fatalf("get_fd_from_pos failed with error:\n%v\n", err)
	}

	if fd.Name.String() != "main" {
		t.Fatalf("main func should be found at position 32.\n")
	}
}

func Test_get_funcdecl_from_fname(t *testing.T) {
	src := `package main

func foo() {
	fmt.Printf("bar...\n")
}

func main() {
	fmt.Println("hello")
	fmt.Println("bye")
	bar()
}`

	fset := token.NewFileSet()

	file_node, err := parser.ParseFile(fset, "main.go", src, parser.SkipObjectResolution)

	if err != nil {
		t.Fatalf("Parsing file failed.\n%s", err)
	}

	// 1
	test_func_decl := get_funcdecl_from_fname(file_node, "main")

	if test_func_decl.Name.Name != "main" {
		t.Fatalf("test_func_decl.Name.Name is %s\n", test_func_decl.Name.Name)
	}

	// 2
	test_func_decl = get_funcdecl_from_fname(file_node, "foo")

	if test_func_decl.Name.Name != "foo" {
		t.Fatalf("test_func_decl.Name.Name is %s\n", test_func_decl.Name.Name)
	}

	// 3
	test_func_decl = get_funcdecl_from_fname(file_node, "")

	if test_func_decl != nil {
		t.Fatalf("test_func_decl is not nil.\n")
	}
}

func Test_get_funcdecl_from_fname_multifile(t *testing.T) {
	// 1
	first_src := `package main

func foo() {
	fmt.Printf("bar...\n")
}

func main() {
	fmt.Println("hello")
	fmt.Println("bye")
	bar()
}`

	fset := token.NewFileSet()

	file_node_one, err := parser.ParseFile(fset, "main.go", first_src, parser.SkipObjectResolution)

	if err != nil {
		t.Fatalf("Parsing file failed.\n%s", err)
	}

	second_src := `package main
func bar() {
	fmt.Println("foo...\n")
	foo()
	foo()
}`

	file_node_two, err := parser.ParseFile(fset, "bar.go", second_src, parser.SkipObjectResolution)

	if err != nil {
		t.Fatalf("Parsing file failed.\n%s", err)
	}

	file_nodes := []ast.File{*file_node_one, *file_node_two}

	bar_fdecl := get_funcdecl_from_fname_multifile(file_nodes, "bar")

	if bar_fdecl.Name.Name != "bar" {
		t.Fatalf("Function name is not bar.\n%s", err)
	}

	// 2

	main_fdecl := get_funcdecl_from_fname_multifile(file_nodes, "main")

	if main_fdecl.Name.Name != "main" {
		t.Fatalf("Function name is not main.\n%s", err)
	}

	// 3 - how to get a method funcdecl?

	var srcs []string

	srcs = append(srcs, `package main

type cli_args struct {
	verbose bool
}

func parse() string {
	return "hello"
}

func (args *cli_args)parse() {
}

func main() {
	cli_args := []cli_args {
		true,
	}

	cli_args.parse()

	parse()
}`)

	afs, fset := strings_to_ast(srcs)

	fd := get_funcdecl_from_fname_multifile(afs, "&{cli_args parse}")

	if fd == nil {
		t.Fatalf("A FuncDecl should be found.")
	}
	if fd.Recv == nil {
		t.Fatalf("List of receivers should have one element.\n")
	}

	expected_fd := `func (args *cli_args) parse() {
}`

	var fd_print bytes.Buffer
	printer.Fprint(&fd_print, fset, fd)

	if fd_print.String() != expected_fd {
		t.Fatalf("fd_print.String() and expected_fd differ.\n"+
			"fd_print.String(): \n%v\n"+
			"expected_fd: \n%v\n", fd_print.String(), expected_fd)
	}
}

func Test_is_funcdecl(t *testing.T) {
	// 1
	srcs := make([]string, 0)

	srcs = append(srcs, `package main

func foo() {
	fmt.Printf("bar...\n")
}

func main() {
	fmt.Println("hello")
	fmt.Println("bye")
	bar()
}`)

	afs, _ := strings_to_ast(srcs)

	if !is_funcdecl(afs, "foo") {
		t.Fatalf("This function should return true.\n")
	}

	if is_funcdecl(afs, "bar") {
		t.Fatalf("This function should return false.\n")
	}
}

func Test_get_calls(t *testing.T) {
	src := `package foo
func my_cool_func(greeting string) {
	fmt.Println(greeting)
	secret_func(greeting)
}
func main() {
fmt.Println("hello")
i := 3
fmt.Println(i)
my_cool_func("foo")
return 0
}`

	fset := token.NewFileSet()

	file_node, err := parser.ParseFile(fset, "", src, parser.SkipObjectResolution)

	if err != nil {
		t.Fatalf("parser.ParseFile failed - %s\n", err)
	}

	my_funcdecl := get_funcdecl_from_fname(file_node, "main")

	cexprs := get_calls(my_funcdecl)

	if "my_cool_func" != fmt.Sprintf("%s", cexprs[2].Fun) {
		t.Fatalf("Callexpression does not match!\n")
	}
}

func Test_get_calls_from_node(t *testing.T) {
	// 1
	src := `package foo
func my_cool_func(greeting string) {
	fmt.Println(greeting)
	secret_func(greeting)
}
func main() {
names := make([]string, 0)
names = append(names, "Joe")
fmt.Println("hello")
i := 3
fmt.Println(i)
fmt.Println(names)
my_cool_func("foo")
return 0
}`

	af, _ := string_to_ast(src)

	my_funcdecl := get_funcdecl_from_fname(af, "main")

	cexprs := get_calls_from_node(my_funcdecl)

	cexprs_strings := callexprs_to_strings(cexprs)

	expected_string_slice := []string{
		"&{fmt Println}",
		"&{fmt Println}",
		"&{fmt Println}",
		"my_cool_func",
	}

	if slices.Compare(cexprs_strings, expected_string_slice) != 0 {
		t.Fatalf("Slices cexprs_strings and expected_string_slice differ.\n"+
			"cexprs_strings:\n%s\n"+
			"expected_string_slice:\n%s\n", cexprs_strings, expected_string_slice)
	}

	// 2 (test if method is detected as a call)
	var srcs []string

	srcs = append(srcs, `package main

type cli_args struct {
	verbose bool
}

func (args *cli_args)parse() {
}

func main() {
	cli_args := []cli_args {
		true,
	}

	cli_args.parse()
}`)

	afs, _ := strings_to_ast(srcs)

	main_fd := get_funcdecl_from_fname_multifile(afs, "main")

	callexprs := get_calls_from_node(main_fd)

	if fmt.Sprintf("%v", callexprs[0].Fun) != "&{cli_args parse}" {
		t.Fatalf("First elem of callexprs should be &{cli_args parse}\n"+
			"But is >%v<\n", callexprs[0].Fun)
	}

	// 3 (calls from makefile_parser source)
	afs, _, _ = parse_dir("test_files/makefile_parser")

	main_fd = get_funcdecl_from_fname_multifile(afs, "main")

	callexprs = get_calls_from_node(main_fd)

	calls := callexprs_to_strings(callexprs)

	// WARNING: need to get the type of "program_flags"
	// (it is of type "cli_args")
	expected_calls := []string{
		"parse_cli_args",
		"&{program_flags evaluate}",
		"&{log Fatal}",
	}

	if slices.Compare(calls, expected_calls) != 0 {
		t.Fatalf("calls and expected_calls differ.\n"+
			"calls:\n%v\n"+
			"expected_calls:\n%v\n", calls, expected_calls)
	}
}

func Test_get_local_calls_from_node(t *testing.T) {
	src := `package foo
func my_cool_func(greeting string) {
	fmt.Println(greeting)
	secret_func(greeting)
}
func main() {
fmt.Println("hello")
i := 3
fmt.Println(i)
my_cool_func("foo")
my_string := make(string, 0)
my_string = append(my_string, "hello")
fmt.Println(my_string)
return 0
}`

	af, _ := string_to_ast(src)

	my_funcdecl := get_funcdecl_from_fname(af, "main")

	cexprs := get_local_calls_from_node(my_funcdecl)

	if fmt.Sprintf("%v", cexprs[0].Fun) != "my_cool_func" {
		t.Fatalf("cexpr[0].Fun should be my_cool_func.\n"+
			"cexprs[0]:\n%v\n", cexprs[0].Fun)
	}
	if fmt.Sprintf("%v", cexprs[len(cexprs)-1].Fun) != "my_cool_func" {
		t.Fatalf("Last element of cexprs should be the my_cool_func function call.\n"+
			"cexprs[last element]:\n%v\n", cexprs[len(cexprs)-1].Fun)
	}
}

func Test_callexprs_to_strings(t *testing.T) {
	src := `package foo
func my_cool_func(greeting string) {
	fmt.Println(greeting)
	secret_func(greeting)
}
func main() {
fmt.Println("hello")
my_cool_func("foo")
my_string := make(string, 0)
my_string = append(my_string, "hello")
my_cool_func("hi")
fmt.Println(my_string)
return 0
}`

	af, _ := string_to_ast(src)

	my_funcdecl := get_funcdecl_from_fname(af, "main")

	cexprs := get_local_calls_from_node(my_funcdecl)

	calls := callexprs_to_strings(cexprs)

	expected_string_slice := []string{
		"my_cool_func",
		"my_cool_func",
	}

	if slices.Compare(calls, expected_string_slice) != 0 {
		t.Fatalf("The slices calls and expected_string_slice should be equal!\n"+
			"calls:\n%s\n"+
			"expected_string_slice:\n%s\n", calls, expected_string_slice)
	}
}

func Test_get_fname_from_call(t *testing.T) {
	// 1
	ae, err := parser.ParseExpr(`my_function("foo")`)

	if err != nil {
		t.Fatalf("An error occured parsing an Expression.")
	}

	fun_name := get_fname_from_call(ae)

	if fun_name != "my_function" {
		t.Fatalf("fun_name should be my_function")
	}

	// 2
	ae, err = parser.ParseExpr(`i + 3 - 14`)

	if err != nil {
		t.Fatalf("An error occured parsing an Expression.\n%s", err)
	}

	fun_name = get_fname_from_call(ae)

	if fun_name != "" {
		t.Fatalf("fun_name should be an empty string")
	}

	// 3
	ae, err = parser.ParseExpr(`fmt.Println("bar")`)

	if err != nil {
		t.Fatalf("An error occured parsing an Expression.\n%s", err)
	}

	fun_name = get_fname_from_call(ae)

	if fun_name != "&{fmt Println}" {
		t.Fatalf("fun_name is %s.\n", fun_name)
	}
}
