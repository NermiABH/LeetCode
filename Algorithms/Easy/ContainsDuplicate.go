package main

func containsDuplicate(nums []int) bool {
	dup := make(map[int]struct{})
	for _, num := range nums {
		if _, exist := dup[num]; exist {
			return true
		}
		dup[num] = struct{}{}
	}
	return false
}
