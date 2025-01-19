package main

import (
	"slices"
	"testing"
)

func Test_get_gofiles(t *testing.T) {
	go_files := get_gofiles("test_files/multifile")

	expected_slice := []string{"test_files/multifile/foo.go", "test_files/multifile/main.go", "test_files/multifile/third_file.go"}

	if slices.Compare(go_files, expected_slice) != 0 {
		t.Fatalf("\"go_files\" %v does not match \"expected_slice\" %v", go_files, expected_slice)
	}
}
