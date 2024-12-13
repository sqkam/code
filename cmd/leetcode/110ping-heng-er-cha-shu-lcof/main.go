package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	fmt.Printf("%v\n", isBalancedV2(nil))
}

var isBalance = true

func isBalancedV2(root *TreeNode) bool {
	height(root)
	return isBalance
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH := height(root.Left)
	rightH := height(root.Right)
	dH := abs(leftH - rightH)
	if dH > 1 {
		isBalance = false
	}
	return max(leftH, rightH) + 1
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	d := abs(depth(root.Left) - depth(root.Right))

	return d <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
