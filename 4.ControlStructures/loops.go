package main

import "fmt"

func main() {
    // Standard for loop
    for i := 0; i < 5; i++ {
        fmt.Println("Iteration:", i)
    }

    // While-like loop
    count := 0
    for count < 3 {
        fmt.Println("Count is:", count)
        count++
    }

    // Range loop
    fruits := []string{"Apple", "Banana", "Cherry"}
    for index, fruit := range fruits {
        fmt.Printf("Fruit at index %d is %s\n", index, fruit)
    }
}
