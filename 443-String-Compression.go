func compress(chars []byte) int {
	last, lenGroup := 1, 1
	for i := 1; i < len(chars); i++ {
		if chars[i-1] == chars[i] {
			lenGroup++
			continue
		}
		if lenGroup != 1 {
			divisor := 1
			for temp := lenGroup; temp >= 10; temp /= 10 {
				divisor *= 10
			}
			for divisor > 0 {
				digit := lenGroup / divisor
				chars[last] = byte(48 + digit)
				last++
				lenGroup %= divisor
				divisor /= 10
			}
			lenGroup = 1
		}
        chars[last] = chars[i]
		last++
	}
	if lenGroup != 1 {
		divisor := 1
		for temp := lenGroup; temp >= 10; temp /= 10 {
			divisor *= 10
		}
		for divisor > 0 {
			digit := lenGroup / divisor
			chars[last] = byte(48 + digit)
			last++
			lenGroup %= divisor
			divisor /= 10
		}
	}
	return last
}