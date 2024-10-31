package main

import "fmt"

func main() {
    temperature := 30

    if temperature > 25 {
        fmt.Println("It's a hot day.")
    } else if temperature < 15 {
        fmt.Println("It's a cold day.")
    } else {
        fmt.Println("It's a pleasant day.")
    }
}
