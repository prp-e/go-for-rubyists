# Go for Rubyists

A practical book for people who already know how to program and happen to think in Ruby. We start with a small terminal program. By the end, we'll build and test a JSON API that can shut down without dropping every request on the floor.

**Read the book:** [PDF edition](compiled/go-for-rubyists.pdf) · [Table of contents](SUMMARY.md)

The examples use Go 1.22 language features and the standard library. Use a supported Go release from [go.dev/dl](https://go.dev/dl/). Ruby 3 or newer is enough for the Ruby examples. There's no Rails app to install and no database server to configure for the exercises.

From the repository root:

```sh
cd examples
go test ./...
go run ./cmd/fizzbuzz
go run ./cmd/readinglist
```

The [manuscript](book/) is in Markdown. The [build instructions](tools/README.md) explain how to regenerate the PDF. The 2020 edition is available in the [original history](https://github.com/prp-e/go-for-rubyists/tree/ad46543057d3502feafeff65ff15ded0593f2695).
