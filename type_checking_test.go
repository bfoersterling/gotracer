package main

import (
	"testing"
)

func Test_get_type_info(t *testing.T) {
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

	info, err := get_type_info(fset, afps)

	if err != nil {
		t.Fatalf("get_type_info() failed with err:\n%v\n", err)
	}

	if len(info.Uses) != 8 {
		t.Fatalf("There should be 8 Uses maps.\n")
	}

	// 2 - check Defs of test_files/makefile_parser

	afps, fset, err = parse_dir_afps("test_files/makefile_parser")

	if err != nil {
		t.Fatalf("Parsing dir test_files/makefile_parser failed with err:\n"+
			"%v\n", err)
	}

	info, err = get_type_info(fset, afps)

	if len(info.Defs) != 119 {
		t.Fatalf("There should be 119 info.Defs.\n")
	}
}
