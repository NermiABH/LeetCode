func asteroidCollision(asteroids []int) []int {
    var stack []int

    for _, asteroid := range asteroids {
        stack = append(stack, asteroid)
        for len(stack) > 1 && stack[len(stack)-2] > 0 && stack[len(stack)-1] < 0 { 
            a, b := stack[len(stack)-2], -stack[len(stack)-1]
            if a < b {
                stack[len(stack)-2] = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            }else if a > b {
                stack = stack[:len(stack)-1]
            }else {
                stack = stack[:len(stack)-2]
            }
        }
    }

    return stack
}

//