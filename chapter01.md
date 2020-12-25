# Chapter 1 : A Journey to the land of Gophers

When Go became a somehow famous language among programmers, I asked myself _Why should I learn it?_. I still ask this, but for now I have some strong answers and the strongest one is that _it helps me find a job easier_. Let's get more realistic about the language and its properties. 

Go, has everything you need at the same time. Primarily it was created to become an alternative to C (but I doubt the whole idea, as most of system programs are still being written in C or C++, and also Rust has been born in pretty much the same time as Go. ) but shortly, Go became a replacement for _every single backend language_. Isn't it beautiful? 

Personally, looked for a way to write a whole new operating system in Go. In [this wiki](https://osdev.org) I found some _bare bones_ and to be honest I even couldn't compile the codes (in my defense though, I was only 16 when I was doing that and I didn't know S*it about Go). After that, one of my friends just made [this](https://github.com/fzerorubigd/tmass) which is a tmux-like tool with Go and I wanted to learn Go, again! Years passed, I just learned how Go works and how does its syntax look like, but never tried to code _seriously_ in that language. I never tried to make a _microservice_ using Go, never tried to make a simple shell or something like that. 

Everything was fine, until I watched [this video](https://www.youtube.com/watch?v=8fi7uSYlOdc). I don't know are you like me or not, but if you're like me, you obviously like to make everything yourself. I always wanted to understand how docker, lxc or any other containerization system works. I personally thought _there's no better way than implementing one myself!_ and of course this idea always worked. For example, I made an operating system kernel when I was 21, just to understand how it works. Or even earlier in my teenage years, I was almost 16, I made an Ubuntu based linux distribution just to understand how Linux distributions are made and how they work. When I watched the video, I realized how cool is Docker, and also, how cool is Go! It was almost a year ago. And then, I had first sparks in my head for seriously learning Go!

Fine, I told you a story just in four paragraphs, and I'm not going to continue, because this is not a work of fiction. This book is going to help you guys migrate from Ruby to Go. So, the rest of the chapter, although it's going to be storytelling, I try to keep it as technical as I can.

## Smoking kills, so does the trend!

It is really good to know that smoking kills. For the love of God, everytime you buy a packet of cigarettes, it's printed in a bold font on the packet! But do you know what also kills? the trend. Let me give you an example. K-POP music (which is short for Korean pop) existed for a long time, but in the past few years, it became _a trend_. So, when you criticize the music or the taste of the fans, you're somehow a Nazi to eyes of the others. 

So why do I tell you this? Let's just take a deep look at the community I'm coming from. I come from Iran. In Iran, for a long long time, _PHP_ was the best! Why? probably because most of the internet was written in PHP back then, facebook used it and there was a lot of resources to learn the language. So, why not? I don't know it was 2004 or 2005, but [Blogfa](http://blogfa.com) became a trendy Iranian blogging service. It was written in _C#_ and _ASP_. This was enough for people to just make a new trend. For years, you couldn't find a job without knowing _C#_. It wasn't good, at least to me. 

Around 2011, Python became trendy here. I know python existed long before. But in 2011, I witnessed some organizations, companies and startups here used Python to create their apps, backends, etc. At least it was something good. For finding a job, you could learn python and django and then you could print money. 

After a few years, when _Go_ became the trend of the world, most of those companies who used _PHP_, did a refactor to _Go_. Now, a new age has been started. You really _need_ to learn Go when this happens. Why? because Go is a new language. It has a lot of potentials to become the next _PHP_ and this means we'll have a web written in Go in a few years. 

Although what I told you in the previous section, wasn't the whole story and it's not the most accurate story about trends, but it gives you the idea about how trends can be killer, too!

## Go in a first glance 

Although I didn't want to put any code here, but hey, why not? This is how _Hello World_ looks like in Go : 

```go
package main 

import "fmt"

func main() {
    fmt.Println("Hello, World")
}
``` 

How does it look like? To me, it looks like _C_ or _C++_. And this is one of the reasons I love the language. Let's write a more complex code. Why not using some functions? Fine, I'll write a code that takes width and height of a rectangle and then gives you the area : 

```go
package main 

import (
    "fmt"
    "math"
)

func rectangle_area(width, height float64) float64 {
    return width * height
}

func main() {
    width := 10 
    height := 22 

    fmt.Printf("Area is %d\n", rectangle_area(width, height))
}

``` 

So, this means this language is a bit harder than Ruby. Is it a bad thing? So if you don't know the fact that codes written in Go can be 1500 times faster than ruby, YES! But if you know the fact, No. You spend more time coding, but you spend less time running your codes. And does it mean Ruby is bad? HELL NO! Ruby is awesome, but you know there's something more awesome. 

## Advantages of using Go

So, I think I told you everything about Go and its advantages, right? But in this section I just list all of them : 

* Modern language, with the strength of _C_ and _C++_. 
* Supported by Google (I know if you're a _FLOSS_ advocate this will piss you off, but do you know having support from a big corporate means what? ) 
* Learning resources 
* Models they introduced in the language for concurrent programming, etc. 
* Packages! (Fine, I think I'm out of words for this section, because having a package manager is pretty much the definition of a _modern_ language. )

And if you still look for more to know, I highly suggest reading [this](https://builtin.com/software-engineering-perspectives/golang-advantages) article. 

## What's next?

In the next chapter, we're going to install Go, and take a look at the differences between installation of Go and Ruby, and we also try to set up a development environment. It won't be a long chapter, but I recommend reading that to not miss my jokes! 