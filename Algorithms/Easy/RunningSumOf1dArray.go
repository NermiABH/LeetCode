package main

func runningSum(nums []int) []int {
	for i, val := range nums[:len(nums)-1] {
		nums[i+1] += val
	}
	return nums
}
