package main

import "fmt"

func findJudge(n int, trust [][]int) int {
	if n == 1 && len(trust) == 0 {
		return 1
	}
	people := make([]int, n+1)
	for _, couple := range trust {
		people[couple[0]]++
		people[couple[1]]--
	}
	for i, human := range people {
		if human == 1-n {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
}
