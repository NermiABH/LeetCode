type RecentCounter struct {
    arr []int
    start int 
}


func Constructor() RecentCounter {
    return RecentCounter{
        arr: make([]int, 0, 100),
    }
}   


func (this *RecentCounter) Ping(t int) int {
    this.arr = append(this.arr, t)

    for this.arr[this.start] < t-3000 {
        this.start++
    }

    return len(this.arr) - this.start
}