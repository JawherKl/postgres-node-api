package main

import (
    "fmt"
    "time"
)

func generateMessage(ch chan<- string, delay time.Duration) {
    time.Sleep(delay)
    ch <- "Message after " + delay.String()
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go generateMessage(ch1, 2*time.Second)
    go generateMessage(ch2, 1*time.Second)

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received:", msg2)
        }
    }

    fmt.Println("Main function finished.")
}
