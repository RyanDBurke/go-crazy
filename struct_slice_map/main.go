package main

import (
	"fmt"
	"golang.org/x/tour/pic"
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

	// arrays
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]				// include low, excludes high
	fmt.Println(s)

	// slices, just references to part of an array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// copy entire array
	c := names[:]
	fmt.Println(c)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// changing the reference will change in array (and all other slices that touch it)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// creating slice w/ make
	a := make([]int, 5)
	printSliceMake("a", a)

	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	// range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	*/

	// stopped @ 18
	pic.Show(Pic)

}

func Pic(dx, dy int) [][]uint8 {
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}