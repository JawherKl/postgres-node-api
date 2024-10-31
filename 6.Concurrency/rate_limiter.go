package main

import (
    "fmt"
    "time"
)

func main() {
    rateLimit := time.Tick(1 * time.Second) // One request per second

    for i := 1; i <= 5; i++ {
        <-rateLimit
        fmt.Printf("Request %d sent at %v\n", i, time.Now())
    }
}
