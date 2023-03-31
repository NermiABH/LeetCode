package main

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	limit := len(haystack) - needleLen
	for i := 0; i <= limit; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}
