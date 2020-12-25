# Chapter 2 : Setting up the environment

So now we're going to set our environment up. There are so many different ways to set up a Ruby environment, right? I personally use [RVM](http://rvm.io) but there are more ways such as `rbenv` or even installing ruby from repositories (in case you use Linux) or installing from `brew` (If you have a Mac) and something like that. In this chapter, I just refer to `rvm` and do not use any other Ruby installation system. The Only reason for this is that I just want to show you differences and show you how to _not_ get confused while working with Go. 

## A deep look at RVM 

Fine, RVM stands for _Ruby Version Manager_. It's somehow like `npm` (if you used javascript, you know what I mean) but not completely. It's a tool that installs and manages different versions of Ruby on your computer. For example, if you want Ruby version 2.6.6, you just do this : 

```
rvm install 2.6.6
``` 

And if you want to use that version, you just do this : 

```
rvm use 2.6.6
``` 

For some reasons, I mostly use Ruby 2.6.5, and I want it to be my default. So what should I do? it's easy : 

```
rvm use 2.6.5 --default
``` 

Isn't it beautiful? By the way I believe this way of installing a language isn't cool. But do you know why we have to keep using this way - specially when we use ruby - and not get over it? The answer is most of apps we created, Use different versions of ruby and ruby _gems_ (I think gems are more like `npm` by the way.) and we need to test them out with different versions on our own machine, then we have to install _the right version_ on the host machine. Do we need the same thing in Go? not really. Why? You'll find out later in the book yourself :) 

## Go installation

There are a lot of ways to install Go on your machine. If you're on Windows, [website](http://golang.org) offers an installation file (`msi` or `exe`) which installs everything for you. If you're on macOS, that's the same, but this time the whole thing is just made for _Darwin_ kernel. And if you use Linux, [this](https://golang.org/doc/install) is how you should install the damn thing. Isn't it too easy then? It is. And this is why I kinda like Go, even more than Ruby. 

After installation, just run this : 

```
go version
``` 

and you'll have this response : 

```
go version go1.15.6 darwin/amd64
``` 

And remember, I use macOS, so that `darwin` thing may differ on your machine. 

## Go environment 

Fine, finally we arrived here. Now, I just run this in my terminal : 

```
go env
``` 

What does it do? It just lists a bunch of environment variables usef by Go. Something like this : 

```
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/prp-e/Library/Caches/go-build"
GOENV="/Users/prp-e/Library/Application Support/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/prp-e/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/prp-e/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GCCGO="gccgo"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/25/6jb_0lsx1_x8xcdslr9ns78w0000gn/T/go-build098575207=/tmp/go-build -gno-record-gcc-switches -fno-common"
```

Let's see what we've got here. The most important one is `GOPATH` for now. As you can see it tells us `/Users/prp-e/go` is my `GOPATH`. What does it mean? It means everything we've done is here. So, let's take a look at my own `GOPATH`. It looks like this : 

```
/Users/prp-e/go
└── src
    └── github.com
        └── prp-e
``` 

Inside my `GOPATH`, I have a `src` directory, and as you know it's short for _source_. And I also have `github.com` directory. Which means I host my projects on github. Depending on the services you use, this can be different. GitLab? your own instance of gitea? doesn't matter, just create a folder for each one! 

> Something you should know - specailly if you're a newbie programmer - is that most of companies, startups or organizations may have their own _internal_ git server as well. So you may need a directory for that as well. 

Inside that github folder, I have a `prp-e` directory. That's my github username. So if I have another user or organization, it's recommended to have another folder for that. 

Great! We have our development environment now. Why not creating a _Hello World_ project? 

## Hello World 

Let's do a flashback to Ruby. In Ruby, _Hello World_ was this simple : 

```ruby
puts "Hello, World"
``` 

but here, it's a bit more coding work, but the result is really cool. The code we have here looks like this : 

```go

package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

``` 

but, let's do it more classy. I did these : 

```
cd ~/go
``` 

which navigates to my `GOPATH`, then : 

```
cd src/github.com/prp-e/
mkdir go-for-rubyists-projects
``` 

and yes, there will be a repository for projects we're making here. After that, I did these : 

```
cd go-for-rubyists-projects
mkdir 01_hello
cd 01_hello
touch main.go
``` 

And copied the code I gave you earlier to `main.go` file. After that, I just ran this : 

```
go run main.go
``` 

and this was the result : 

```
 ~/go/src/github.com/prp-e/go-for-rubyists-projects/01_hello/ [master] go run main.go
Hello World!
```

Yes, we ran our very first code written in Go! 

## What's next? 

In the next chapter, we just try to take a look at our knowledge of Ruby programming language. We just do a quick review on Ruby and in future chapters of this book, we're going to re-learn everything we knew from Ruby, but this time in Go!