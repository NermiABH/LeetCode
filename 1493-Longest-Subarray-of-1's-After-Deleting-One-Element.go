func longestSubarray(nums []int) int {
	m, l, zero := 0, 0, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if zero == -1 {
				zero = i
			} else {
                if m < l {
				    m = l
			    }
				l = i - zero - 1
				zero = i
			}
		} else {
			l++
		}
	}
    if m < l {
		m = l
	}
    if zero == -1 {
        m--
    }
	return m
}