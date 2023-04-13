package main

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	sum, preSum := 3, 2
	for i := 4; i <= n; i++ {
		sum, preSum = sum+preSum, sum
	}
	return sum
}

// Recursive method
var climbStairsMap = make(map[int]int)

func climbStairsRec(n int) int {
	if n <= 3 {
		return n
	}
	if k, ok := climbStairsMap[n]; ok {
		return k
	}
	k := climbStairs(n-1) + climbStairs(n-2)
	climbStairsMap[n] = k
	return k
}
