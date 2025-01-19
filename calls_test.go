package main

import (
	"testing"
)

func Test_get_calls_from_afps(t *testing.T) {
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

	afps, _, err := strmap_to_ast(srcs)

	if err != nil {
		t.Fatalf("strmap_to_ast() failed with err:\n%v\n", err)
	}

	callexprs, err := get_calls_from_afps(afps)

	if err != nil {
		t.Fatalf("get_calls_from_afps() failed with err:\n%v\n", err)
	}

	if len(callexprs) != 2 {
		t.Fatalf("There should be 2 ast.CallExprs in callexprs.\n")
	}
}
