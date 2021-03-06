package main

import (
	"fmt"
	"time"
)

/* built in interface
type error interface {
    Error() string
}
*/

type MyError struct {
	When time.Time
	What string
}


// basically toString method for error prints
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {

	// nil errors mean SUCCESS
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
