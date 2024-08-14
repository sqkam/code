package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1}}}
	kk := insertionSortList(l1)
	fmt.Printf("%v\n", kk)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	newList := &ListNode{}

	for newHead.Next != nil {
		node := &ListNode{Val: newHead.Next.Val}
		//del
		newHead.Next = newHead.Next.Next
		p := newList
		for p.Next != nil && p.Next.Val < node.Val {
			p = p.Next
		}
		node.Next = p.Next
		p.Next = node
	}
	return newList.Next
}
