package main

import "fmt"

// Hello App
func main() {
	fmt.Println("Hello, App!")

	// Variables and Data types
	// Variables
	var foo string
	var foo string = "Go is awesome!"
	var foo, bar string = "Go", "is awesome!"
	var (
		foo string = "Go"
		bar string = "is awesome!"
	)
	var foo = "What's my type?"

	// Constants
	const constant = "This is a constant"

	// Data Types
	// String
	var name string = "My name is Go"
	var bio string = `I am statically typed.
										I was designed at Google.`

	// Bool
	var value bool = false
	var isItTrue bool = true

	// Numeric types
	type byte = uint8
	type rune = int32

	var b byte = 'a'
	var r rune = '🍕'

	var f32 float32 = 1.7812 // IEEE-754 32-bit
	var f64 float64 = 3.1415 // IEEE-754 64-bit

	var c1 complex128 = complex(10, 1)
	var c2 complex64 = 12 + 4i

	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}