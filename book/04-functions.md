# 4. Functions, errors, and files

A Ruby method gives you its last expression unless it returns explicitly. Go makes the return part of the contract. Here's the rectangle from the old manuscript, this time with a decision about invalid input.

```go
func rectangleArea(width, height float64) (float64, error) {
    if !(width > 0 && height > 0) {
        return 0, fmt.Errorf("dimensions must be positive")
    }
    area := width * height
    if math.IsInf(area, 0) || area == 0 {
        return 0, fmt.Errorf("area is out of range")
    }
    return area, nil
}
```

The argument list reads as two names of the same type. The result list describes a number and an error. A nil error means the operation succeeded. The number is zero on failure because callers are supposed to inspect the error before using it. In real work, decide what happens for unusual inputs before writing the tests. The comparison here rejects NaN and nonpositive dimensions. The multiplication can also overflow or underflow, even though both inputs passed the first check.

Functions are values. You can accept `func(string) bool` to let a caller choose which titles to keep. A function literal can close over local variables. Unlike passing a Ruby block with `yield`, the callback appears in the signature. For a variadic function, declare its last parameter as `names ...string`. Inside it's a slice. Call it with individual strings or with `names...` to expand a slice.

## Handle errors where you can make a decision

Ruby code often raises an exception and lets a caller rescue it. In Go, most expected failures return an `error` value. When reading a file, the caller might decide that a missing settings file means defaults, and an unreadable one means startup failure. Preserve that choice by wrapping the original error:

```go
data, err := os.ReadFile(path)
if err != nil {
    return fmt.Errorf("read settings %q: %w", path, err)
}
```

`%w` lets a caller use `errors.Is(err, fs.ErrNotExist)` through the wrapper. Use `errors.As` when a concrete error type has fields you need. Keep error strings for humans and use values or types for program decisions. Return errors from libraries. Print one useful message and set the exit status at the command boundary.

## Defer and streaming I/O

For a small file `os.ReadFile` is convenient. For a large input, open it and process it a piece at a time. `defer file.Close()` immediately after a successful open keeps cleanup on every return path. Deferred calls run in reverse order when the enclosing function returns. Arguments to a deferred call are evaluated when it is registered. In a long loop, put per-file work in a helper so deferred closes happen after each file. For a writable file, report flush and close errors because they can mean data failed to reach its destination. `os.Exit` skips deferred calls, so do cleanup inside a helper that returns before calling it.

The command in `examples/cmd/linecount` uses `bufio.Scanner` and checks `scanner.Err()` after the loop. Scanner has a token size limit. This example sets a larger ceiling and still rejects an oversized line rather than reporting a partial count as success. For arbitrary-length records, use a buffered reader and set limits that suit the format.

`panic` unwinds the goroutine and runs deferred calls. An unrecovered panic ends the program. `recover` only stops the panic when called directly by an appropriate deferred function in that goroutine. It's useful at a narrow boundary with a deliberate policy. A missing file is still a normal error.

## Reader and Writer are more useful than a filename

If a function accepts a filename, it decides what the input is and how long it stays open. If it accepts an `io.Reader`, the caller can pass a file, standard input, an HTTP body, or `strings.NewReader` in a test. You can put the parsing loop in a function that accepts a reader and let the command own the file. In Ruby terms this is the part where `StringIO` stops being a trick and starts saving you from creating a directory of test files.

When calling `Read` directly, use the bytes it returned before handling the error. An implementation can return both data and EOF in the same call. Prefer a library that already handles the loop, such as `io.Copy`, when its behavior is what you need. For untrusted or unbounded input, set a limit where the bytes enter your application.

```go
func copyReport(dst io.Writer, src io.Reader) error {
    if _, err := io.Copy(dst, src); err != nil {
        return fmt.Errorf("copy report: %w", err)
    }
    return nil
}
```

The caller still owns both streams. If someone passes a buffered writer, it is that caller's job to flush it and check the flush error. Giving each resource a clear owner makes it easier to see which one wasn't closed when a command runs out of file descriptors.

**Try it.** Count only nonempty lines. Run the command on a nonexistent filename and on a file with one very long line. Make sure both failures are reported.
