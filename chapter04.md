# Chapter 4 : FizzBuzz

FizzBuzz, is a commonly asked programming question. It doesn't matter if it's a simple programming class exam or a coding interview in a big corporate. It's simple, but needs a lot of focus and attention. In this chapter, I write it in ruby, then we try to re-write it in Go. So, why should we do this, when we don't know shit about Go? because we're going to learn different parts of a Go application in a _real world_ application and not useless examples. If you want a lot of useless examples, you can just search over the web, and you'll witness how many times people repeated pretty much the same thing, over and over again! 

Fine, instead of being angry, let's do our job. First, we need to know what is FizzBuzz (I assume you know, but I also explain it here.) After that, We will see how we can implement the algorithm in Ruby, and then we re-write that in Go. 

## What's FizzBuzz? 

FizzBuzz is a gimmicky algorithm and determines if a number is divisble by 3, 5 or 15 in a very fun way. We just give the algorithm a bunch of numbers (For example every number from 1 to 100) and then if it's divisible by 3, it outputs _Fizz_. If it's divisible by 5 it outputs _Buzz_. Finally, if it's divisible by 15, it outputs _FizzBuzz_. In most of coding interviews, you're asked to code this. Why? because although it's simple, you can make a world of mistakes while writing the code. 

## Ruby 

```ruby

counter = 1

while counter < 100 do
    if counter % 15 == 0
        puts "FizzBuzz"
    elsif counter % 3 == 0
        puts "Fizz"
    elsif counter % 5 == 0
        puts "Buzz"
    else
        puts counter
    end
end
``` 

Well, this is our ruby code. But we should understand Go is not that easy. But it's still easy enough to have fun with! Before writing our version of _FizzBuzz_, we need to learn alternatives in Go. I really didn't want to make another disaster like previous chapter, so I kept if fun and cool in this chapter. Well, let's find out what we need to learn about Go!

## Before we write the code in Go 

### General code structure 

Remember this : 

```go
package main

import "fmt"

func main(){
    fmt.Println("Hello, World")
}
``` 

So, what's this? Imagine we have a file called `main.go`. The _pacakge_ keyword here determines the name of the file, and what it does in the whole application. The next thing is `import "fmt"`. If you did code in ruby, you know we use `require` to import a library. This is the same thing. I also think if you're from a python backgorund, you're more familiar, because it's the same keyword in python. `"fmt"` is _formatter_. If you want to print something, you need that. And of course, it does more than printing but for now, we just need the print function. 

The part we said `func main(){}`, we declare _main function_. This _main function_ is inherited from _C_ and _C++_, which were the main inspirations for Go. Of course, we need to mind that the one and only Ken Thompson had a great roll in the creation of Go. So, it means we have to deal with a lot of C inheritance here. 

> Although I've always been a Ruby developer, but I also used to work with AVR microcontrollers and Arduino boards. I am really familiar with _C_ and I highly recommend to learn a bit of _C_. It will make learning programming much easier. 

Fine, when we pretty much know what everything does here, we need to just know what `fmt.Println("Hello World")` does. It simply prints out the input text with a `\n` appended at the ent. 

Please, remember this _general code structure_, because it's what we're using all the time in this book and pretty much our career as a Go programmer!

### Conditions and statements

Again, these structures are pretty much similar to _C_. Let's make a simple program that checks three things (I hate this example) : 

1. A is greater than B 
2. A is equal to B
3. A is less than B 

This is written like this in Go: 

```go 
package main

import "fmt"

func main(){
    a := 10
    b := 5

    if a > b {
        fmt.Printf("%d is greater than %d", a, b)
    } else if a == b {
        fmt.Printf("%d is equal to %d", a, b)
    } else {
        fmt.Printf("%d is less than %d", a, b)
    }
}
``` 

Are we cool now? Let's take a deeper look. First, I used `fmt.Printf` because I'm more familiar with it's structure (remember my hardware programming background?) and of course there are better ways to print out integers. But the focus here is on `if {}` statements. The structure is :

```go
if CONDITION {
    STATEMENR
}
``` 

and other stuff like `else if` and `else` are pretty much the same. Remember, Go is more like _C_ than Ruby, so it looks (trust me, it just looks like this) a bit more difficult to understand. 

Now only loops remain... 

### Loops