package main

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	var p int
	for x > p {
		p = p*10 + x%10
		x /= 10
	}
	return x == p || x == p/10
}
