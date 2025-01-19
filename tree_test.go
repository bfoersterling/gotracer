package main

import (
	"testing"
)

func Test_silent_filetree(t *testing.T) {
	src := `package main
func inner_func() {
	fmt.Println("im inside my_func")
	fmt.Println("inner_func is done")
}
func foo_func() {
	fmt.Println("im foo_func")
	inner_func()
	inner_func()
}
func my_func() {
	foo_func()
	fmt.Println("this was my func")
}
func main() {
	fmt.Println("hi world")

	my_func()

	my_func()

	fmt.Println("main is done")
}`

	expected_result :=
		"|--my_func\n" +
			"|  `--foo_func\n" +
			"|     |--inner_func\n" +
			"|     `--inner_func\n" +
			"`--my_func\n" +
			"   `--foo_func\n" +
			"      |--inner_func\n" +
			"      `--inner_func\n"

	af, fset := string_to_ast(src)

	tree_string := silent_filetree(fset, *af)

	if tree_string != expected_result {
		t.Fatalf("tree_string and expected_result differ.\ntree_string:\n%s\n"+
			"\nexpected_result:\n%s\n", tree_string, expected_result)
	}
}

func Test_silent_dirtree(t *testing.T) {
	// 1
	src_main := `package main
		func main() {
			fmt.Println("hi world")

			my_func()

			my_func()

			fmt.Println("main is done")
		}`

	src_my_func := `package main
		func inner_func() {
			fmt.Println("im inside my_func")
			fmt.Println("inner_func is done")
		}
		func foo_func() {
			fmt.Println("im foo_func")
			inner_func()
			inner_func()
		}
		func my_func() {
			foo_func()
			fmt.Println("this was my func")
		}`

	srcs := []string{src_main, src_my_func}

	expected_result :=
		"main\n" +
			"|--my_func\n" +
			"|  `--foo_func\n" +
			"|     |--inner_func\n" +
			"|     `--inner_func\n" +
			"`--my_func\n" +
			"   `--foo_func\n" +
			"      |--inner_func\n" +
			"      `--inner_func\n"

	afs, fset := strings_to_ast(srcs)

	dirtree := silent_dirtree(fset, afs)

	if dirtree != expected_result {
		t.Fatalf("dirtree and expected_result differ.\n"+
			"dirtree:\n%s\nexpected_result:\n%s\n", dirtree, expected_result)
	}

	// 2
	ast_files, fset, _ := parse_dir("test_files/recursion")

	dirtree = silent_dirtree(fset, ast_files)

	expected_result =
		"main\n" +
			"|--foo\n" +
			"`--rec_func\n" +
			"   |--foo\n" +
			"   |--rec_func->rec_func (recursive)\n" +
			"   `--foo\n"

	if dirtree != expected_result {
		t.Fatalf("dirtree and expected_result differ.\n"+
			"dirtree:\n%s\nexpected_result:\n%s\n", dirtree, expected_result)
	}

	// 3
	ast_files, fset, _ = parse_dir("test_files/make_append")

	dirtree = silent_dirtree(fset, ast_files)

	expected_result =
		"main\n" +
			"|--foo\n" +
			"|  `--second_foo\n" +
			"`--second_foo\n"

	if dirtree != expected_result {
		t.Fatalf("dirtree and expected_result differ.\n"+
			"dirtree:\n%s\n"+
			"\nexpected_result:\n%s\n", dirtree, expected_result)
	}
}

func Test_silent_calltree(t *testing.T) {
	afps, fset, err := parse_dir_afps("test_files/makefile_parser")

	if err != nil {
		t.Fatalf("parse_dir_afps failed with err:\n%v\n", err)
	}

	tree_string, err := silent_calltree(fset, afps)

	if err != nil {
		t.Fatalf("silent_calltree failed with err:\n%v\n", err)
	}

	expected_result := "main\n" +
		"|--parse_cli_args\n" +
		"`--(cli_args).evaluate\n" +
		"   |--is_flag_passed\n" +
		"   |--is_flag_passed\n" +
		"   |--search_result\n" +
		"   |  |--find_makefiles\n" +
		"   |  `--pretty_print_makefiles\n" +
		"   `--single_file_result\n" +
		"      |--parse_makefile\n" +
		"      |  |--file_to_string\n" +
		"      |  |--new_statement\n" +
		"      |  |--(*statement).read_statement\n" +
		"      |  |--(*statement).remove_comment\n" +
		"      |  `--(*statement).parse\n" +
		"      `--has_relevant_target\n" +
		"         `--get_special_targets\n"

	if tree_string != expected_result {
		t.Fatalf("tree_string and expected_result differ.\n"+
			"tree_string:\n%+v\n"+
			"expected_result:\n%v\n", tree_string, expected_result)
	}
}
