package main

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	//1
	//	/ \
	//	2   3
	//	/ \ / \
	//	4 5 6   7

	node := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}
	fmt.Printf("%v\n", "pre")
	preOrder(node)
	fmt.Printf("%v\n", "mid")
	midOrder(node)
	fmt.Printf("%v\n", "after")
	afterOrder(node)
	//fmt.Printf("%v\n", node)
}

func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%v\n", node.Val)
	preOrder(node.Left)

	preOrder(node.Right)
}

func midOrder(node *Node) {
	if node == nil {
		return
	}

	midOrder(node.Left)
	fmt.Printf("%v\n", node.Val)
	midOrder(node.Right)
}

func afterOrder(node *Node) {
	if node == nil {
		return
	}
	afterOrder(node.Left)
	afterOrder(node.Right)
	fmt.Printf("%v\n", node.Val)
}

func findNodeR(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	if node.Val == data {
		return node
	} else if data < node.Val {
		return findNodeR(node.Left, data)
	} else {
		return findNodeR(node.Right, data)
	}
}

func findNode(node *Node, data int) *Node {
	p := node
	for p != nil {
		if p.Val == data {
			return p
		} else if data < p.Val {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return nil
}

// 二叉树 插入
func insertR(node *Node, data int) {
	if node == nil {
		return
	}
	if data < node.Val {
		if node.Left == nil {
			node.Left = &Node{Val: data}
		} else {
			insertR(node.Left, data)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Val: data}
		} else {
			insertR(node.Right, data)
		}
	}
}

func insert(node *Node, data int) {
	if node == nil {
		return
	}
	p := node
	for p != nil {
		if data < p.Val {
			if p.Left == nil {
				p.Left = &Node{Val: data}
				//exit
				return
			} else {
				p = p.Left
			}
		} else {
			if p.Right == nil {
				p.Right = &Node{Val: data}
				return
			} else {
				p = p.Right
			}
		}
	}
}
