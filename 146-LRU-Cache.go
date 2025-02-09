type LRUCache struct {
	head, tail *Node
	mp map[int]*Node
	cap int
}

type Node struct {
	next, prev *Node
	key, value int
}

func newNode(key, value int) *Node{
	return &Node{
		key: key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	head, tail := newNode(0, 0), newNode(0, 0)
	head.next, tail.prev = tail, head
	return LRUCache {
		head: head,
		tail: tail,
		mp: make(map[int]*Node, capacity),
		cap: capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	if n, ok := this.mp[key]; ok {
		this.remove(n)
		this.insert(n)
		return n.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	n, ok := this.mp[key]
	if ok {
		this.remove(n)
        n.value = value
	}else {
		n = newNode(key, value)
	}
	if this.cap == len(this.mp) {
		this.remove(this.tail.prev)
	}
    this.insert(n)
}

func (this *LRUCache) remove(node *Node) {
	delete(this.mp, node.key)
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (this *LRUCache) insert(node *Node) {
	this.mp[node.key] = node
	next := this.head.next
	this.head.next = node
	node.prev = this.head
	node.next = next
	next.prev = node
}

