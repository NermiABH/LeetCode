func findDifference(nums1 []int, nums2 []int) [][]int {
    answer := make([][]int, 2)
    answer[0] = make([]int, 0, len(nums1))
    answer[1] = make([]int, 0, len(nums2))
    nums := make(map[int]int, len(nums1) + len(nums2))
    for _, v := range nums1 {
        nums[v] =-1
    }
    for _, v := range nums2 {
        nums[v]+=2
    }
    for k, v := range nums {
        if v==-1{
            answer[0]=append(answer[0],k)
        }else if v%2==0{
            answer[1]=append(answer[1],k)
        }
    }
    return answer
}