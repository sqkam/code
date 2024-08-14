package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", search(nums, 3))
}

func search(nums []int, target int) int {
	return bSearch(nums, 0, len(nums)-1, target)
}
func bSearch(nums []int, start, end, target int) int {
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
