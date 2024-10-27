package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

var count = 0
var val = 0

func inorder(root *TreeNode, k int) {
	if root == nil {
		return
	}
	inorder(root.Left, k)
	if k == count {
		val = root.Val
	}
	count++
	inorder(root.Right, k)
	return
}

// todo 失败
func kthSmallest(root *TreeNode, k int) int {
	inorder(root, k)
	return val
}
