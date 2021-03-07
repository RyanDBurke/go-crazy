package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil	{ Walk(t.Left, ch) }
	ch <- t.Value
    if t.Right != nil 	{ Walk(t.Right, ch) }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	const CAP = 10
	
	ch1 := make(chan int, CAP)
	ch2 := make(chan int, CAP)

	go func() {
		Walk(t1, ch1)
	 	Walk(t2, ch2)
	}()

	defer close(ch1)
	defer close(ch2)

	for i := 0; i < CAP; i++ {
		x := <-ch1
		y := <-ch2
		fmt.Println("x: ", x, ", y: ", y)
		if x != y {
			return false
		}
	}

	return true
}


func main() {

	// create channel
	ch := make(chan int)

	// run goroutine
	go func() {
		defer close(ch)
		Walk(tree.New(1), ch)
	}()

	res := Same(tree.New(2), tree.New(2))
	if res {
		fmt.Println("They are equal")
	} else {
		fmt.Println("They are NOT equal")
	}
}
