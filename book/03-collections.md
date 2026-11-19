# 3. Slices, maps, and text

When I left the first version of this book, I promised to come back and write about Ruby arrays and hashes. It took a little longer than I expected. The nice thing about a slice is that growing one by one item takes considerably less time than coming back to finish a book.

## Arrays and slices

In Ruby, `books = ["Dune", "Solaris"]` makes an array that can grow. In Go, `[2]string{"Dune", "Solaris"}` is an array of exactly two strings. The length is part of the type and assigning the array copies its elements. Most lists you pass around will be slices, written `[]string{"Dune", "Solaris"}`. A slice describes part of an underlying array. It has a length and a capacity, and copying the slice value copies the description while sharing storage.

```go
books := []string{"Dune", "Solaris", "Kindred"}
first := books[:2]
first[0] = "Neuromancer"
fmt.Println(books[0]) // Neuromancer

copyOfFirst := slices.Clone(first)
copyOfFirst[0] = "Dune"
fmt.Println(books[0]) // Neuromancer
```

`append` returns the resulting slice. Assign the result, as in `books = append(books, "Ficciones")`. With spare capacity, append can write into the existing array, which other slices may share. When capacity runs out it allocates a new array. If callers shouldn't be able to affect one another, copy at the ownership boundary. `slices.Clone` is shallow, so elements containing maps or pointers can still share data. A nil slice has length zero and can be appended to. Empty slices usually behave the same way, though JSON encodes a nil slice as `null` and a non-nil empty slice as `[]`.

Use `for i, title := range books` to get index and value. Use `_` to discard an index. The value variable is a copy. To update a slice of structs, use `books[i].Field = value`. Taking a pointer to the loop value gives you a pointer to the copy.

## Hashes and maps

Ruby's `Hash#fetch` gives you a way to distinguish missing keys from stored values. Go's map lookup has an optional second result for exactly that question.

```go
ratings := map[string]int{"Dune": 0}
rating, found := ratings["Dune"]
fmt.Println(rating, found) // 0 true
rating, found = ratings["Solaris"]
fmt.Println(rating, found) // 0 false
```

A nil map is safe to read and panics on assignment. Initialize it with a literal or `make(map[string]int)` before writing. Delete with `delete(ratings, "Dune")`. A key type must support equality, so a string or a comparable struct works. A slice doesn't. Map iteration order is unspecified. Collect and sort the keys for output a reader or a test needs to see in a stable order. Sharing a map between goroutines involves ownership or synchronization, which we'll get to in Chapter 8.

## Bytes, runes, and visible characters

Go strings are immutable sequences of bytes. `len("سلام")` counts 8 bytes and `utf8.RuneCountInString("سلام")` counts 4 Unicode code points. `for offset, r := range text` decodes UTF-8 and gives the byte offset and rune. Given invalid UTF-8 it yields the replacement rune for invalid bytes. Some visible characters consist of several runes, such as a letter plus a combining accent. If your task is cursor movement or trimming to a visible character count, you need Unicode grapheme segmentation. For parsing ASCII delimiters, byte indexing may be exactly what you want. Decide which unit the problem speaks before reaching for a loop.

**Try it.** Count whitespace-separated words with `strings.Fields` and `map[string]int`. Print counts in alphabetical order. Then make two slices share an array, append through one of them with spare capacity, and see which elements changed. The collection demonstration is in `examples/cmd/collections`.
