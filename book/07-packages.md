# 7. Packages, modules, and tests

I used to look for the Gemfile first when opening someone else's Ruby project. In a Go module, start with `go.mod`. It tells you how packages in this repository are addressed, which Go version the module requires, and which versions of dependencies it requests. `go.sum` records checksums used to authenticate module downloads. Commit both when present. A module without third-party dependencies can have no `go.sum` at all.

## A place for each package

The examples use this layout:

```text
examples/
    go.mod
    cmd/hello/
    cmd/readinglist/
    internal/readinglist/
```

`cmd` is a convention. `internal` has an enforced import restriction. Here it keeps the implementation inside the examples tree. Give packages a short name that makes sense at the call site. `readinglist.NewStore()` says more than `utils.New()`. Avoid cycles. If two packages repeatedly need each other's types, consider keeping those types together or making the consumer declare a narrower interface.

Use `go get module/path@version` to add or update a dependency deliberately. `go mod tidy` reconciles the dependency files with imports, including tests. Tagged major versions starting at v2 ordinarily have a major version suffix in the module path. Go's minimal version selection picks the highest required version for each module path from the graph. Read the diff after an upgrade. `go list -m all` shows selected versions and `go mod why -m module/path` helps explain a dependency's presence.

## A test is an ordinary Go file with a special name

Go discovers files ending in `_test.go` and functions with names such as `TestLabel`. An assertion is commonly just an if and a readable failure. Here's a table for two FizzBuzz cases:

```go
func TestLabel(t *testing.T) {
    for _, tc := range []struct {
        input int
        want  string
    }{
        {3, "Fizz"},
        {15, "FizzBuzz"},
    } {
        t.Run(tc.want, func(t *testing.T) {
            if got := label(tc.input); got != tc.want {
                t.Errorf("label(%d) = %q, want %q",
                    tc.input, got, tc.want)
            }
        })
    }
}
```

Put the test in the same package to reach unexported functions. Use a package with a `_test` suffix when you want to exercise what an external caller can see. `t.Fatal` exits the current test goroutine, so propagate errors from any workers to the test before calling it. Use `t.TempDir()` for temporary fixtures and `t.Cleanup` for test-lifetime resources. Avoid real network dependencies or a process-wide clock in a unit test when a local fake will tell you the same thing.

Run `go test ./...` for all packages in the module, `go test -run TestLabel ./cmd/fizzbuzz` for one test, and `go test -race ./...` to detect races on exercised paths. `go test -cover ./...` reports which statements ran. Ask whether you checked failure modes as well as happy paths. Coverage can tell you that execution visited a line while the assertion missed the bug entirely.

## Turn a bug report into a regression test

Suppose someone says that an empty catalog returns `null` and their client expects a JSON array. Write the smallest test that fails for that report. Compare the returned JSON, or decode it into a slice and check that it is non-nil and empty. Changing the snapshot to use `make([]Book, 0)` fixes the contract at the source. With the test in place, the next person who writes `var snapshot []Book` will learn why that little allocation exists.

When a helper or fixture appears in more than one test, ask whether the cases are still readable at the call site. A helper that calls `t.Helper()` makes failure locations point to its caller. A test using `t.Parallel()` must avoid changing process-wide state and must manage any resources shared with sibling tests. Try running a suspicious package with `go test -count=20 ./internal/readinglist`. A failure that occurs only in a particular order is often telling you something about global state or cleanup.

**Try it.** Add cases for 1, 5, and 30. Put a deliberate bug in the switch and make sure the test catches it. Restore the switch before continuing.
