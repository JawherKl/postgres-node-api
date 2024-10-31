package main

import "fmt"

// Define an interface
type Animal interface {
    Speak() string
}

// Dog implements the Animal interface
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

// Cat implements the Animal interface
type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    var a Animal

    a = Dog{}
    fmt.Println("Dog:", a.Speak())

    a = Cat{}
    fmt.Println("Cat:", a.Speak())
}
