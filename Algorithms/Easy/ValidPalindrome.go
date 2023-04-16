package main

import "fmt"

func isPalindrome_(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isAlphanumeric(s[j]) {
			j--
			continue
		}
		if toLower(s[i]) != s[j] {
			fmt.Println(i, j)
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphanumeric(i byte) bool {
	return i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' || i >= '0' && i <= '9'
}
func toLower(i byte) byte {
	if 'A' <= i && i <= 'Z' {
		return i + 32
	}
	return i
}
