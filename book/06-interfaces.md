# 6. Interfaces and generics

In Ruby I can hand an object to a method and hope it responds to `write`. In Go I can ask for the same capability at the function boundary. A type satisfies an interface by having the required methods. It doesn't have to announce the relationship.

```go
func printTitle(w io.Writer, title string) error {
    _, err := fmt.Fprintln(w, title)
    return err
}
```

`io.Writer` asks for `Write([]byte) (int, error)`. Pass `os.Stdout` in the command and a `*bytes.Buffer` in the test. This is a small interface with a big payoff. You can also make an interface in the package that consumes it. If your handler only needs to look up a record, a one-method reader is easier to replace than an interface for half the database.

## Method sets and nil

For a named type `T`, the value method set includes methods with receiver `T`. The pointer method set for `*T` includes those and methods with receiver `*T`. An ordinary method call on an addressable value can borrow its address. Assigning that value to an interface uses its method set as-is. This is why `bytes.Buffer{}` fails as an `io.Writer` while `&bytes.Buffer{}` works.

An interface value can hold a concrete type and a nil pointer at the same time. That interface compares unequal to nil. It's especially easy to trip over this with errors. When a function returning `error` succeeds, return `nil` directly instead of returning a nil pointer of your custom error type. A compile-time assignment such as `var _ io.Writer = (*bytes.Buffer)(nil)` makes an interface expectation explicit, where the expectation is useful to readers.

`any` means the empty interface. Use a type assertion such as `s, ok := value.(string)` when you really do have an open-ended value. Use a type switch for several possible concrete types. Typed structs are easier to use when you know the shape at compile time.

## When the operation is the same for many types

An interface value lets a function call behaviors. A type parameter lets you reuse an implementation across types and retain their static types. Here's a helper for getting a value from a map when a missing entry has a meaningful default:

```go
func getOr[K comparable, V any](
    m map[K]V, key K, fallback V,
) V {
    if value, ok := m[key]; ok {
        return value
    }
    return fallback
}
```

The compiler can infer both parameters from a call such as `getOr(ratings, "Solaris", 0)`. `comparable` allows key equality, which the map requires. For custom numeric constraints, `~int` means types with underlying type `int`, including your own named type. The operations in the function body have to be valid for all types allowed by the constraint. Check `slices`, `maps`, and `cmp` before writing a helper of your own. If the one call site reads well as two lines, it can stay two lines.

**Try it.** Print a title into a buffer and compare the bytes with the expected newline. Then use a writer that returns an error and verify the same operation returns it. For the generics exercise, use `getOr` with a map whose key is a named string type.
