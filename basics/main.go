package main

import (
	"fmt"
	_ "time"
	_ "math"
	_ "math/rand"
)

func main() {
	/*

	// time
	fmt.Println("The time is ", time.Now())

	// math
	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))
	fmt.Println(math.Pi)

	// add func
	fmt.Println(add(10, 2))

	// swap func
	a, b := swap("hi", "ryan")
	fmt.Println(a, b)

	// split func
	fmt.Println(split(17))
	
	// Type
	var name string = "ryan"
	fmt.Printf("Type: %T\n", name)
	*/

	// constants
	const World = "World"
	fmt.Println("Hello", World)

}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) (int, int) {
	return (x + 2), (y + 2)
}
