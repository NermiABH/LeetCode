func longestSubarray(nums []int) int {
    maxLen := 0
	start := 0
    zeroCount := 0
    for end := 0; end < len(nums); end ++ {
        if nums[end] == 0 {
            zeroCount++
        }
        if zeroCount > 1 {
            if nums[start] == 0 {
                zeroCount--
            }
            start++
        }
        l := end - start + 1 - zeroCount 
        if l > maxLen {
            maxLen = l
        }
    }
    if maxLen == len(nums){
        maxLen --
    }
    return maxLen
}