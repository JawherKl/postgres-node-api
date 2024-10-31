package main

import (
    "fmt"
    "time"
)

func producer(ch chan<- string) {
    for i := 1; i <= 5; i++ {
        msg := fmt.Sprintf("Message %d", i)
        ch <- msg
        time.Sleep(500 * time.Millisecond)
    }
    close(ch) // Close channel after sending all messages
}

func consumer(ch <-chan string) {
    for msg := range ch {
        fmt.Println("Received:", msg)
    }
}

func main() {
    ch := make(chan string)

    go producer(ch) // Start producer goroutine
    consumer(ch)    // Start consumer

    fmt.Println("Main function finished.")
}
