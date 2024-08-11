package main

import (
	"fmt"
)

type Node struct {
	Next  *Node
	Value int64
}

var head = &Node{}
var tail = head

func main() {
	insertTail(1)
	insertTail(2)
	insertTail(3)
	insertTail(4)
	//insertHead(99)

	//deleteThisNode(head.Next)
	DeleteByIndex(4)
	fmt.Printf("%v\n", head)

}
func deleteThisNode(n *Node) {
	if n == nil && head == nil {
		return
	}
	//有头节点
	h := head
	for h.Next != nil && h.Next != n {
		h = h.Next
	}
	if h.Next == nil {
		return
	}
	h.Next = h.Next.Next

}

func insertTail(val int64) {
	tail.Next = &Node{Value: val}
	tail = tail.Next
}
func insertHead(val int64) {
	// 因为有头节点
	head.Next = &Node{Value: val, Next: head.Next}
}
func DeleteByIndex(index int64) {
	var i int64
	h := head
	for h.Next != nil && i < index {
		h = h.Next
		i++
	}
	if h.Next != nil {
		h.Next = h.Next.Next
	}

}
