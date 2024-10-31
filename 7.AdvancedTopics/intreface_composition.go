package main

import "fmt"

type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

type IO interface {
    Reader
    Writer
}

type MyIO struct{}

func (m MyIO) Read() string {
    return "Reading data"
}

func (m MyIO) Write(data string) {
    fmt.Println("Writing:", data)
}

func main() {
    var io IO = MyIO{}
    fmt.Println(io.Read())
    io.Write("Hello, World!")
}
