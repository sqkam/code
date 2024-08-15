package main

func main() {

	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	lRUCache.Get(1)    // 返回 1
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Get(2)    // 返回 -1 (未找到)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.Get(1)    // 返回 -1 (未找到)
	lRUCache.Get(3)    // 返回 3
	lRUCache.Get(4)    // 返回 4
}

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}
type LRUCache struct {
	Capacity int
	Size     int
	Head     *Node
	Tail     *Node
	m        map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{Capacity: capacity, Size: 0, Head: head, Tail: tail, m: make(map[int]*Node)}
}

func (l *LRUCache) addHead(key int, val int) *Node {
	newNode := l.Head.Next
	node := &Node{Key: key, Val: val, Next: l.Head.Next, Prev: l.Head}
	newNode.Prev = node
	l.Head.Next = node
	l.Size++
	l.m[key] = node
	return node
}
func (l *LRUCache) removeNode(node *Node) {
	if l.Head.Next == l.Tail {
		return
	}
	preNode := node.Prev
	nextNode := node.Next
	preNode.Next = nextNode
	nextNode.Prev = preNode
	delete(l.m, node.Key)
	l.Size--

}

func (l *LRUCache) Get(key int) int {
	v, ok := l.m[key]
	if !ok {
		return -1
	}
	l.removeNode(v)
	l.addHead(v.Key, v.Val)
	return v.Val
}

func (l *LRUCache) Put(key int, value int) {
	v, ok := l.m[key]
	if ok {
		v.Val = value
		l.removeNode(v)
		l.addHead(v.Key, v.Val)
		return
	}
	if l.Size == l.Capacity {
		l.removeNode(l.Tail.Prev)
	}
	l.addHead(key, value)
}
