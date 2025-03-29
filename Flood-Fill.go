type Node struct {
    i, j int
    next *Node
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    oldColor := image[sr][sc]
    if oldColor == color {
        return image
    }

    m, n := len(image), len(image[0])
    head := &Node{i: sr, j: sc, next: nil}

    for head != nil {
        i, j := head.i, head.j
        head = head.next
        if image[i][j] != oldColor {
            continue
        }

        image[i][j] = color

        if i-1 >= 0 && image[i-1][j] == oldColor {
            head = &Node{i: i - 1, j: j, next: head}
        }
        if i+1 < m && image[i+1][j] == oldColor {
            head = &Node{i: i + 1, j: j, next: head}
        }
        if j-1 >= 0 && image[i][j-1] == oldColor {
            head = &Node{i: i, j: j - 1, next: head}
        }
        if j+1 < n && image[i][j+1] == oldColor {
            head = &Node{i: i, j: j + 1, next: head}
        }
    }

    return image
}
