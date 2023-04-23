package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	res := make([]string, 0, len(nums))
	start, startIndex := nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			if startIndex == i-1 {
				res = append(res, strconv.Itoa(start))
			} else {
				res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i-1]))
			}
			start, startIndex = nums[i], i
		}
	}
	if startIndex == len(nums)-1 {
		res = append(res, strconv.Itoa(start))
	} else {
		res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[len(nums)-1]))
	}
	return res
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}
