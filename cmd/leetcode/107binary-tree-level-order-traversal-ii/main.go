package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ch := make(chan *TreeNode, 2000)
	ch <- root
	result := make([][]int, 0)
	for len(ch) > 0 {
		size := len(ch)
		q := make([]int, 0)
		for _ = range size {
			v := <-ch
			q = append(q, v.Val)
			if v.Left != nil {
				ch <- v.Left
			}
			if v.Right != nil {
				ch <- v.Right
			}
		}
		result = append(result, q)

	}
	slices.Reverse(result)
	return result

}
