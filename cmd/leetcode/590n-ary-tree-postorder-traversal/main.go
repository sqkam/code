package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	result := []int{}
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}

		for _, v := range root.Children {
			dfs(v)
		}
		result = append(result, root.Val)
	}

	dfs(root)
	return result
}
