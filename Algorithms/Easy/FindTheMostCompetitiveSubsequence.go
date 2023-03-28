package main

func maximumWealth(accounts [][]int) int {
	var max int
	for _, val := range accounts {
		if current := sumInts(val); max < current {
			max = current
		}
	}
	return max
}

func sumInts(nums []int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}
