func increasingTriplet(nums []int) bool {
    m1, m2 := math.MaxInt, math.MaxInt
    for _, v := range nums {
        if v <= m1 {
            m1 = v
        } else if v <= m2 {
            m2 = v
        }else {
            return true
        }
    }
    return false
}