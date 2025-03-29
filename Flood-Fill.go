func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    oldColor := image[sr][sc]
    if oldColor == color {
        return image
    }

    m, n := len(image), len(image[0])
    stack := [][2]int{{sr, sc}}
    image[sr][sc] = color 

    for len(stack) > 0 {
        i, j := stack[len(stack)-1][0], stack[len(stack)-1][1]
        stack = stack[:len(stack)-1]

        if i-1 >= 0 && image[i-1][j] == oldColor {
            image[i-1][j] = color
            stack = append(stack, [2]int{i - 1, j})
        }
        if i+1 < m && image[i+1][j] == oldColor {
            image[i+1][j] = color
            stack = append(stack, [2]int{i + 1, j})
        }
        if j-1 >= 0 && image[i][j-1] == oldColor {
            image[i][j-1] = color
            stack = append(stack, [2]int{i, j - 1})
        }
        if j+1 < n && image[i][j+1] == oldColor {
            image[i][j+1] = color
            stack = append(stack, [2]int{i, j + 1})
        }
    }

    return image
}
