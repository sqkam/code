package main

import "fmt"

func main() {
	fmt.Printf("%v\n", search([]int{4, 5, 6, 7, 0, 1, 2}, 7))
}
func search(a []int, target int) int {
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == target {
			return mid
		} else if a[start] <= a[mid] {
			if target >= a[start] && target < a[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target <= a[end] && target > a[mid] {

				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
