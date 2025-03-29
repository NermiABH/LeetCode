func maxProfit(prices []int) int {
    mProfit := 0
    for i, j := 0, 1; j < len(prices); j++ {
        l, r := prices[i], prices[j]
        if l > r {
            i = j
        }else {
            mProfit = max(mProfit, r-l)
        }
    }
    return mProfit
}