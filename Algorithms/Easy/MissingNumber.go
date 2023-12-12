package main

import (
	"fmt"
	"slices"
)

func missingNumber(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums[0] ^ 1
	}
	slices.Sort(nums)
	if nums[0] != 0 {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		if (nums[i]&1)^(nums[i+1]&1) == 0 {
			return nums[i] + 1
		}
	}
	return nums[len(nums)-1] + 1
}

func main() {
	//nums := []int{1}
	fmt.Println(1 ^ 1)
}
