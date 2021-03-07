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
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		Walk(t1, ch1)
		Walk(t2, ch2)
	}()

	defer close(ch1)
	defer close(ch2)

	for i := range ch1 {
		fmt.Println(i)
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

	res := Same(tree.New(1), tree.New(1))
	fmt.Println(res)
}
