package main

func majorityElement(nums []int) int {
	var elem, count int
	for _, num := range nums {
		if count == 0 {
			elem, count = num, 1
		} else if elem == num {
			count++
		} else {
			count--
		}
	}
	return elem
}
