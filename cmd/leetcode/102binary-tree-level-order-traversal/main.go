package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	flood int
}

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	levelOrderV4(node)

}

// use size
func levelOrderV4(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ch := make(chan *TreeNode, 1000)
	result := make([][]int, 0)

	ch <- root
	for len(ch) != 0 {
		q := []int{}
		size := len(ch)
		for range size {
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

	return result
}

// 层号
func levelOrderV3(root *TreeNode) {
	if root == nil {
		return
	}
	ch := make(chan *TreeNode, 1000)
	root.flood = 0
	ch <- root
	result := make([][]int, 0)
	for {
		select {
		case v := <-ch:
			if len(result) < v.flood+1 {
				result = append(result, make([]int, 0))
			}
			//fmt.Println(v.Val, v.flood)
			result[v.flood] = append(result[v.flood], v.Val)
			if v.Left != nil {
				v.Left.flood = v.flood + 1
				ch <- v.Left
			}
			if v.Right != nil {
				v.Right.flood = v.flood + 1
				ch <- v.Right
			}
		default:
			fmt.Println(result)
			return
		}

	}

}

// nil来分割层数
func levelOrderV2(root *TreeNode) {
	if root == nil {
		return
	}
	ch := make(chan *TreeNode, 999)
	ch <- root
	ch <- nil
	lastNil := false
	for {
		select {
		case v := <-ch:
			if v == nil {
				if !lastNil {
					ch <- nil
					lastNil = true
				}
				fmt.Println("asdfsadf")
				continue
			}
			lastNil = false
			fmt.Println(v.Val)
			if v.Left != nil {
				ch <- v.Left
			}
			if v.Right != nil {
				ch <- v.Right
			}
		default:
			return
		}

	}
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	temp := []int{root.Val}
	ltemp := dfs(root.Left)
	rtemp := dfs(root.Right)
	return append(temp, append(ltemp, rtemp...)...)
}
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	arr := make([][]int, 0)
	arr = append(arr, []int{root.Val})
	return nil
}
