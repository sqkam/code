package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3}}
	kk := addTwoNumbers(l1, l2)
	fmt.Printf("%v\n", kk)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newNode := &ListNode{}
	p := newNode
	tail := 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			total := tail + l2.Val
			tail = total / 10
			p.Next = &ListNode{Val: total % 10}
			l2 = l2.Next
		} else if l2 == nil {
			total := tail + l1.Val
			tail = total / 10
			p.Next = &ListNode{Val: total % 10}
			l1 = l1.Next
		} else {
			total := l1.Val + l2.Val + tail
			tail = total / 10
			p.Next = &ListNode{Val: total % 10}
			l1 = l1.Next
			l2 = l2.Next
		}
		p = p.Next
	}
	if tail > 0 {
		p.Next = &ListNode{Val: tail}
	}
	return newNode.Next
}
