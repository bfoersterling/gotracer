package main

import (
	"testing"
)

func Test_get_all_fds(t *testing.T) {
	// 1
	afs, _, err := parse_dir("test_files/makefile_parser")

	if err != nil {
		t.Fatalf("parse_dir returned error: %v\n", err)
	}

	fds := get_all_fds(afs)

	if len(fds) != 17 {
		t.Fatalf("There should be 17 function declarations in the makefile_parser source.\n")
	}
}

func Test_node_to_string(t *testing.T) {
}

func Test_is_selectorstr(t *testing.T) {
	// 1 - positive test
	if !is_selectorstr("&{fmt Println}") {
		t.Fatalf("&{fmt Println} should be detected as selector string\n.")
	}

	// 2 - negative test
	if is_selectorstr("&{fmtPrintln}") {
		t.Fatalf("&{fmtPrintln} should NOT be detected as selector string\n.")
	}

	// 3 - negative test
	if is_selectorstr("") {
		t.Fatalf("Empty string should NOT be detected as selector string\n.")
	}

	// 4 - negative test
	if is_selectorstr("&{}") {
		t.Fatalf("Squirly brackets with empty string should NOT be detected as selector string\n.")
	}
}

func Test_split_selectorstr(t *testing.T) {
	// 1
	selector_str := "&{fmt Println}"

	selector_slice := split_selectorstr(selector_str)

	if selector_slice[0] != "fmt" {
		t.Fatalf("First elem of selector_slice should be fmt.\n")
	}

	if selector_slice[1] != "Println" {
		t.Fatalf("Second elem of selector_slice should be Println.\n")
	}
}

func Test_has_nonstd_import(t *testing.T) {
	// 1
	afps, _, err := parse_dir_afps("test_files/cowsay")

	if err != nil {
		t.Fatalf("parse_dir_afps failed with err:\n"+
			"%s\n", err)
	}

	if !has_nonstd_import(afps) {
		t.Fatalf("A non std import should be detected.\n")
	}

	// 2
	afps, _, err = parse_dir_afps("test_files/makefile_parser")

	if err != nil {
		t.Fatalf("parse_dir_afps failed with err:\n"+
			"%s\n", err)
	}

	if has_nonstd_import(afps) {
		t.Fatalf("No non std import should be detected.\n")
	}
}
