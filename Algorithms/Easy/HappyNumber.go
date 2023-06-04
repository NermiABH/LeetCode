package main

// Подлежит улучшению
func isHappy(n int) bool {
	nums := make(map[int]struct{})
	for n != 1 {
		var sum int
		for ; n != 0; n /= 10 {
			v := n % 10
			sum += v * v
		}
		if _, ok := nums[sum]; ok {
			return false
		}
		nums[sum] = struct{}{}
		n = sum
	}
	return true
}
