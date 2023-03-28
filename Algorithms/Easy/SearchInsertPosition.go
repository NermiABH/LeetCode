package main

import "fmt"

func searchInsert(nums []int, target int) int {
	first, last := 0, len(nums)-1
	mid := 0
	for first <= last {
		mid = (first + last) / 2
		if current := nums[mid]; current < target {
			first = mid + 1
		} else if current > target {
			last = mid - 1
		} else {
			return mid
		}
	}
	return first
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 4))
}
