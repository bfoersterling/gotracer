package main

import (
	"testing"
)

func Test_new_func_center(t *testing.T) {
	// 1
	srcs := make(map[string]string, 0)

	srcs["main.go"] = `package main
func main() {
	cli_args := cli_args {
		true,
	}

	cli_args.parse()

	parse()
}`

	srcs["cli_args.go"] = `package main

type cli_args struct {
	verbose bool
}

func parse() string {
	return "hello"
}

func (args *cli_args)parse() {
}`

	afps, fset, err := strmap_to_ast(srcs)

	if err != nil {
		t.Fatalf("strmap_to_ast() should not return an error in this test case.\n")
	}

	fc, err := new_func_center(fset, afps)

	if err != nil {
		t.Fatalf("new_func_center failed.\n")
	}

	if len(fc.fds) != 3 {
		t.Fatalf("fc.fds should contain 3 func decls.\n")
	}

	if len(fc.func_uses) != 2 {
		t.Fatalf("fc.func_uses should contain 2 uses.\n")
	}

	if len(fc.func_defs) != 3 {
		t.Fatalf("fc.func_defs should contain 3 elements.\n")
	}
}

func Test_get_funcdecl(t *testing.T) {
	// 1
	srcs := make(map[string]string, 0)

	srcs["main.go"] = `package main
import "fmt"
func main() {
	cli_args := cli_args {
		true,
	}
	my_slice := make([]string, 0)

	cli_args.parse()

	parse()

	fmt.Printf("%v\n", my_slice)
}`

	srcs["cli_args.go"] = `package main

type cli_args struct {
	verbose bool
}

func parse() string {
	return "hello"
}

func (args *cli_args)parse() {
}`

	afps, fset, err := strmap_to_ast(srcs)

	if err != nil {
		t.Fatalf("strmap_to_ast() should not return an error in this test case.\n")
	}

	fc, err := new_func_center(fset, afps)

	if err != nil {
		t.Fatalf("new_func_center failed with err:\n%s.\n", err)
	}

	callexprs, err := get_calls_from_afps(afps)

	if err != nil {
		t.Fatalf("get_calls_from_afps() failed with err:\n%s\n", err)
	}

	// 1 - use CallExpr that has a corresponding funcdecl
	fd, err := fc.get_funcdecl(callexprs[1])

	if err != nil {
		t.Fatalf("fc.get_funcdecl() failed with err:\n%s\n", err)
	}

	if fd.Recv == nil {
		t.Fatalf("Second call should be a method and have a Recv field.\n")
	}

	// 2 - use CallExpr that does not have a corresponding funcdecl
	fd, err = fc.get_funcdecl(callexprs[0])

	if err.Error() != "CallExpr was not found in types info." {
		t.Fatalf("fc.get_funcdecl() failed with err:\n%s\n", err)
	}
}

func Test_get_fcalls(t *testing.T) {
	// 1
	srcs := make(map[string]string, 0)

	srcs["main.go"] = `package main
import "fmt"
func main() {
	cli_args := cli_args {
		true,
	}
	my_slice := make([]string, 0)

	cli_args.parse()

	parse()

	fmt.Printf("%v\n", my_slice)
}`

	srcs["cli_args.go"] = `package main

type cli_args struct {
	verbose bool
}

func parse() string {
	return "hello"
}

func (args *cli_args)parse() {
}`

	afps, fset, err := strmap_to_ast(srcs)

	if err != nil {
		t.Fatalf("strmap_to_ast() should not return an error in this test case.\n")
	}

	fc, err := new_func_center(fset, afps)

	if err != nil {
		t.Fatalf("new_func_center failed with err:\n%s.\n", err)
	}

	fcalls, _ := fc.get_fcalls()

	if len(fcalls) != 3 {
		t.Fatalf("3 fcalls should be returned and not %d.\n", len(fcalls))
	}

	for _, fcall := range fcalls {
		if (fcall.call_name == "parse") && (fcall.is_method) {
			t.Fatalf("Call with name \"parse\" should not be a method.")
		}
		if (fcall.call_name == "*cli_args.parse") && (fcall.is_method == false) {
			t.Fatalf("Call with name \"*cli_args.parse\" should be a method.")
		}
	}
}

// benchmarks

func Benchmark_new_func_center(b *testing.B) {
	// 1
	afps, fset, _ := parse_dir_afps("test_files/makefile_parser")

	for i := 0; i < b.N; i++ {
		new_func_center(fset, afps)
	}
}
