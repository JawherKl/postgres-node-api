package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from Goroutine!")
}

func main() {
    go sayHello() // Start goroutine

    // Wait for the goroutine to finish
    time.Sleep(1 * time.Second)
    fmt.Println("Main function finished.")
}
