package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
        Walk(t.Left, ch)
    }
    ch <- t.Value
    if t.Right != nil {
        Walk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {

	// https://tour.golang.org/concurrency/8

	// create channel
	ch := make(chan int, 20)

	// run goroutine
	go Walk(tree.New(1), ch)
}
