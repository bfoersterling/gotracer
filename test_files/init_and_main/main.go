package main

import (
	"fmt"
	"os"
)

var foo int
var cwd string

// use init func to initialize global vars
func init() {
	var err error
	foo = 3

	cwd, err = os.Getwd()

	if err != nil {
		foobar()
	}
}

func main() {
	foobar()
	fmt.Println(foo)
}
