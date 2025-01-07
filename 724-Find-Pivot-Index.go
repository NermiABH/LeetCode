func pivotIndex(nums []int) int {
    m := 0
    for _, v := range nums {
        m+=v
    }
    l := 0
    for i, v := range nums {
        if l == m - l -v {
            return i 
        }
        l += v
    }
    return -1
}