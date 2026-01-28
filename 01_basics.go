package main

import (
	"fmt"
	"strings"
)

func Basics() {

	// Numbers
	entero := 10
	decimal := 3.14
	sum := entero + int(decimal)

	fmt.Println(sum) // 13

	// Text
	messageText := "Hola, "
	concat := messageText + "Gentleman"
	inUpperCase := strings.ToUpper(concat)
	fmt.Println(inUpperCase) // HOLA, GENTLEMAN

	// Booleans
	isTrue := true
	fmt.Println(isTrue) // true

	// Arrays  and Slices
	arrays := [3]int{1, 2, 3}
	fmt.Println(arrays) // [1 2 3]

	sliceVariable := []int{4, 5, 6}
	fmt.Println(sliceVariable) // [4 5 6]

	sliceVariable = append(sliceVariable, 7)
	fmt.Println(sliceVariable) // [4 5 6 7]

	// Maps
	dictionary := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(dictionary)

	// Structs
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Kevin", Age: 25}
	fmt.Println(person)
}

// bool true / false - flag o condicionales - default value false
// string cadena de caracteres - para representar texto - default value ""
// int, int8, int16, int32, int64 - entero - controlar el tamaño de los enteros - default value 0
// unit, uint8, uint32, uint64 - entero sin signo - cuando no queremos negativos - default value 0
// float32, float64 - representar valores numéricos reales - números con punto => 32, 64 el sistema - default value 0
// byte === uint8 - trabajar con datos binarios - default value 0
// rune === int32 - cuando queremos representar un solo carácter que ocupa más de un byte - default value 0
// complex64, complex128 - cuando tiene una parte real y una parte imaginaria - N + iN - default value 0 + 0i
