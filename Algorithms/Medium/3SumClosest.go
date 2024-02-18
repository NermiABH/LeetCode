package main

import (
	"fmt"
	"math"
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	closed := math.MaxInt
	slices.Sort(nums)
	for i := 0; i < n-2; i++ {
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			diff1, diff2 := absInt(target-closed), absInt(target-sum)
			if diff1 > diff2 {
				if target == sum {
					return sum
				}
				closed = sum
			}
			if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return closed
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(threeSumClosest([]int{
		-1000, -5, -5, -5, -5, -5, -5, -1, -1, -1},
		-7111))
}
