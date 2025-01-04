func max(nums ...int) int {
	if len(nums) == 0 {
		panic("max: please provide at least one number")
	}
	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	m := max(candies...)
	resp := make([]bool, len(candies))
	for i, v := range candies {
		if v+extraCandies >= m {
			resp[i] = true
		}
	}
	return resp
}

