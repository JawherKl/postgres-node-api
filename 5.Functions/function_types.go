package main

import "fmt"

// Define a function type
type operation func(int, int) int

// Function that performs addition
func add(a int, b int) int {
    return a + b
}

// Function that performs multiplication
func multiply(a int, b int) int {
    return a * b
}

func calculate(op operation, a int, b int) int {
    return op(a, b)
}

func main() {
    sum := calculate(add, 3, 4)
    product := calculate(multiply, 3, 4)

    fmt.Println("Sum:", sum)
    fmt.Println("Product:", product)
}
