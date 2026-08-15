package main

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}
// # This is a simple public function for Adding two numbers.
//
// Check out this [link] for in-depth reference.
//
// [link]: https://www.mathsisfun.com/numbers/addition.html
//
func Add[T Number](a, b T) T {
	return a + b
}
