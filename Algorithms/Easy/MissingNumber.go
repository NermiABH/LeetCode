package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	sum := len(nums) * (len(nums) + 1) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}

func main() {
	//nums := []int{1}
	fmt.Println(1 ^ 1)
}
