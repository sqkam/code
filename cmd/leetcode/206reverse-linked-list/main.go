package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	kk := reverseList(l1)
	fmt.Printf("%v\n", kk)
}

func reverseList(head *ListNode) *ListNode {
	newList := &ListNode{}
	newHead := &ListNode{
		Next: head,
	}
	p := newHead
	for p.Next != nil {

		newList.Next = &ListNode{Val: p.Next.Val, Next: newList.Next}
		p = p.Next
	}
	return newList.Next
}
