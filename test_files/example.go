package main

import "fmt"

func concat_string(one string, two string) string {
	return one + two
}

func foobar() {
	fmt.Println("doing foo")
	fmt.Println("doing bar")

	united_string := concat_string("foo", "bar")

	fmt.Println(united_string)
}

func main() {
	fmt.Println("Hi")

	foobar()
}
