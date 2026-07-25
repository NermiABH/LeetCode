func minCostClimbingStairs(cost []int) int {
    var first, second int
    for _, c := range cost {
        first, second = second, min(first, second) + c
    }
    return min(first, second)
}