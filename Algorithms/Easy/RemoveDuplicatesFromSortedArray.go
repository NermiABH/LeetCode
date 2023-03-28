package main

func removeDuplicates(nums []int) int {
	elem, k := nums[0], 1
	for _, num := range nums {
		if num != elem {
			nums[k], elem = num, num
			k++
		}
	}
	return k
}
