// Package declaration
package main

// Importing necessary package
import "fmt"

// main function: entry point of any Go program
func main() {
    // Printing a message to the console
    fmt.Println("Hello, Go!")

    // Declaring and initializing variables
    var greeting string = "Welcome to the Go language!"
    var year int = 2024
    version := 1.18  // Short variable declaration for floats

    // Displaying variable values
    fmt.Println(greeting)
    fmt.Println("The current year is:", year)
    fmt.Printf("Using Go version: %.2f\n", version)

    // Basic types and operations
    number := 10
    increment := 2
    result := number + increment
    fmt.Printf("Result of %d + %d is %d\n", number, increment, result)

    // Demonstrating conditionals
    if year == 2024 {
        fmt.Println("It's 2024!")
    } else {
        fmt.Println("It's not 2024.")
    }

    // Loop example
    for i := 1; i <= 3; i++ {
        fmt.Printf("Loop iteration: %d\n", i)
    }
}
