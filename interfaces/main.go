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

// using pointer here allows us to change the values
// within Vertex v!
// can take pointers or values :)
func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

// function HAVE to take pointers if specified
func Add(v *Vertex, f float64) {
	v.x = v.x + f
	v.y = v.y + f
}

func PrintV(v Vertex) {
	fmt.Println(v.x, v.y)
}

// INTERFACES
type PrintName interface {
	PrintName()
}

type Name struct {
	name string
}

func (n Name) PrintName() {
	fmt.Println(n.name)
}

func main() {

	// methods and structs (3)
	v := Vertex {3, 4}
	fmt.Println(v.Abs())

	// Pointer Recievers (4)
	v = Vertex {3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	// pointers & functions (5)
	v = Vertex {2, 3}
	Add(&v, 10)
	PrintV(v)

	// interface (9)
	// a set of method signatures
	// interface can hold any value that implements those methods
	var p PrintName
	myName := Name {"Ryan"}
	p = myName
	p.PrintName()
}