func maxSubArray(nums []int) int {
    maxSum, currentSum := nums[0], nums[0]
    for _, v := range nums[1:]{
        if currentSum < 0 {
            currentSum = v
        }else {
            currentSum += v
        }
        maxSum = max(currentSum, maxSum)
    }
    return maxSum
}