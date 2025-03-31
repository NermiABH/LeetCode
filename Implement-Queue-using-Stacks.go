type MyQueue struct {
    s1, s2 Stack
}


func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    this.s1.Push(x) 
}


func (this *MyQueue) Pop() int {
    if this.s2.Empty() {
        for !this.s1.Empty() {
            this.s2.Push(this.s1.Pop())
        }
    }
    return this.s2.Pop()
}


func (this *MyQueue) Peek() int {
    if this.s2.Empty(){
        for !this.s1.Empty() {
            this.s2.Push(this.s1.Pop())
        }
    }
    return this.s2.Peek()
}


func (this *MyQueue) Empty() bool {
    return this.s1.Empty() && this.s2.Empty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */


type Stack []int

func (s *Stack) Push(v int) {
    *s = append(*s, v)
}


func (s *Stack) Pop() int {
    top := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return top
}


func (s *Stack) Peek() int {
    return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
    return len(*s) == 0
}