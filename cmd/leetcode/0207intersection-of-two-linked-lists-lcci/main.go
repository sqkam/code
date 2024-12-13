package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	newHeadA := &ListNode{Next: headA}
	newHeadB := &ListNode{Next: headB}
	h1p := newHeadA
	h2p := newHeadB
	len1 := 0
	len2 := 0

	for h1p.Next != nil {
		len1++
		h1p = h1p.Next
	}

	for h2p.Next != nil {
		len2++
		h2p = h2p.Next
	}
	h1p = newHeadA
	h2p = newHeadB
	if len2 > len1 {
		len1, len2 = len2, len1
		h1p, h2p = h2p, h1p
	}

	for range len1 - len2 {
		h1p = h1p.Next
	}

	for h1p.Next != nil {
		if h1p.Next.Val == h2p.Next.Val {
			return h1p.Next
		}
		h1p = h1p.Next
		h2p = h2p.Next
	}

	return nil
}
