package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	maxDepth()
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := maxDepth(root.Left)
	rh := maxDepth(root.Right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}
