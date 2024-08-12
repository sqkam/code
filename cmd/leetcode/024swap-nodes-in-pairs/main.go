package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}

	kk := swapPairs(l1)
	fmt.Printf("%v\n", kk)
}

func swapPairs(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	p := newHead
	newList := &ListNode{}
	np := newList
	for p.Next != nil && p.Next.Next != nil {
		np.Next = &ListNode{Val: p.Next.Next.Val}
		np.Next.Next = &ListNode{Val: p.Next.Val}
		p = p.Next.Next
		np = np.Next.Next
	}
	if p.Next != nil {
		np.Next = &ListNode{Val: p.Next.Val}

	}
	return newList.Next
}
