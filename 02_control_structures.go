package main

import "fmt"

func ControlStructures() {
	defer fmt.Println("END")

	// conditionals
	// Assertive / negative
	age := 20

	if age < 18 {
		fmt.Println("You are not old enough")
		return
	}

	fmt.Println("You are old enough")

	// Loop classic
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration: %d\n", i)
	}

	// Loop While
	n := 0
	for n < 3 {
		fmt.Printf("Iteration: %d\n", n)
		n++
	}

	//
	n = 0
	for {
		n++

		if n == 5 {
			continue
		}

		fmt.Printf("n on the loop: %d\n", n)

		if n >= 7 {
			break
		}
	}

	// range
	slice := []string{"one", "two", "three"}

	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s \n", index, value)
	}
	for _, value := range slice {
		fmt.Printf("Value: %s \n", value)
	}
}
