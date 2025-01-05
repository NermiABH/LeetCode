func maxArea(height []int) int {
	m := 0
	for i, j := 0, len(height)-1; i < j; {
		v := min(height[i], height[j]) * (j - i)
		if v > m {
			m = v
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return m
}