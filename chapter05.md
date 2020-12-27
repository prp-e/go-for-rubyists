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
