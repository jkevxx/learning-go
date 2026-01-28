package main

import (
	"errors"
	"fmt"
)

// classic function
func sum(a, b int) int {
	return a + b
}

// function that return more than one value
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("it cannot divide for 0")
	}

	cociente := a / b

	return cociente, nil
}

// Closure
func counter() func() int {
	c := 0
	return func() int {
		c++
		return c
	}
}

func Functions() {
	// cociente, error := divide(10, 0)
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }

	// fmt.Println(cociente)

	// Closure
	count := counter()
	fmt.Println("Counter: ", count())
	fmt.Println("Counter: ", count())
	fmt.Println("Counter: ", count())
	fmt.Println("Counter: ", count())
	fmt.Println("Counter: ", count())
	fmt.Println("Counter: ", count())

}
