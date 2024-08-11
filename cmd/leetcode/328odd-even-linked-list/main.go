package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	kk := oddEvenList(l1)
	fmt.Printf("%v\n", kk)
}
func oddEvenList(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	p := newHead
	newList := &ListNode{}
	np := newList
	evenList := &ListNode{}
	ep := evenList
	count := 1
	for p.Next != nil {
		if count%2 == 0 {
			ep.Next = &ListNode{Val: p.Next.Val}
			ep = ep.Next
		} else {
			np.Next = &ListNode{Val: p.Next.Val}
			np = np.Next
		}
		count++
		p = p.Next
	}
	np.Next = evenList.Next
	return newList.Next
}

func oddEvenListV2(head *ListNode) *ListNode {

	p := head
	newList := &ListNode{}
	np := newList
	evenList := &ListNode{}
	ep := evenList
	count := 1
	for p != nil {
		if count%2 == 0 {
			ep.Next = &ListNode{Val: p.Val}
			ep = ep.Next
		} else {
			np.Next = &ListNode{Val: p.Val}
			np = np.Next
		}
		count++
		p = p.Next
	}
	np.Next = evenList.Next
	return newList.Next
}
