func productExceptSelf(nums []int) []int {
    out, buf := make([]int, len(nums)), 1
    for i := range nums {
        out[i], buf = buf, buf * nums[i]
    }
    buf = 1
    for i := len(nums)-1; i >= 0; i-- {
        out[i], buf = out[i]*buf, buf*nums[i]
    }
    return out
}