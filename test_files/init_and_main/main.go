package main

import "fmt"

var foo int

// use init func to initialize global vars
func init() {
	foo = 3
}

func main() {
	foobar()
	fmt.Println(foo)
}
