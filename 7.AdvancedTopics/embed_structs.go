package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person  // Embed Person
    Company string
}

func main() {
    emp := Employee{
        Person: Person{Name: "John", Age: 30},
        Company: "TechCorp",
    }
    fmt.Printf("Name: %s, Age: %d, Company: %s\n", emp.Name, emp.Age, emp.Company)
}
