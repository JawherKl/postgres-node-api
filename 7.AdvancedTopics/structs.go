package main

import "fmt"

// Define a struct
type Rectangle struct {
    width, height float64
}

// Method to calculate area
func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func main() {
    rect := Rectangle{width: 10, height: 5}
    fmt.Println("Area of rectangle:", rect.Area())
}
