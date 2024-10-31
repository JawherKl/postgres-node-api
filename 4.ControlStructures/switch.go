package main

import "fmt"

func main() {
    day := "Wednesday"

    switch day {
    case "Monday":
        fmt.Println("It's Monday!")
    case "Tuesday":
        fmt.Println("It's Tuesday!")
    case "Wednesday":
        fmt.Println("It's Wednesday!")
    case "Thursday":
        fmt.Println("It's Thursday!")
    case "Friday":
        fmt.Println("It's Friday!")
    default:
        fmt.Println("It's the weekend!")
    }
}
