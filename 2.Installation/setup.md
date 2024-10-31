# Installation and Setup of Go

This guide provides step-by-step instructions for installing Go on Windows, macOS, and Linux. By the end, you’ll have a working Go environment to start developing!

---

## 📥 1. Download Go

Go to the official [Go download page](https://golang.org/dl/) and download the installer for your operating system.

---

## 💻 2. Installation Instructions

### Windows

1. Run the downloaded `.msi` installer and follow the installation prompts.
2. By default, Go will install to `C:\Program Files\Go`.
3. After installation, add Go’s `bin` directory to your PATH:
   - Open the **Start Menu** and search for **Environment Variables**.
   - Under **System Variables**, find and select the **Path** variable, then click **Edit**.
   - Add `C:\Program Files\Go\bin` to the list.

### macOS

1. Open a terminal.
2. Install Go using Homebrew:
   ```bash
   brew install go
   ```
3. Alternatively, run the downloaded `.pkg` installer and follow the prompts. Go will install to `/usr/local/go` by default.

### Linux

1. Open a terminal.
2. Download the Go tarball for Linux:
   ```bash
   wget https://golang.org/dl/go1.18.linux-amd64.tar.gz
   ```
3. Extract the archive to `/usr/local`:
   ```bash
   sudo tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz
   ```
4. Add Go’s `bin` directory to your PATH by adding this line to your shell profile (`~/.bashrc`, `~/.zshrc`, etc.):
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```
5. Refresh your profile:
   ```bash
   source ~/.bashrc  # or ~/.zshrc for Zsh users
   ```

---

## 🛠 3. Verify Your Installation

1. Open a terminal or command prompt.
2. Type the following command to verify your Go installation:
   ```bash
   go version
   ```
   You should see output like `go version go1.18 <your_os/arch>`, confirming that Go is installed.

---

## ⚙️ 4. Set Up Your Go Workspace

Go’s workspace is where all Go code resides. Setting up a workspace helps manage projects efficiently.

1. **Create a Workspace Folder** (e.g., `~/go`):
   ```bash
   mkdir ~/go
   ```
2. Set the `GOPATH` environment variable to point to this folder by adding the following to your shell profile:
   ```bash
   export GOPATH=$HOME/go
   ```
3. Confirm that your workspace is set up by creating a basic Go project.

---

## 🧪 5. Test Your Setup

Create a simple “Hello, World!” program to confirm your setup:

1. Inside your workspace folder, create a new file:
   ```bash
   mkdir -p $GOPATH/src/hello
   cd $GOPATH/src/hello
   touch hello.go
   ```

2. Add the following code to `hello.go`:

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World!")
   }
   ```

3. Run the program:
   ```bash
   go run hello.go
   ```

If everything is set up correctly, you should see:
```
Hello, World!
```

---

Congratulations! You’re now ready to start coding with Go. For more information, see the [Go Documentation](https://golang.org/doc/).