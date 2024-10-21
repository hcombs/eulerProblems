package main

import (
	"fmt"
	"math"
	"strconv"
)

func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func question1 () (sum int){
	sum = 0
	for i:=1; i < 1000; i++ {
		sum += If(i%3 == 0 || i%5 == 0, i ,0)
	}
	return
}

func question2 ()(sum int){
	sum = 2
	var term int = 2
	var prev int = 1
	for term < 4000000 {
		term = term + prev
		prev = term - prev
		sum += If(term%2 == 0, term, 0)
	}
	return
}

func isFactor(number int) bool {
	return 600851475143 % number == 0
}

func isPrime(number int) bool{
	for i:= 2; i < int(math.Sqrt(float64(number))); i++ {
		if(number%i == 0) {
			return false
		}
	}
	return true
}

func reverse (str string) string {
	asciiArray := []rune(str)
	for i, j := 0, len(asciiArray)-1; i<j;  i,j = i+1, j-1 {
		asciiArray[i],asciiArray[j] = asciiArray[j],asciiArray[i]
	}
	return string(asciiArray)
}

func isPalindrome(number int) bool {
	str := strconv.Itoa(number)
	return str == reverse(str)
}

func question3 ()(primefactor int){
	primefactor = 0
	var start int = int(math.Floor(math.Sqrt(600851475143)))
	for i:= start; i > 0; i-- {
		if isFactor(i) && isPrime(i){
			primefactor = i
			return
		}
	}

	return
}

func question4() int {
	var ans int = -1
	for i:= 999; i >= 100; i-- {
		for j:= i; j >= 100; j-- {
			if isPalindrome(i*j) && i * j > ans {
				ans = i * j
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(question4())
}
