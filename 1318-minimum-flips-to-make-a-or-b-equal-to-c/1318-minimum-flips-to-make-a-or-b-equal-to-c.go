func minFlips(a int, b int, c int) int {
	res := 0
	for a > 0 || b > 0 || c > 0 {
		ai, bi, ci := a%2, b%2, c%2
		a, b, c = a/2, b/2, c/2
		if ai|bi != ci {
			res++
			if ci == 0 && ai == 1 && bi == 1 {
				res++
			}
		}
	}
	return res
}