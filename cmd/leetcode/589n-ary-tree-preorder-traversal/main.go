package main

import (
	"fmt"
	"slices"
)

type Node struct {
	Val      int
	Children []*Node
}

// 这种才行,v2不行不知道为什么,就是先定义dfs,定义出口,狠狠造
func preorder(root *Node) []int {
	var result []int
	var dfs = func(root *Node) {}
	dfs = func(root *Node) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		for _, v := range root.Children {
			dfs(v)
		}
	}
	dfs(root)
	return result
}

var resultV2 []int

func preorderV2(root *Node) []int {
	if root == nil {
		return nil
	}
	resultV2 = append(resultV2, root.Val)
	for _, v := range root.Children {
		preorderV2(v)
	}

	return resultV2
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
	//preorder(root)
	fmt.Printf("%v\n", preorder(root))
	preorderV2(root)
	fmt.Printf("%v\n", resultV2)
	fmt.Printf("%v\n", slices.Equal(resultV2, preorder(root)))
}
