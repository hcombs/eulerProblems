# Question summaries and solutions

## 1) Sum all numbers divisible by 3 or 5 under 1000

Loop from 0 to 999 and check if modding 3 or 5 equals 0 and add that to a running total variable. 

## 2) Sum all even numbers in the fibonacci sequence under 4 million

While the fibonacci term is under 4 million, starting at 2, and with a previous term of 1 get the 
next term by adding the previous term plus the current. If this number modded by 2 equals zero add that to a 
running total. 

## 3) Find the largest prime factor of 600851475143  

A potential largest factor for any number can not exceed the square root of that number, so by looping
from the square root rounded down to the nearest int to 0 check if the number provided in the question 
modded by current increment equals zero. If yes, then test if that number is prime by looping from 2 
to the square root of the current increment rounded down to the nearest int ensuring that no number 
in that loop mods the square root to zero. 

## 4) Find the largest number that is the same backwards and forwards (palindrome) from multiplying two 3 digit numbers

Using an outer and inner for loop going from 999 to 100, multiply the outer and inner loop incrementors. 
Convert the product to a string and reverse it. If the reverse equals the original then we have a palindrome.
Store the palindrome product in a variable and change it if you encounter a bigger palindrome while the loop 
is executing.
