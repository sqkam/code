package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}

	kk := deleteDuplicates(l1)
	fmt.Printf("%v\n", kk)
}

func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	p := newHead
	newList := &ListNode{}
	np := newList

	m := make(map[int]int)
	for p.Next != nil {
		v, _ := m[p.Next.Val]
		m[p.Next.Val] = v + 1
		p = p.Next
	}
	p = newHead
	for p.Next != nil {
		v, _ := m[p.Next.Val]
		if v == 1 {
			np.Next = &ListNode{Val: p.Next.Val}
			np = np.Next
		}
		p = p.Next
	}
	return newList.Next
}
