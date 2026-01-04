1func singleNumber(nums []int) int {
2	res := 0
3	for _, n := range nums {
4		res ^= n
5	}
6	return res
7}