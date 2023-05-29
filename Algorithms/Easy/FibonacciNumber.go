package main

func fib(n int) int {
	sum, preSum := 0, 1
	for i := 0; i < n; i++ {
		sum, preSum = sum+preSum, sum
	}
	return sum
}
