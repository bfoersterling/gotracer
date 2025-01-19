package main

import "fmt"

func bar() string {
	return "hello world"
}

func foo() {
	fmt.Printf("%v\n", bar())
	fmt.Printf("%v\n", bar())
	another_func()
}

func second_foo() {
	fmt.Println("second foo")
}
