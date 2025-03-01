package main

import (
	"testing"
)

func Test_list_uncalled_funcs(t *testing.T) {
	afps, fset, err := parse_dir_afps("test_files/uncalled_funcs")

	if err != nil {
		t.Fatalf("parse_dir_afps failed with err: %v\n", err)
	}

	fc, err := new_func_center(fset, afps)

	if err != nil {
		t.Fatalf("new_func_center failed with err:\n"+
			"%v\n", err)
	}

	result_string := list_uncalled_funcs(fc)

	expected_string := "dont_call_this\n" +
		"never_used"

	if result_string != expected_string {
		t.Fatalf("result_string and expected_string differ.\n"+
			"result_string:\n%v\n\n"+
			"expected_string:\n%v\n", result_string, expected_string)
	}
}
