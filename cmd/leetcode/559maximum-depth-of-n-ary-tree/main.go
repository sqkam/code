package main

import (
	"fmt"
	"slices"
)

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	// 就是记住这两个条件就行了
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}
	var depths = make([]int, 0)
	for _, v := range root.Children {
		depths = append(depths, maxDepth(v))
	}

	return slices.Max(depths) + 1
}
func main() {

	root := &Node{Val: 1, Children: []*Node{
		{Val: 2},
		{Val: 3, Children: []*Node{
			{Val: 6}, {Val: 7, Children: []*Node{{Val: 11, Children: []*Node{{Val: 14}}}}},
		}},
		{Val: 4, Children: []*Node{{Val: 8, Children: []*Node{{Val: 12}}}}},
		{Val: 5, Children: []*Node{{Val: 9, Children: []*Node{{Val: 13}}}, {Val: 10}}},
	}}
	ret := maxDepth(root)
	fmt.Printf("%v\n", ret)
}
