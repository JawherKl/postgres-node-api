package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 2; j++ {
            fmt.Printf("Outer Loop: %d, Inner Loop: %d\n", i, j)
        }
    }
}
