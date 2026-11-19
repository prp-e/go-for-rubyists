# Appendix B. Ruby-to-Go reference

- **Running code.** `ruby main.rb` runs a script. `go run .` builds and runs the current command package. `go build` can leave an executable.
- **Dependencies.** Start at Gemfile in Ruby and `go.mod` in Go. Go import paths identify packages. The module path identifies a versioned collection of packages.
- **Variables.** Ruby locals start with an assignment. A Go function uses `:=` to declare and `=` to update. Use `var` when a zero value or an explicit type reads better.
- **Truth.** Ruby conditions accept objects. Go conditions require booleans. Check a string length, an error value, or a map lookup flag directly.
- **Arrays.** Use a Go slice for growing lists. A Go array has a fixed length as part of its type. A slice can share element storage with another slice.
- **Hashes.** A Go map lookup with two results gives a value and a presence flag. Sort keys when you need deterministic order. Use a lock or ownership when sharing a mutable map across goroutines.
- **Blocks.** Use a function literal as a callback. The parameter type describes what arguments it receives and what results it returns.
- **Exceptions.** Return and check an `error`. Use `%w` to wrap and `errors.Is` or `errors.As` to examine the chain.
- **Ensure.** Use `defer` for cleanup within the current function. Write a small helper if cleanup should happen at the end of each iteration.
- **Objects.** Use structs for state, methods for operations, interfaces for capabilities, and explicit composition for relationships.
- **RSpec or Minitest.** Go's `testing` package recognizes `TestXxx` functions in `_test.go` files. Use `t.Run` for independently named cases.
- **Sinatra routing.** Register HTTP methods and paths on `http.ServeMux`. `httptest` can exercise the mux without binding a public port.
- **Active Record.** Use a database driver with `database/sql` or another database package. Keep transaction work on the same transaction and check the commit error.
- **Async and threads.** Start a goroutine with `go`. Coordinate lifetimes explicitly, synchronize shared state, and pass context to operations that may block.

## Keep these commands close

```sh
go doc net/http.ServeMux
go fmt ./...
go test ./...
go test -race ./...
go vet ./...
go list -m all
go mod tidy
go build ./cmd/readinglist
```

Go's [language specification](https://go.dev/ref/spec) settles language questions. The [module reference](https://go.dev/ref/mod) explains version selection. The [memory model](https://go.dev/ref/mem) is the place to check synchronization guarantees. The [Ruby documentation](https://docs.ruby-lang.org/) is still there when you want to remember what started all of this.
