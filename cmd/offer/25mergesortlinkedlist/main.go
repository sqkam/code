package main

import (
	"fmt"
)

type ListNode struct {
	Next  *ListNode
	Value int64
}

func main() {
	l1 := &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 4}}}
	l2 := &ListNode{Value: 1, Next: &ListNode{Value: 3}}
	kk := mergeV0(l1, l2)
	fmt.Printf("%v\n", kk)
}

func mergeV0X(l1, l2 *ListNode) *ListNode {
	//套路版本
	newL1 := &ListNode{Next: l1}
	newL2 := &ListNode{Next: l2}
	newList := &ListNode{}
	p := newList // 跟踪的
	for newL1.Next != nil || newL2.Next != nil {
		if newL1.Next == nil {
			p.Next = &ListNode{Value: newL2.Next.Value}
			newL2.Next = newL2.Next.Next
		} else if newL2.Next == nil {
			p.Next = &ListNode{Value: newL1.Next.Value}
			newL1.Next = newL1.Next.Next
		} else if newL1.Next.Value < newL2.Next.Value {
			p.Next = &ListNode{Value: newL1.Next.Value}
			newL1.Next = newL1.Next.Next
		} else {
			p.Next = &ListNode{Value: newL2.Next.Value}
			newL2.Next = newL2.Next.Next
		}
		p = p.Next

	}
	//一定的
	return newList.Next
}
func mergeV0(l1, l2 *ListNode) *ListNode {
	newListNode := &ListNode{}
	p := newListNode // 跟踪的
	for l1 != nil || l2 != nil {
		if l1 == nil {
			p.Next = &ListNode{Value: l2.Value}
			l2 = l2.Next
		} else if l2 == nil {
			p.Next = &ListNode{Value: l1.Value}
			l1 = l1.Next
		} else if l1.Value < l2.Value {
			p.Next = &ListNode{Value: l1.Value}
			l1 = l1.Next
		} else {
			p.Next = &ListNode{Value: l2.Value}
			l2 = l2.Next
		}
		p = p.Next

	}
	//一定的
	return newListNode.Next
}
func mergeV1(l1, l2 *ListNode) *ListNode {
	newListNode := &ListNode{}
	p := newListNode // 跟踪的
	for l1 != nil || l2 != nil {
		if l1 == nil {
			p.Next = l2
			break
		} else if l2 == nil {
			p.Next = l1
			break
		} else if l1.Value < l2.Value {
			p.Next = &ListNode{Value: l1.Value}
			l1 = l1.Next
		} else {
			p.Next = &ListNode{Value: l2.Value}
			l2 = l2.Next
		}
		p = p.Next
	}
	//一定的
	return newListNode.Next
}
func mergeV2(l1, l2 *ListNode) *ListNode {
	newListNode := &ListNode{}
	p := newListNode // 跟踪的
	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			p.Next = &ListNode{Value: l1.Value}
			l1 = l1.Next
		} else {
			p.Next = &ListNode{Value: l2.Value}
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	} else if l2 != nil {
		p.Next = l2
	}
	//一定的
	return newListNode.Next
}
