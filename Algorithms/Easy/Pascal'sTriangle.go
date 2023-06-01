package main

func generate(numRows int) [][]int {
	triangle := make([][]int, 1, numRows)
	triangle[0] = []int{1}
	for i := 1; i < numRows; i++ {
		raw := make([]int, i+1)
		raw[0], raw[len(raw)-1] = 1, 1
		for j := 1; j <= i/2; j++ {
			n := triangle[i-1][j-1] + triangle[i-1][j]
			raw[j], raw[len(raw)-1-j] = n, n
		}
		triangle = append(triangle, raw)
	}
	return triangle
}
