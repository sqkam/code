package main

import "fmt"

func main() {
	fmt.Printf("%v\n", findMin([]int{3, 4, 5, 1, 2}))
}

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if start == end {
			return nums[mid]
		}
		if nums[(len(nums)-1+mid)%len(nums)] > nums[mid] {
			return nums[mid]
		} else if nums[mid] > nums[end] {
			start = mid + 1

		} else {
			end = mid - 1
		}

	}
	return -1

}
