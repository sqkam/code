package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%v\n", generateMatrix(3))
}

func generateMatrix(n int) [][]int {

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	left, right, top, bottom := 0, n-1, 0, n-1

	idx := 1
	target := n * n
	for idx <= target {
		for i := left; i <= right; i++ {
			matrix[top][i] = idx
			idx++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = idx
			idx++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = idx
			idx++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = idx
			idx++
		}
		left++

	}
	return matrix

}
