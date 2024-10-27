package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(depth(node.Left), depth(node.Right)) + 1
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
	return (d <= 1) && isBalanced(root.Right) && isBalanced(root.Left)
}
