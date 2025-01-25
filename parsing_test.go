package main

import (
	"testing"
)

func Test_parse_dir(t *testing.T) {
	// 1
	afs, fset, _ := parse_dir("test_files/multifile")

	if afs[0].FileEnd != 209 {
		t.Fatalf("FileEnd of the ast file should be 209, but is %d.", afs[0].FileEnd)
	}

	if afs[1].FileStart != 210 {
		t.Fatalf("FileStart of the ast file should be 210, but is %v.", afs[1].FileStart)
	}

	if afs[1].Name.Name != "main" {
		t.Fatalf("Package name of the ast file should be main.")
	}

	af_string := node_to_string(&afs[1], fset)

	expected_string := `package main

func main() {
	foo()
	second_foo()
}
`

	if af_string != expected_string {
		t.Fatalf("af_string and expected_string should be equal.\naf_string: %s", af_string)
	}

	// 2 - test if program will stop if mixed packages are detected

	_, _, err := parse_dir("test_files/mixed_packages")

	if err == nil {
		t.Fatalf("An error should be thrown when you have mixed packages.\n")
	}

	if err.Error() != "All files need to be of the same package." {
		t.Fatalf("A different error should be returned.\nerr:\n%s", err.Error())
	}

	// 3 - error when parsing file (no package declaration)

	_, _, err = parse_dir("test_files/parser_error")

	if err == nil {
		t.Fatalf("err should not be nil when parsing an invalid Go file.\n")
	}

	// 4 - parsing dir without go files

	_, _, err = parse_dir("test_files/no_go_files")

	if err == nil {
		t.Fatalf("An error should be thrown when trying to parse a dir without go files.\n")
	}
}

func Test_parse_dir_afps(t *testing.T) {
	// 1
	afps, fset, _ := parse_dir_afps("test_files/multifile")

	if afps[0].FileEnd != 209 {
		t.Fatalf("FileEnd of the ast file should be 209, but is %d.", afps[0].FileEnd)
	}

	if afps[1].FileStart != 210 {
		t.Fatalf("FileStart of the ast file should be 210, but is %v.", afps[1].FileStart)
	}

	if afps[1].Name.Name != "main" {
		t.Fatalf("Package name of the ast file should be main.")
	}

	af_string := node_to_string(afps[1], fset)

	expected_string := `package main

func main() {
	foo()
	second_foo()
}
`

	if af_string != expected_string {
		t.Fatalf("af_string and expected_string should be equal.\naf_string: %s", af_string)
	}

	// 2 - test if program will stop if mixed packages are detected

	_, _, err := parse_dir_afps("test_files/mixed_packages")

	if err == nil {
		t.Fatalf("An error should be thrown when you have mixed packages.\n")
	}

	if err.Error() != "All files need to be of the same package." {
		t.Fatalf("A different error should be returned.\nerr:\n%s", err.Error())
	}

	// 3 - error when parsing file (no package declaration)

	_, _, err = parse_dir_afps("test_files/parser_error")

	if err == nil {
		t.Fatalf("err should not be nil when parsing an invalid Go file.\n")
	}

	// 4 - parsing dir without go files

	_, _, err = parse_dir_afps("test_files/no_go_files")

	if err == nil {
		t.Fatalf("An error should be thrown when trying to parse a dir without go files.\n")
	}
}
