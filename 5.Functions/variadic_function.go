package main

import "fmt"

// A variadic function that sums multiple integers
func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
    result := sum(1, 2, 3, 4, 5)
    fmt.Println("Total Sum:", result)
}
