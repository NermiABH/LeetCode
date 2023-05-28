package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if value, ok := m[num]; ok && i-value <= k {
			return true
		}
		m[num] = i
	}
	return false
}
