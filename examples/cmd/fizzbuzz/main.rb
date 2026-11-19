(1..100).each do |i|
  puts case
       when (i % 15).zero? then "FizzBuzz"
       when (i % 3).zero? then "Fizz"
       when (i % 5).zero? then "Buzz"
       else i
       end
end
