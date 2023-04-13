package main

func titleToNumber(columnTitle string) int {
	var result int
	for _, c := range columnTitle {
		result = result*26 + int(c-'A'+1)
	}
	return result
}
