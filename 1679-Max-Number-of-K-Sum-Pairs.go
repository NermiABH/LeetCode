func maxOperations(nums []int, k int) int {
	n, out := make(map[int]int), 0
	for _, v := range nums {
		diff := k - v
		if count := n[diff]; count == 0 {
			n[v]++
		} else {
			out++
			n[diff]--
		}
	}
	return out
}