package main

import "fmt"

func riskyOperation() {
    defer recoverPanic()
    panic("something went wrong")
}

func recoverPanic() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from:", r)
    }
}

func main() {
    riskyOperation()
    fmt.Println("After risky operation")
}
