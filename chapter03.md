# Chapter 3 : Flashback to Ruby

In the last section of the previous chapter, I've mentioned that in this chapter, we just review our knowledge of Ruby programming. I don't think it takes so long, specially if you're a rubyist or you have ideas about Ruby. You also may use this chapter as a _quick tutorial_ of Ruby, specially if you have experience in programming and your primary language is not ruby. 

As I mentioned before, Ruby is a little bit too easy to understand, so this chapter obviously is a bunch of long and boring codes. But, in the next chapter we'll see all of them in Go! So, why not just get bored before we see the true magic? 

## Let's cut our Rubies!

In this section, I am going to show you how ruby codes work. Ruby is a scripting language and developers even have provided an [online](https://repl.it/languages/ruby) interpreter which lets you run these codes. To be honest, we won't do so fancy ruby here. That `repl.it` thing is enough to just see the idea. 

Enough talk, let's write the code. 

### Hello World 

```ruby
puts "Hello, World!"
``` 

### Variable declaration 

```ruby
a = 1 
b = "Hello"
``` 

### Conditionals 

```ruby
a = 5
b = 10 

if a > b 
    puts "#{a} is bigger than #{b}"
else
    puts "#{b} is bigger than #{a}"
end
``` 

### Arrays and Hashes 

```ruby
array = [1, 2, 3, 4]
hash = {:a => "a", :b => "b" }
``` 

### Loops 

Fine, here I have to explain different scenarios. All of them pretty much do the same thing, except some of them are directly from the entities like arrays, strings or everything that has a similar behavior. 

#### While loops 

```ruby

counter = 1

while counter < 10 do 
    puts counter
    counter += 1
```

#### For loops 

```ruby

a = [1, 2, 3, 4, 5, 6, 7, 8, 9]

for i in a do
    puts i
end 
``` 

But as I mentioned, we can iterate through the whole thing much easier. Pretty much like this : 

#### Iteration through the array 

```ruby

a = [1, 2, 3, 4, 5, 6, 7, 8, 9]

a.each do |i|
    puts i
end
``` 

And this concept applies to hashes and strings as well. 

### Functions and classes 

For functions and classes, we'll have a separate chapter. I prefer to explain functions and classes of Ruby in the chapter we're coding Go. Why? because Ruby is a _mostly object oriented_ language. But Go is not even object oriented. So I need you to feel the difference. 

## What's next? 