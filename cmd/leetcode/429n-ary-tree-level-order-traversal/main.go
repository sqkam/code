package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	ch := make(chan *Node, 2000)
	ch <- root
	result := [][]int{}
	for len(ch) != 0 {
		size := len(ch)
		q := make([]int, 0)
		for _ = range size {
			v := <-ch
			q = append(q, v.Val)
			for _, c := range v.Children {
				ch <- c
			}
		}
		result = append(result, q)
	}
	return result

}
