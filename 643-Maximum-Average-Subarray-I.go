func findMaxAverage(nums []int, k int) float64 {
    m := 0
    for _, v := range nums[:k] {
        m += v
    }
    l := m
    for i := k; i < len(nums); i++{
        v := l - nums[i-k] + nums[i]
        if  v > m {
            m = v
        }
        l = v
    }
    return float64(m) / float64(k)
}