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