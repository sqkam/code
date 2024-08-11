package main

import "fmt"

type CircularQueue struct {
	items []int
	head  int
	tail  int
	n     int
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{items: make([]int, n), n: n}
}

func (s *CircularQueue) enqueue(item int) bool {
	if (s.tail+1)%s.n == s.head {
		return false
	}

	s.items[s.tail] = item
	s.tail = (s.tail + 1) % s.n
	return true
}
func (s *CircularQueue) dequeue() int {
	if s.head == s.tail {
		return -1
	}
	val := s.items[s.head]
	s.head = (s.head + 1) % s.n
	return val
}

type Node struct {
	Next *Node
	data int
}
type LinkQueue struct {
	head *Node
	tail *Node
}

func NewLinkQueue() *LinkQueue {
	n := &Node{}
	return &LinkQueue{head: n, tail: n}
}

func (s *LinkQueue) enqueue(item int) bool {
	s.tail.Next = &Node{data: item}
	s.tail = s.tail.Next
	return true
}
func (s *LinkQueue) dequeue() int {
	if s.head.Next == nil {
		return -1
	}
	val := s.head.Next.data
	s.head.Next = s.head.Next.Next
	if s.head.Next == nil {
		s.tail = s.head
	}
	return val

}

func main() {
	as := NewCircularQueue(5)
	as.enqueue(1)
	as.enqueue(2)
	as.enqueue(3)
	as.enqueue(4)
	as.enqueue(5)
	as.enqueue(6)
	as.enqueue(7)
	as.enqueue(8)
	as.dequeue()
	as.dequeue()
	as.dequeue()
	as.dequeue()
	as.dequeue()
	as.dequeue()
	as.dequeue()
	fmt.Printf("%v\n", as)

	ls := NewLinkQueue()
	ls.enqueue(1)
	ls.enqueue(2)
	ls.enqueue(3)
	ls.enqueue(4)
	ls.enqueue(5)
	ls.enqueue(6)
	ls.enqueue(7)
	ls.enqueue(8)
	ls.dequeue()
	ls.dequeue()
	ls.dequeue()
	ls.dequeue()
	ls.dequeue()
	ls.dequeue()
	ls.dequeue()
	ls.dequeue()
	ls.dequeue()
	ls.dequeue()
	ls.dequeue()
	ls.dequeue()
	ls.dequeue()
	ls.enqueue(2)
	ls.enqueue(3)
	fmt.Printf("%v\n", ls)

}
