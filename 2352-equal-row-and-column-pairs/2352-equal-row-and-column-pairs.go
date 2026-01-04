func equalPairs(grid [][]int) int {
	n := len(grid)
	m := make(map[string]int, n)

	for i := 0; i < n; i++ {
		b := make([]byte, n*4)
		off := 0
		for j := 0; j < n; j++ {
			v := uint32(grid[i][j])
			
			b[off+0] = byte(v)
			b[off+1] = byte(v >> 8)
			b[off+2] = byte(v >> 16)
			b[off+3] = byte(v >> 24)
			off += 4
		}
		key := unsafe.String(unsafe.SliceData(b), len(b)) 
		m[key]++
	}


	tmp := make([]byte, n*4)
	res := 0

	for c := 0; c < n; c++ {
		off := 0
		for r := 0; r < n; r++ {
			v := uint32(grid[r][c])
			tmp[off+0] = byte(v)
			tmp[off+1] = byte(v >> 8)
			tmp[off+2] = byte(v >> 16)
			tmp[off+3] = byte(v >> 24)
			off += 4
		}
		k := unsafe.String(unsafe.SliceData(tmp), len(tmp)) 
		res += m[k]
	}

	return res
}