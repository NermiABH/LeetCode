1func minFlips(a int, b int, c int) int {
2	res := 0
3	for a > 0 || b > 0 || c > 0 {
4		ai, bi, ci := a%2, b%2, c%2
5		a, b, c = a/2, b/2, c/2
6		if ai|bi != ci {
7			res++
8			if ci == 0 && ai == 1 && bi == 1 {
9				res++
10			}
11		}
12	}
13	return res
14}