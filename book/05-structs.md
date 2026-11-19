# 5. Structs, pointers, and methods

If I were writing the reading list in Ruby, I might start with `Book = Struct.new(:id, :title, :read, keyword_init: true)` and add methods when the hash started doing too much. Go has structs and methods too. The way they're put together is slightly different, which is a polite way of saying that my first few guesses were wrong.

```go
type Book struct {
    ID    int64
    Title string
    Read  bool
}

func (b Book) Label() string {
    if b.Read {
        return b.Title + " (read)"
    }
    return b.Title
}

func (b *Book) MarkRead() {
    b.Read = true
}
```

A method declares its receiver before the method name. `Label` receives a copy of the struct. `MarkRead` receives its address and changes the original. Go often inserts the address or dereference for a method call when it has an addressable value. If `book := Book{Title: "Kindred"}`, you can call `book.MarkRead()` directly. A map element is not addressable. Read a value from a map, change it, and assign it back, or deliberately store pointers and document who may mutate them.

Go passes arguments by value, including pointers. Copying a pointer copies the address. `&book` gets the address and `*p` follows it. A nil pointer points nowhere, and dereferencing it panics. A function can return the address of a local value safely. The compiler handles the allocation and lifetime. Start with the API you want callers to use, then measure if allocation is important.

## Give the zero value a job

For `Book`, an empty title is invalid at creation, but the false value of `Read` is a useful default. A constructor function such as `NewBook(title string) (Book, error)` can trim and validate the title. A package can keep fields unexported if callers should always go through that function. Exported fields are useful for plain data transfers, including values we marshal to JSON. The service later validates at the storage boundary because it accepts title strings and constructs the stored record itself.

Embedding can promote fields and methods from another type:

```go
type OwnedBook struct {
    Book
    Owner string
}
```

You can call `owned.Label()` through the embedded `Book`. The embedded field is still a field. Its methods act on that field, and passing an `OwnedBook` where a concrete `Book` is expected requires selecting `owned.Book`. Prefer a named field when promotion would make the API harder to understand.

## Copying has consequences

Copying a struct copies its immediate fields. If it contains a slice, the copied slice can share its backing array. If it contains a mutex that has been used, copying the struct is a bug. These two details explain why you'll see value receivers for simple records and pointer receivers for stateful types. Keep the same receiver style across a type's methods when the type is meant to be used by pointer. Chapter 6 shows how that choice affects interface satisfaction.

**Try it.** Write a `Rename` method that trims whitespace, rejects a blank title, and preserves the old title on error. Call it on a variable and then try to call it on a map element. Explain the compiler error.
