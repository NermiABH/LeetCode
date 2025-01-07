func longestOnes(nums []int, k int) int {
	m, n := 0, 0
	for i := range nums {
		if nums[i] == 0 {
			if n == k {
				break
			}
			n++
		}
		m++
	}
	l := m
	for i := l; i < len(nums); i++ {
		if v := nums[i]; v == 0 {
			for j := i - l; j <= i; j++ {
				if nums[j] == 0 {
					if n >= k {
						n--
					} else {
						break
					}
				} else {
					l--
				}
			}
		} else {
			l++
			if m < l {
				m = l
			}
		}

	}
	return m
}
