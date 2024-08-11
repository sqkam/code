package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	//先来一层头节点
	newHead := &ListNode{Next: head}
	fp := newHead
	sp := newHead
	for fp.Next != nil {
		if fp.Next.Next != nil {
			fp = fp.Next.Next
		} else {
			break
		}
		sp = sp.Next
	}
	return sp.Next

}
