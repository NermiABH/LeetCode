func intersect(nums1 []int, nums2 []int) []int {
    nums1Set := make(map[int]int, len(nums1))
    
    for _, v := range nums1 {
        nums1Set[v]++
    }
    nums2Set := make(map[int]int, len(nums2))
    for _, v := range nums2 {
        nums2Set[v]++
    }
    resp := make([]int, 0, 0)
    for k, v := range nums1Set {
        if v2, ok := nums2Set[k]; ok {
            for count := min(v, v2); count > 0; count-- {
                resp = append(resp, k)
            }
        }
    }
    return resp
}