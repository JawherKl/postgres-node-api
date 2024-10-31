package main

import (
    "fmt"
    "time"
)

func producer(ch chan<- int) {
    for i := 0; i < 5; i++ {
        ch <- i
        fmt.Println("Produced:", i)
        time.Sleep(200 * time.Millisecond)
    }
    close(ch)
}

func consumer(ch <-chan int) {
    for num := range ch {
        fmt.Println("Consumed:", num)
        time.Sleep(300 * time.Millisecond)
    }
}

func main() {
    ch := make(chan int, 3) // Buffered channel with capacity of 3

    go producer(ch) // Start producer goroutine
    consumer(ch)    // Start consumer

    fmt.Println("Main function finished.")
}
