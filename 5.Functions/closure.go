package main

import "fmt"

// Function that returns a closure
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    counter := createCounter()

    fmt.Println("Counter:", counter())
    fmt.Println("Counter:", counter())
    fmt.Println("Counter:", counter())
}
