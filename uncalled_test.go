package main

import (
	"testing"
)

func Test_list_uncalled_funcs(t *testing.T) {
	afps, fset, err := parse_dir_afps("test_files/uncalled_funcs")

	if err != nil {
		t.Fatalf("parse_dir_afps failed with err: %v\n", err)
	}

	result_string := list_uncalled_funcs(fset, afps)

	expected_string := "dont_call_this\n" +
		"main\n" +
		"never_used"

	if result_string != expected_string {
		t.Fatalf("result_string and expected_string differ.\n"+
			"result_string:\n%v\n\n"+
			"expected_string:\n%v\n", result_string, expected_string)
	}
}
