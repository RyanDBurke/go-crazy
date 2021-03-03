package main

import (
	"fmt"
	"runtime"
)

func main() {

	// defers execution until surrounding func is
	// finished
	defer fmt.Println("Hello World!")

	/*
	fmt.Println(sum(10))
	fmt.Println(sum_while(11))

	x := 10
	if x >= 10 {
		fmt.Println(sum_while(x))
	}

	// interesting conditional
	if y:= 1; y > 0 {
		fmt.Println("y")
	}

	fmt.Println(sqrt(2))
	*/

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
}


func sqrt(x float64) (z float64) {
	z = float64(1)

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}

	return
}

func sum(x int) (answer int) {
	answer = 0
	for i := 0; i < x; i++ {
		answer += i
	}
	return
}

func sum_while(x int) int {
	answer := 0

	for x >= 0 {
		answer += x
		x--
	}
	return answer
}