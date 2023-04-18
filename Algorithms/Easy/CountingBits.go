package main

import "fmt"

func countBits(n int) []int {
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i/2] + i&1
	}
	return arr
}

func main() {
	fmt.Println(countBits(5))
}
