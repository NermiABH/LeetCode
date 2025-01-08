func uniqueOccurrences(arr []int) bool {
    nums := make(map[int]int, len(arr))
    for _, v := range arr {
        nums[v]++
    }
    nn := make(map[int]struct{}, len(nums))
    for _, v := range nums {
        if _, ok := nn[v]; ok{
            return false
        }else {
            nn[v] = struct{}{}
        }
    }
    return true
}