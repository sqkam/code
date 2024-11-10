package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next: head}
	fastp := newHead
	slowp := newHead
	for range n {
		fastp = fastp.Next
	}
	for fastp.Next != nil {
		fastp = fastp.Next
		slowp = slowp.Next
	}
	slowp.Next = slowp.Next.Next
	return newHead.Next
}
