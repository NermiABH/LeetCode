func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return nil
    }else if len(nums) == 1 {
        return []string{strconv.Itoa(nums[0])}
    }
    resp := make([]string, 0, 0)

    first := nums[0]
    count := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] - count - first -1 != 0 {
            if count == 0 {
                resp = append(resp, strconv.Itoa(first))
            }else {
                resp = append(resp, strconv.Itoa(first)+\->\+strconv.Itoa(nums[i-1]))
            }
            count = 0
            first = nums[i]
        }else {
            count++
        }
    }
     if count == 0 {
                resp = append(resp, strconv.Itoa(first))
            }else {
                resp = append(resp, strconv.Itoa(first)+\->\+strconv.Itoa(nums[len(nums)-1]))
            }
    
    return resp
}