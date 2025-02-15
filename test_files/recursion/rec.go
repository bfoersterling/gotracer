package main

import "fmt"

func foo() {
	fmt.Println("foo")
}

func rec_func(bar string) {
	foo()
	if bar == "bar" {
		rec_func("bar")
	}
	foo()
}
