package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack, pair := make([]rune, 0), map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else if lastIndex := len(stack) - 1; lastIndex == -1 {
			return false
		} else if v == pair[stack[lastIndex]] {
			stack = stack[:lastIndex]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
