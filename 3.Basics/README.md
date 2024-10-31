# Basics of Go Programming

Welcome to the **Basics** section! This module covers the foundational elements of Go, from simple syntax to essential control structures, types, and variables. Mastering these basics is key to becoming proficient in Go and understanding the core principles that make it unique.

---

## 📖 Topics Covered

1. **Hello, World!**: Your first Go program.
2. **Variables and Constants**: Declaring and initializing variables.
3. **Data Types**: Understanding basic and composite types in Go.
4. **Control Structures**: Using `if`, `for`, `switch`, and more.
5. **Functions**: Declaring and calling functions.
6. **Error Handling**: Introducing basic error handling.

---

## 🏗 Structure of Go Code

A simple Go program begins with a **package** declaration, followed by imports, and a `main` function that serves as the entry point of the program.

```go
package main  // Declare the main package for executables

import "fmt"   // Import the fmt package for formatted I/O

// Main function, the program entry point
func main() {
    fmt.Println("Hello, World!")
}
```

### Running the Program
To execute the code, save it in a file (e.g., `hello.go`), then run:
```bash
go run hello.go
```

---

## 🔢 Variables and Constants

### Declaring Variables
Variables in Go can be declared in several ways:

- **Explicit declaration with type**:
  ```go
  var name string = "Alice"
  ```
- **Type inference**:
  ```go
  age := 30
  ```

### Declaring Constants
Constants are immutable and declared with the `const` keyword:
```go
const Pi = 3.14
```

### Built-in Data Types
Go includes several fundamental data types:
- **Basic Types**: `int`, `float64`, `string`, `bool`
- **Composite Types**: `array`, `slice`, `map`, `struct`

### Type Conversion
Type conversions are explicit:
```go
var x int = 42
var y float64 = float64(x)
```

---

## 🔄 Control Structures

### 1. **`if` and `else` Statements**
   ```go
   if x > 10 {
       fmt.Println("x is greater than 10")
   } else {
       fmt.Println("x is 10 or less")
   }
   ```

### 2. **Loops (`for`)**
Go has a single looping construct, `for`:
   ```go
   for i := 0; i < 5; i++ {
       fmt.Println(i)
   }
   ```

### 3. **Switch Statement**
   ```go
   switch day := "Monday"; day {
   case "Monday":
       fmt.Println("Start of the work week")
   case "Friday":
       fmt.Println("Almost weekend!")
   default:
       fmt.Println("It's just a regular day")
   }
   ```

### 4. **Defer, Panic, and Recover**
   - **`defer`**: Used to delay the execution of a function until the surrounding function returns.
   - **`panic`**: Terminates the program; used for unrecoverable errors.
   - **`recover`**: Catches panics for graceful handling.

---

## 📦 Composite Types

1. **Arrays and Slices**:
   ```go
   var arr [3]int = [3]int{1, 2, 3}
   slice := []int{4, 5, 6}  // A dynamic array
   ```

2. **Maps**:
   ```go
   colors := map[string]string{"red": "#FF0000", "green": "#00FF00"}
   ```

3. **Structs**:
   Custom data types can be defined using `structs`:
   ```go
   type Person struct {
       Name string
       Age  int
   }
   ```

---

## ✨ Functions

### Defining and Calling Functions
Functions are defined with the `func` keyword:
```go
func add(a int, b int) int {
    return a + b
}
```
You can call this function with:
```go
sum := add(3, 4)
```

### Returning Multiple Values
Go supports multiple return values:
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

---

## 🔍 Error Handling

Go doesn’t have traditional exceptions; instead, functions often return an `error` as the second return value:
```go
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

---

## 🚀 Practice

Each file in this section demonstrates a specific Go concept. Run each example, and try modifying the code to see how Go responds. This hands-on approach will deepen your understanding and prepare you for more advanced topics!

For more details on each concept, see the official [Go documentation](https://golang.org/doc/).
