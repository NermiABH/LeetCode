func search(nums []int, target int) int {
    
    left, right := 0, len(nums)-1

    for left <= right {
        mid := (left + right)/2
        n := nums[mid]
        if target < n {
            right = mid - 1
        }else if target > n {
            left = mid +1
        }else {
            return mid
        }
    }
    return -1
}