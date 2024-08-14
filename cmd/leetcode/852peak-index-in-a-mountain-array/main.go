package main

import "fmt"

func main() {
	res := peakIndexInMountainArray([]int{1, 2, 3, 0})
	fmt.Printf("%v\n", res)
}

func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		if mid == 0 {
			start = mid + 1
		} else if mid+1 == len(arr) {
			end = mid - 1

		} else if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid-1] < arr[mid] {
			start = mid + 1

		} else {
			end = mid - 1

		}

	}
	return -1
}
