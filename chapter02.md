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