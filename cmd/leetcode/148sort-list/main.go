package main

import (
	"fmt"
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	kk := sortList(l1)
	fmt.Printf("%v\n", kk)
}

func sortList(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	newList := &ListNode{}
	var b []int
	p := newHead
	for p.Next != nil {
		b = append(b, p.Next.Val)
		p = p.Next
	}

	//mergeSort(b, 0, len(b)-1)
	slices.SortFunc(b, func(a, b int) int {
		if (a - b) > 0 {
			return -1
		}
		return 0
	})
	for _, v := range b {
		newList.Next = &ListNode{Next: newList.Next, Val: v}
	}

	return newList.Next
}

func mergeSort(b []int, p int, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	mergeSort(b, p, q)
	mergeSort(b, q+1, r)
	merge(b, p, q, r)

}

func merge(b []int, p int, q int, r int) {
	i := p
	j := q + 1
	tmp := make([]int, 0)
	for i <= q && j <= r {
		if b[i] > b[j] {
			tmp = append(tmp, b[i])
			i++
		} else {
			tmp = append(tmp, b[j])
			j++
		}
	}
	for i <= q {
		tmp = append(tmp, b[i])
		i++
	}
	for j <= r {
		tmp = append(tmp, b[j])
		j++
	}
	for idx, v := range tmp {
		b[p+idx] = v
	}
}

func sortListV2(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	newList := &ListNode{}

	for newHead.Next != nil {
		p := newHead
		maxNext := newHead
		for p.Next != nil {
			if p.Next.Val > maxNext.Next.Val {
				maxNext = p
			}
			p = p.Next
		}

		newList.Next = &ListNode{Next: newList.Next, Val: maxNext.Next.Val}

		//删除
		maxNext.Next = maxNext.Next.Next
	}
	return newList.Next
}
