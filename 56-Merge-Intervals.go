func merge(intervals [][]int) [][]int {
    if len(intervals) < 2 {
        return intervals
    }

    slices.SortFunc(intervals, func(a, b []int) int {
\t\tif a[0] < b[0] {
\t\t\treturn -1
\t\t} else {
\t\t\treturn 1
\t\t}
\t})

    resp := make([][]int, 0, 0)
    last := intervals[0]
    for _, v := range intervals[1:] {
        if last[0] <= v[0] && v[0] <= last[1] {
            last = []int{last[0], max(v[1], last[1])}
        }else {
            resp = append(resp, last)
            last = v
        }
    } 
    resp = append(resp, last)
    
    return resp
}