package main

import "strconv"

func fizzBuzz(n int) []string {
	slice := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%5 == 0 && i%3 == 0:
			slice[i-1] = "FizzBuzz"
		case i%5 == 0:
			slice[i-1] = "Buzz"
		case i%3 == 0:
			slice[i-1] = "Fizz"
		default:
			slice[i-1] = strconv.Itoa(i)
		}
	}
	return slice
}
