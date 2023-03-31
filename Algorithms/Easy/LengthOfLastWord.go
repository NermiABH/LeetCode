package main

func lengthOfLastWord(s string) int {
	last := len(s) - 1
	for last >= 0 {
		if s[last] != ' ' {
			break
		}
		last--
	}
	for i := last; i >= 0; i-- {
		if s[i] == ' ' {
			return last - i
		}
	}
	return last + 1
}
