package main

import "fmt"

// A simple function that adds two integers
func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(3, 5)
    fmt.Println("Sum:", sum)
}
