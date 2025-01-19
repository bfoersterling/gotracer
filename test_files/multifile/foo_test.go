package main

import "testing"

func Test_bar(t *testing.T) {
	my_string := bar()

	if my_string != "hello world" {
		t.Fatalf("my_string should be \"hello world\"")
	}
}
