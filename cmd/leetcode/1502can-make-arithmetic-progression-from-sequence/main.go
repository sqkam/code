package main

import (
	"fmt"
)

func main() {
	slice := []int{-68, -96, -12, -40, 16}
	fmt.Printf("%v\n", canMakeArithmeticProgression(slice))
}

func canMakeArithmeticProgression(arr []int) bool {
	QuickSort(arr, 0, len(arr)-1)
	d := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		if (arr[i] - arr[i-1]) != d {
			return false
		}
	}
	return true
}
func QuickSort(arr []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(arr, p, r)
	QuickSort(arr, p, q-1)
	QuickSort(arr, q+1, r)
}

func partition(arr []int, p, r int) int {
	i := p
	for j := p; j < r; j++ {
		if arr[j] < arr[r] {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]

	return i
}
