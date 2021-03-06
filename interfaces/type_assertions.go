package main

import (
	"fmt"
	_"math"
)

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) 			// hello

	s, ok := i.(string)
	fmt.Println(s, ok)		// hello, true

	f, ok := i.(float64)
	fmt.Println(f, ok)		// 0, false

	//f = i.(float64) 		// panic
	//fmt.Println(f)			// no float64

	// note similarities to maps..

	// Type switches
	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}