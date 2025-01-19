package main

import (
	"bytes"
	"fmt"
	"go/printer"
	"testing"
)

func Test_get_core(t *testing.T) {
	myvar := "hey"

	myvar_core := get_core(myvar)

	if myvar_core != "string - hey" {
		t.Fatalf("myvar_core should be \"string - hey\"")
	}
}

func Test_string_to_ast(t *testing.T) {
	src := `package main

func main() {
	i := 3
	fmt.Println(i)
}
`

	ast_file, fset := string_to_ast(src)

	compare_str := node_to_string(ast_file, fset)

	if src != compare_str {
		t.Fatalf("src should be equal to compare_str\n")
	}
}

func Test_strings_to_ast(t *testing.T) {
	src_one := `package main

func main() {
	i := 3
	bar()
}
`
	src_two := `package main
func bar() {
	fmt.Println("bar")
}`

	srcs := []string{src_one, src_two}

	afs, _ := strings_to_ast(srcs)

	if (afs[0].FileStart != 1) || (afs[1].FileStart != 47) {
		t.Fatalf("FileStarts of ast.Files do not fit.\n")
	}
}

func Test_strmap_to_ast(t *testing.T) {
	// 1
	srcs := make(map[string]string, 0)

	srcs["main.go"] = `package main

func main() {
	println("hi")
	foo()
}`

	srcs["foo.go"] = `package main

func foo() {
	println("foo")
}`

	afs, fset, err := strmap_to_ast(srcs)

	if err != nil {
		t.Fatalf("Function strmap_to_ast should not return an error in this test.\n")
	}

	if len(afs) != 2 {
		t.Fatalf("There should be exactly two ast files in afs.\n")
	}

	if fset.Base() != 98 {
		t.Fatalf("fset.Base() should be 98.\n")
	}
}

func Test_strmap_to_afs(t *testing.T) {
	// 1
	srcs := make(map[string]string, 0)

	srcs["main.go"] = `package main

func main() {
	println("hi")
	foo()
}`

	srcs["foo.go"] = `package main

func foo() {
	println("foo")
}`

	afs, fset, err := strmap_to_afs(srcs)

	if err != nil {
		t.Fatalf("Function strmap_to_ast should not return an error in this test.\n")
	}

	if len(afs) != 2 {
		t.Fatalf("There should be exactly two ast files in afs.\n")
	}

	if fset.Base() != 98 {
		t.Fatalf("fset.Base() should be 98.\n")
	}
}

func Test_get_funcdecls_from_afs(t *testing.T) {
	// 1
	src_one := `package main

func main() {
	i := 3
	joe := person{
		"Joe",
		44,
	}

	joe.greet()
}
`
	src_two := `package main
type person struct {
	name string
	age int
}
func (p person) greet() {
	fmt.Println("Hello, ", p.name)
}`

	srcs := []string{src_one, src_two}

	afs, _ := strings_to_ast(srcs)

	fds := get_funcdecls_from_afs(afs)

	if fds[0].Recv != nil {
		t.Fatalf("main function should not have receivers.")
	}

	if fds[1].Name.String() != "greet" {
		t.Fatalf("Second FuncDecl should be \"greet\".")
	}

	if fmt.Sprintf("%v", fds[1].Recv.List[0].Type) != "person" {
		t.Fatalf("Receiver type should be person.")
	}

	// 2 - use source of makefile_parser

	afs, fset, _ := parse_dir("test_files/makefile_parser")

	fds = get_funcdecls_from_afs(afs)

	expected_fn := "evaluate"

	var fname bytes.Buffer

	printer.Fprint(&fname, fset, fds[1].Name)

	if fname.String() != expected_fn {
		t.Fatalf("fname.String() and expected_fn differ.\n"+
			"fname.String():\n%v\n"+
			"expected_fn:\n%v\n", fname.String(), expected_fn)
	}
}
