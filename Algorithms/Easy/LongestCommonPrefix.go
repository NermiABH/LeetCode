package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		if prefix = commonPrefix(prefix, strs[i]); prefix == "" {
			return ""
		}

	}
	return prefix
}
func commonPrefix(s1, s2 string) string {
	var prefix string
	for i := 0; i < int(math.Min(float64(len(s1)), float64(len(s2)))); i++ {
		if s1[i] != s2[i] {
			break
		}
		prefix += string(s1[i])
	}
	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
