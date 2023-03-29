package main

import "fmt"

var pair = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var sum int
	var i = len(s) - 1
	for i > 0 {
		if v1, v2 := s[i-1], s[i]; pair[v1] < pair[v2] {
			sum += pair[v2] - pair[v1]
			i--
		} else {
			sum += pair[v2]
		}
		i--
	}
	if i == 0 {
		return sum + pair[s[0]]
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
