package main

import "fmt"

func foo() []string {
	my_slice := make([]string, 0)

	my_slice = append(my_slice, "Peter")

	second_foo()

	return my_slice
}

func second_foo() {
	fmt.Println("never mind")
}
