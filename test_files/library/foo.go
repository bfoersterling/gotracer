package foo

import "fmt"

type foo struct {
	index int
	color string
}

func (f foo) print() {
	fmt.Printf("%+v", f)
}

func (f foo) combine_bar() {
	fmt.Printf(bar() + fmt.Sprintf("%v", f))
}
