package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

// method for Vertex struct
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func main() {

	// methods and structs
	v := Vertex {3, 4}
	fmt.Println(v.Abs())

	// Pointer Recievers (4) https://tour.golang.org/methods/4
}