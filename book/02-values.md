# 2. Values, decisions, and FizzBuzz

In Ruby I can set `count = 3` and later decide that `count = "three"`. Maybe I had a very good reason. Maybe I forgot the type of a value coming out of JSON. Go wants that decision expressed in the code. It can infer the type at the declaration, but the variable keeps that type afterward.

```go
count := 3
var title string
var active bool
price := float64(count) * 1.25
```

Inside a function, `:=` declares variables and usually gives you all the inference you want. Use `=` to assign to an existing variable. With multiple names on the left, `:=` can reuse names declared in the same block if at least one non-blank name is new. Inside a nested block it can create a new variable with a familiar name. That is called shadowing, and it is a fine way to spend half an hour wondering why a value never changed. At package scope use `var` or `const`. Here the zero values of `title` and `active` are the empty string and false. The zero value of a number is zero.

`int` is a sensible default for counts and indexes. Pick an explicit width when a file format, database representation, or protocol requires it. Go converts between named numeric types explicitly. Integer division truncates toward zero. If you want a mean of 2.5 from 5 and 2, convert before dividing. For money, keep integer minor units when the currency and rounding rules permit it. A floating-point type was trying to do us a favor when it represented 0.1 approximately. The accountant is unlikely to see it that way.

## FizzBuzz again, for science

Ruby's `(1..100).each` gives us an inclusive range. Go has one loop keyword, `for`. Use a three-part loop for initialization, condition, and increment, a condition alone for a while-style loop, or `for range` to iterate (we'll see collections next). In Go 1.22 and newer, `for i := range 100` gives 0 through 99.

```ruby
(1..100).each do |i|
  puts case
       when (i % 15).zero? then "FizzBuzz"
       when (i % 3).zero? then "Fizz"
       when (i % 5).zero? then "Buzz"
       else i
       end
end
```

Here is the Go version's decision-making function. The complete command is in `examples/cmd/fizzbuzz`.

```go
func label(n int) string {
    switch {
    case n%15 == 0:
        return "FizzBuzz"
    case n%3 == 0:
        return "Fizz"
    case n%5 == 0:
        return "Buzz"
    default:
        return strconv.Itoa(n)
    }
}
```

The multiples of 15 come first because those numbers also match the shorter tests. A Go `switch` exits after the chosen case by default. A switch without an expression tests boolean cases in order. `strconv.Itoa` converts an integer to its decimal text. `string(n)` interprets the integer as a Unicode code point, so it would make for a surprising FizzBuzz.

## Conditions and scope

Ruby treats everything except `false` and `nil` as truthy, including zero and empty strings. Go requires a boolean expression. Write `if len(name) == 0` or `if err != nil`. If the condition needs one temporary value, declare it before the condition in the same statement:

```go
if n, err := strconv.Atoi(input); err == nil && n > 0 {
    fmt.Println(n)
}
```

Both variables are local to that `if` and its `else`, if it has one. Short-circuiting means the right side of `&&` is evaluated only if the left side is true. Later we'll handle the parse error separately so the caller can tell the difference between malformed text and a valid number outside the allowed range.

## Constants, named types, and the cost of a conversion

`const secondsPerMinute = 60` is an untyped constant. It can be used where an `int` or a `float64` is needed, provided the value is representable in the destination type. Once you write `seconds := secondsPerMinute`, the local has the inferred type `int`. The constant is more flexible than the variable. A constant can depend on another constant and cannot depend on something computed at runtime.

```go
type BookID int64
type UserID int64

var book BookID = 12
var user UserID = 12
```

The two variables have the same underlying representation and different types. Comparing or assigning one to the other requires an explicit conversion. That moment gives you a chance to ask why a user ID was about to get passed to the book lookup. A type alias, written with an equals sign, gives an existing type another spelling. It doesn't create a distinct type. For a small application a plain integer is often enough. The named type becomes valuable when two values have similar representations and very different jobs.

**Try it.** Turn FizzBuzz into a command that accepts a positive upper bound and prints usage when it receives invalid input. Test 0, 1, 15, 100, and a string.
