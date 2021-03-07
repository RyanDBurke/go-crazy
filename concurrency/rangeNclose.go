package main

import (
	"fmt"
	"time"
)

func seq_fibonacci(n int) []int {
	var res []int
	x, y := 0, 1
	for i := 0; i < n; i++ {
		res = append(res, x)
		x, y = y, x+y
	}

	return res
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // close channel
}

func main() {

	// concurrency
	cStart := time.Now()
	c := make(chan int, 100000000)
	go fibonacci(cap(c), c)
	cEnd := time.Now()
	cElapsed := cEnd.Sub(cStart)
	fmt.Println("concurrency:\t",cElapsed)

	// sequential
	sStart := time.Now()
	seq_fibonacci(100000000)
	sEnd := time.Now()
	sElapsed := sEnd.Sub(sStart)
	fmt.Println("sequential:\t", sElapsed)
}
