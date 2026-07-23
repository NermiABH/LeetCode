func tribonacci(n int) int {
    switch n {
        case 0: return 0
        case 1, 2: return 1
    }

    a, b, c := 0, 1, 1

    n -= 2
    for ; n > 0 ; n-- {
        a, b, c  = b, c, a + b + c
    }

    return c
}