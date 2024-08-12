package main

import (
	"fmt"
)

type Node struct {
	Val   int32
	Count int
}

func main() {
	fmt.Printf("%v\n", removeDuplicates("abbaca", 2))
}
func removeDuplicates(s string, k int) string {

	stack := make([]*Node, 0)
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, &Node{Val: v, Count: 1})
		} else {
			peek := stack[len(stack)-1]
			if peek.Val == v {
				peek.Count++
			} else {
				stack = append(stack, &Node{Val: v, Count: 1})
			}
			if peek.Count >= k {
				stack = stack[:len(stack)-1]
			}
		}
	}

	result := make([]byte, 0)
	for _, v := range stack {
		for v.Count > 0 {
			result = append(result, byte(v.Val))
			v.Count--
		}
	}

	return string(result)
}
