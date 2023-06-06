package main

func maxArea(height []int) int {
	var result int
	left, right := 0, len(height)-1
	for left < right {
		result = Max(result, Min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
