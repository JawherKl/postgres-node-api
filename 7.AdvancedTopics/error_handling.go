package main

import (
    "errors"
    "fmt"
)

// Custom error
var ErrNegativeNumber = errors.New("negative number not allowed")

func squareRoot(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeNumber
    }
    return x * x, nil
}

func main() {
    result, err := squareRoot(-4)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Square Root:", result)
    }
}
