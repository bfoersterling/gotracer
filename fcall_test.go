package main

import (
	"testing"
)

func Test_fcall_get_calls(t *testing.T) {
	// 1

	afps, fset, err := parse_dir_afps("test_files/makefile_parser")

	if err != nil {
		t.Fatalf("parse_dir_afps() failed with err:\n%s\n", err)
	}

	// this step takes 500ms!
	fc, err := new_func_center(fset, afps)

	if err != nil {
		t.Fatalf("new_func_center() failed with err:\n%s\n", err)
	}

	fcalls, err := fc.get_fcalls()

	if err != nil {
		t.Fatalf("fc.get_fcalls() failed with err:\n%s\n", err)
	}

	fcall, err := get_fcall_from_slice(fcalls, "main")

	if err != nil {
		t.Fatalf("get_fcall_from_slice() failed with err:\n%s\n", err)
	}

	main_calls := fcall.get_calls()

	if len(main_calls) != 3 {
		t.Fatalf("There should be 3 calls in main, but there are %v.\n",
			len(main_calls))
	}
}
