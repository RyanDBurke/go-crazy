package main

import (
	"fmt"
)

var (
	v = Vertex{2, 3}
)

type Vertex struct {
	x int
	y int
}

func main() {
	/*
	i, j := 42, 2701

	p := &i 		// points to i
	fmt.Println(*p)	// prints 42
	*p = 21			// change *p (i.e i) to 21
	fmt.Println(i)	// prints 21

	p = &j			// points to j
	*p = *p / 37	// j / 37
	fmt.Println(j) 	// new j value

	// structs
	vp := &v
	vp.x = 4
	fmt.Println(v.x)
	*/

	// arrays
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]				// include low, excludes high
	fmt.Println(s)
}