package main

import (
	"fmt"
	_ "math"
)

type I interface {
	M()
}

type T struct {
	S string
}

// handling null pointer exception
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println(t.S)
	}
	
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {

	// declaring
	var i I
	var t *T

	i = t
	i.M()

	i = &T{"Ryan"}
	describe(i)
	i.M()
}	