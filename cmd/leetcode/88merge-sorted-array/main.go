package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func merge(A []int, m int, B []int, n int) {
	for i, v := range B {
		A[m+i] = v
		for j := m + i; j > 0; j-- {
			if A[j-1] > A[j] {
				A[j-1], A[j] = A[j], A[j-1]
			} else {
				break
			}
		}
	}
}
