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
	kk = addTwoNumbersV2(l1, l2)
	fmt.Printf("12341234%v\n", kk)
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	newNode := &ListNode{}
	p := newNode
	newL1 := &ListNode{Next: l1}
	newL2 := &ListNode{Next: l2}
	carry := 0
	for newL1.Next != nil && newL2.Next != nil {
		sum := newL1.Next.Val + newL2.Next.Val + carry
		carry = sum / 10

		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
		newL1 = newL1.Next
		newL2 = newL2.Next
	}
	for newL1.Next != nil {
		sum := newL1.Next.Val + carry
		carry = sum / 10
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
		newL1 = newL1.Next
	}
	for newL2.Next != nil {
		sum := newL2.Next.Val + carry
		carry = sum / 10
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
		newL2 = newL2.Next
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
		p = p.Next
	}
	return newNode.Next

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
