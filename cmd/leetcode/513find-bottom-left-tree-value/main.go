package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func findBottomLeftValue(root *TreeNode) int {
	ch := make(chan *TreeNode, 2000)
	ch <- root
	lastElem := root.Val
	for len(ch) > 0 {
		v := <-ch
		lastElem = v.Val
		if v.Right != nil {
			ch <- v.Right
		}
		if v.Left != nil {
			ch <- v.Left
		}
	}
	return lastElem

}
