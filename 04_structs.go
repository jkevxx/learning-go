package main

import "fmt"

type Person struct {
	Name      string
	Lastname  string
	Age       int32
	IsAproved bool
}

// Method with a receiver (p *Person)
// Think of this like an instance method in Java or a class method in Python/JS.
//
//   - `p` is similar to `this` (Java/JS) or `self` (Python)
//   - `*Person` means we receive the REAL object, not a copy
//     (like passing the object by reference, not by value)
//
// What it does:
// - Toggles IsAproved (true → false, false → true)
// - Updates the struct itself
// - Returns the new value
func (p *Person) ChangeAprove() bool {
	p.IsAproved = !p.IsAproved
	return p.IsAproved
	// fmt.Println("the approve has changed", p.IsAproved)
}

func structs() {
	p1 := Person{
		Name:      "Kevin",
		Lastname:  "Paez",
		Age:       25,
		IsAproved: true,
	}

	fmt.Println("This is the person: ", p1)
	p1.ChangeAprove()
	fmt.Println("This is the person: ", p1)

}
