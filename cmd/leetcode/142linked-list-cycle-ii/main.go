package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	fastp := newHead
	slowp := newHead

	for fastp.Next != nil && fastp.Next.Next != nil {
		fastp = fastp.Next.Next
		slowp = slowp.Next
		if fastp.Next == slowp.Next {
			idx1 := fastp.Next
			idx2 := newHead.Next
			for idx1 != idx2 {
				idx1 = idx1.Next
				idx2 = idx2.Next
			}
			return idx2
		}

	}

	return nil
}
