package main

import "fmt"

func main() {
    number := 10

    // If-Else
    if number%2 == 0 {
        fmt.Println("Even number")
    } else {
        fmt.Println("Odd number")
    }

    // Switch
    switch number {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
    case 10:
        fmt.Println("Ten")
    default:
        fmt.Println("Other number")
    }
}
