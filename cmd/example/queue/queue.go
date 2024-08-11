package main

import "fmt"

type ArrayStack struct {
	items []int
	count int
	n     int
}

func NewArrayStack(n int) *ArrayStack {
	return &ArrayStack{items: make([]int, n), n: n}
}

func (s *ArrayStack) push(item int) bool {
	if s.count == s.n {
		return false
	}
	s.items[s.count] = item
	s.count++
	return true
}
func (s *ArrayStack) pop() int {
	if s.count == 0 {
		return -1
	}
	val := s.items[s.count-1]

	s.count--
	return val
}

func (s *ArrayStack) peek() int {
	if s.count == 0 {
		return -1
	}
	val := s.items[s.count-1]
	return val
}

type Node struct {
	Next *Node
	data int
}
type LinkStack struct {
	head *Node
}

func NewLinkStack() *LinkStack {
	return &LinkStack{head: &Node{}}
}

func (s *LinkStack) push(item int) bool {
	s.head.Next = &Node{Next: s.head.Next, data: item}
	return true
}
func (s *LinkStack) pop() int {
	if s.head.Next == nil {
		return -1
	}
	val := s.head.Next.data
	s.head.Next = s.head.Next.Next
	return val

}

func (s *LinkStack) peek() int {
	if s.head.Next == nil {
		return -1
	}
	val := s.head.Next.data
	return val
}

func main() {
	as := NewArrayStack(5)
	as.push(1)
	as.push(2)
	as.push(3)
	as.push(4)
	as.push(5)
	as.push(6)
	as.push(7)
	as.push(8)
	as.pop()
	as.pop()
	as.pop()
	as.pop()
	as.pop()
	as.pop()
	as.pop()
	fmt.Printf("%v\n", as)

	ls := NewLinkStack()
	ls.push(1)
	ls.push(2)
	ls.push(3)
	ls.push(4)
	ls.push(5)
	ls.push(6)
	ls.push(7)
	ls.push(8)
	ls.pop()
	ls.pop()
	ls.pop()
	ls.pop()
	ls.pop()
	ls.pop()
	ls.pop()
	fmt.Printf("%v\n", ls)

}
