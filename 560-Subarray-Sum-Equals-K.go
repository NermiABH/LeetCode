func subarraySum(nums []int, k int) int {
    if len(nums) == 0 {
        return 0
    }
    count := 0
    for i := 0; i < len(nums); i++ {
        value := nums[i]
        if value == k {
            count++
        }
        for j := i+1; j < len(nums); j++ {
            value += nums[j]
            if value == k {
                count++
            }
        }
    }
    return count
}