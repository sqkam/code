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
	kk := reverseBookList(l1)
	fmt.Printf("%v\n", kk)
}

func reverseBookList(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	if head.Next == nil {
		return []int{head.Val}
	}
	result := reverseBookList(head.Next)
	return append(result, head.Val)
}
