package main

import (
	"fmt"
	"math"
)

// las interfaces NO se implementan al menos manualmente
// las interfaces se cumplen

type Forma interface {
	Area() float64
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func printArea(f Forma) {
	fmt.Printf("El area es: %.2f\n", f.Area())
}

func interfaces() {
	c := Circulo{Radio: 5}

	printArea(c)
}
