# Chapter 5 : Functions

In the [Chapter 3](chapter03.md) I told you that I planned to discuss functions in a different chapter, right? This is it. First of all, I have to say, writing a function in a language like _C_ or _Go_ is a bit difficult. Why? Because you need to difine pretty much everything you want as the input and the output. You'll find out what I say later. In the other hand, this way of writing and declaring functions has something people like to call _type_ safety. I have examples of when Ruby can go mad when you do not provide type safety. 

Does this mean Ruby is bad? of course not. Ruby is the definition of the word love for me, but let's be honest, something like type safety is much easier to achieve in Go.

In this chapter, we'll write two simple functions both in Ruby and Go. Then we run the code. In this chapter, you really experience writing codes which are similar to the real world applications!

## A simple function : Area of a rectangle 

I know this is a simple mathematical operation, but we just need to see how a function works, right? And I tried to keep it as simple as possible. Here, we need an algorithm. The algorithm will become a function and we can add our functions to our codes. Enough said, let's see what algorithm we have. 

The whole algorithm we have is `f(x, y) = x * y`. I know this is not how scientific notation of functions work, but let's take it easy. Here, we assume x is our height and y is our width. So the result is x * y. The function in Ruby will be :

```ruby

def rectangle_area(x, y)
    return x * y 
end
```

And this is very simple, isn't it? In Go we have a more complex but _more accurate_ and _type safe_ structure for writing our functions. The same function in Go will be: 

```go

func rectangleArea(x, y float64) float64 {
	return x * y
}
``` 

Fine, so I think you need explanations for this code. The keyword `func` acts like `def` in Ruby. It's short for _function_ and when you see that, you notice this is a function. Then, we have to name our function (and I called it `rectangleArea`, because Go suggests camel case naming.) and after that, we need to declare _arguments_ in parantheses. As we don't always have integers as widths and heights, we need to declare those variables as _float64_. And here our _type safety_ comes. After parantheses, we have _float64_ again. What's that? That is our output type! We need that in order to get our output correctly. 

Now, we know how it works, let's see how our whole code looks like in Go : 

```go

package main

import "fmt"

func rectangleArea(x, y float64) float64 {
	return x * y
}

func main() {
	x, y := 5.0, 10.0
	fmt.Println(rectangleArea(x, y))
}
```

I think we all know how it works, so we just need to move forward and write the second function. But before that, we need to learn something. 

## Data Types in Go

We all know even in languages like Python, Ruby, PHP and pretty much every scripting language, we have _data types_. The simple way to find out the type of data is using methods provided. For example in Python, if we define `a = 2` and then run `type(a)`, we'll find it's an integer, although we don't tell the interpreter we have an integer. In Ruby we just run `a.class` and so on. 

In Go, although we can declare variables without type as easy as typing `a := 5` we still need to define types. As these types are literally a lot, I highly suggest taking a look at [this](https://golangbyexample.com/all-data-types-in-golang-with-examples/) link. We'll need this link in future, specially when we try to do web stuff using Go!

## Factorial function

Remember the scene from the movie Doctor Strange, when he returned over and over to kick Dormammu's ass? We're going to do the same thing here. But remember that recursive functions are not usually the wisest choice you can make. They easily can ruin your memory. I included this one here, because I like the factorial function. It's simple and it still can show us the whole concept of recursive functions. 

A recursive function, is a function that calls itself. One of the most famous recursive functions is _Fibonacci_. And it's up to you to learn and code that function in Ruby and Go. Here, we're just going to write codes for Factorial. 

Factorial algorithm is very easy. We assume 0 and 1 both have the factorial of 1. You know when you multiply something to 1, result is pretty much the same. So 1 has no effects on the results. Now, we have the algorithm : 

```
n! = n * factorial(n-1)
``` 

This cycle continues until we reach 1 or 0 (depending on how you code the algorithm) and the final number, is the `n!` we've been looking for. 

The code in Ruby is :

```ruby
def factorial(n)
    if n == 0
        return 1
    else
        return n * factorial(n)
    end
end
```

It's simple and it works. Now, it's time to write the whole thing in Go. This time, I only write the function. Printing and other stuff is up to you. 

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    } else {
        return n * factorial(n -1)
    }
}
```

It got much simpler when you keep practicing coding in Go. Functions are very important here, because we're dealing with a language which is not _object oriented_ and it's mostly functional. In the future, we'll learn how these functions can be usueful to us. 

## What's next? 

To be honest, it took more than three hours of thinking to find out what should I write here. Go is a very big language (in terms of features, frameworks, community, etc.) and we covered most of it in these five chapters. I am almost out of content for the rest of the book. 

Do not worry, in the next chapter, we'll take a look at `slice` and `map` data structures in Go. Why? I think most used structures in Ruby are arrays and hash tables. So, why not learning the equivalent in Go? Stay tuned for the next chapter. In the mean time, go ahead and write code for different algorithms with the knowledge you have. 