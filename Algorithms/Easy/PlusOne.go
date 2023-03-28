package main

func plusOne(digits []int) []int {
	one := true
	for i := len(digits) - 1; one && i >= 0; i-- {
		digits[i]++
		if digits[i]%10 != 0 {
			one = false
		} else {
			digits[i] = 0
		}
	}
	if one {
		digits = append([]int{1}, digits...)
	}
	return digits
}
