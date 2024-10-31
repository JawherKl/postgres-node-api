# Setting Up Go

## Installation

### For macOS:
1. Download the installer from [Go's official page](https://go.dev/dl/).
2. Run the installer and follow the prompts.
3. Add Go to your PATH, if necessary: `export PATH=$PATH:/usr/local/go/bin`

### Verify Installation:
Run the following command to check the Go version:
```bash
go version
```

### Hello World Program
In hello_world.go, write a simple program to print "Hello, World!" to verify the installation.

### 4. **Code Examples**

Create `.go` files within each folder. For instance, in **03_Basics/variables.go**:

```go
// variables.go
package main

import "fmt"

func main() {
    var message string = "Hello, Go!"
    fmt.Println(message)
}
```