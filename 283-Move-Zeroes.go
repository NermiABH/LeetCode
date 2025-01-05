func moveZeroes(nums []int) {
	noZero := 0
	for i, v := range nums {
		if v != 0{
            nums[noZero] = v
            if noZero != i {
                nums[i] = 0
            }
            noZero++
        }
	}
}
