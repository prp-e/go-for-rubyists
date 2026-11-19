# Introduction

In 2012, I went looking for an argument on a Linux forum. You know the kind. Someone says Linux, someone corrects them to GNU/Linux, and by page seven nobody remembers what the original question was. Instead I found a PDF about Go. I decided I'd learn the language when my English got better and I had enough free time. Both things eventually happened. I learned Ruby.

Ruby led me to Sinatra and Rails and gave me a reason to care about backend programming. Years later I watched someone build a little container from scratch in Go. I wanted to understand what happened between running a command and getting a process that thought it had the machine to itself. I usually learn these things by building a small version and finding out where I was wrong. This book uses the same trick.

We'll take things we already know how to do in Ruby and write them again in Go. A Ruby array will get us to a slice, and then the slice will do something the array never warned us about. We'll use a Sinatra route to understand a Go HTTP handler, and work our way up to a service with tests, request limits, cancellation, and graceful shutdown. If you're here from Python or JavaScript, the Ruby examples are small enough to read as you go.

## How to read this book

Run the small examples. Change an input before you read the explanation of what it does. When the compiler complains, read the first error and fix it before investigating the six errors that followed it. Some snippets illustrate one piece of a program. Complete programs live under `examples/` and the chapters tell you where to find them. Commands assume a Unix-style shell unless noted.

The exercises build on each other, but you can jump ahead. If you've already written Go, try the slice exercise in Chapter 3, then go straight to the service in Chapter 11. Appendix A gives you a way to check your work. Appendix B is for the days when you remember that Ruby calls it `fetch` and have completely forgotten what Go calls it.

One more thing. I'm still fond of Ruby. There are days when all I need is twenty lines and a hash. There are other days when I want the compiler to tell me precisely where I handed a string to something that expected a number. Let's find out what those days look like.
