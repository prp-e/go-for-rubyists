# 1. Hello, Gopher

In Ruby I can ask the interpreter a question before deciding what the program is called: `ruby -e 'puts "Hello, World!"'`. This is a great way to test what I remember about an API. For our first Go program we'll give the code a file, a package, and an entry point. It sounds like a lot of paperwork for hello. The good news is that the paperwork is almost the same for the reading-list server at the end of the book.

## Install it and make a module

Install a supported toolchain from [the official download page](https://go.dev/dl/). Open a fresh terminal and run `go version`. On macOS and Linux, `command -v go` shows which executable your shell picked. In PowerShell use `Get-Command go`. If you see an older installation, check PATH before installing another copy out of spite.

```sh
mkdir hello
cd hello
go mod init example.com/hello
```

The new `go.mod` names the module and records its Go version. A module is a collection of packages versioned together. `example.com/hello` is fine for an experiment. Use the repository's real import path when you intend others to import it. You can put the directory anywhere you normally work. Go still has a GOPATH used for defaults such as caches and installed commands, but this project has its own address in `go.mod`.

## Give your program a front door

Save the following in `main.go`.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

`package main` tells Go we're building a command. `main` is the function called at startup. `fmt` is a standard-library package, and `Println` is exported, which is why the P is capitalized. A Go package is a group of files in a directory. You could put more functions in another file in this directory and call them without a `require`.

```sh
go fmt ./...
go run .
go build -o hello .
./hello
```

On Windows, build with `-o hello.exe` and run `.\hello.exe` from PowerShell. `go run .` builds the current package and runs it. `go build` leaves the executable behind. When you pass a single filename to `go run`, Go uses the named files, which can leave out a helper in the same directory. That's why the period is handy.

## Meet the compiler now, while the stakes are low

Add `import "os"` and leave it unused. Then run the program. Go rejects unused imports, and it rejects unused local variables. That can feel like someone tidying the desk while you're still using it. Over time I found it useful when moving code around. Remove the import and run `go doc fmt.Println` to find out what our call actually returns.

**Try it.** Print two words passed as separate arguments to `Println`. Then replace it with `Printf("%s/%s\n", ...)`. Predict the spaces and newline before running it. The repository version is in `examples/cmd/hello`.
