package main

func removeElement(nums []int, val int) int {
	var k int
	for _, num := range nums {
		if val != num {
			nums[k] = num
			k++
		}
	}
	return k
}
