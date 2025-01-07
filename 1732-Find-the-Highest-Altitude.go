func largestAltitude(gain []int) int {
    m, sum := 0, 0
    for _, v := range gain {
        sum += v
        if m < sum {
            m = sum
        }
    }
    return m
}