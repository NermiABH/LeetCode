package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if length := len(nums1); length%2 == 0 {
		mid := length / 2
		return float64(nums1[mid]+nums1[mid-1]) / 2
	} else {
		return float64(nums1[length/2])
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
