package main

import "fmt"

func fibonacci(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	numOfFib := 10
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < numOfFib; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fmt.Println("...")
	fibonacci(c, quit)
}
