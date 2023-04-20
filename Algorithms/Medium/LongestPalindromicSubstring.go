package main

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var palStart, palLen int
	check := func(start, end int) {
		for start >= 0 && end < len(s) && s[start] == s[end] {
			start--
			end++
		}
		if end-start > palLen {
			palStart, palLen = start+1, end-start-1
		}
	}
	for i := 0; i < len(s); i++ {
		check(i, i)
		check(i, i+1)
	}
	return s[palStart : palStart+palLen]
}
