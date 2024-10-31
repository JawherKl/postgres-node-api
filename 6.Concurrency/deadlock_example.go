package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    ch1 := make(chan struct{})
    ch2 := make(chan struct{})

    wg.Add(2)

    // Goroutine 1
    go func() {
        defer wg.Done()
        <-ch1 // Wait for signal
        fmt.Println("Goroutine 1 finished.")
    }()

    // Goroutine 2
    go func() {
        defer wg.Done()
        <-ch2 // Wait for signal
        fmt.Println("Goroutine 2 finished.")
    }()

    // Uncomment to create deadlock
    // ch1 <- struct{}{}
    // ch2 <- struct{}{}

    // Resolve deadlock by signaling
    ch1 <- struct{}{}
    ch2 <- struct{}{}

    wg.Wait()
    fmt.Println("Both goroutines completed successfully.")
}
