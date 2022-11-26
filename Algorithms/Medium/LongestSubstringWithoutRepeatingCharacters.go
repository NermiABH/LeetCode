package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	var current string
	maxLength := 0
outerLoop:
	for _, val := range s {
		for i2, val2 := range current {
			if val == val2 {
				current = current[i2+1:] + string(val)
				continue outerLoop
			}
		}
		current += string(val)
		if maxLength < len(current) {
			maxLength = len(current)
		}
	}
	return maxLength
}
