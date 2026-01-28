package main

import "fmt"

// * -> Pointer
// & -> Reference (address of memory)
// Razones para usar punteros -> queremos modificar el elemento original

// var a = 1
// referencia de nombre "a" que apunta a un espacio de memoria que contiene en su interior el valor 1

// v1
// COPIA LOS ARGUMENTOS
func increment_v1(num int) {
	num++
}

// v2
func increment_v2(num *int) {
	// NUNCA modifiques el argumento directamente
	*num++
}

func Pointers() {
	value := 10
	// v1
	fmt.Println("Value before increment_v1", value)
	increment_v1(value)
	fmt.Println("Value after increment_v1", value)
	// value after increment is 10
	// Because go makes a copy of Arguments

	// v2
	fmt.Println("Value before increment_v2", value)
	increment_v2(&value)
	fmt.Println("Value after increment_v2", value)

	// new()
	pointer := new(int)                              // puntero int inicializado en 0
	fmt.Println("Initial value with new:", pointer)  // 0x14000182020
	fmt.Println("Initial value with new:", &pointer) // 0x140000a2038
	fmt.Println("Initial value with new:", *pointer) // 0
}
