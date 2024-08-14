package main

import "fmt"

func main() {
	fmt.Printf("%v\n", bSearchV1([]int{1, 2, 2, 2, 2, 3}, 2))
	fmt.Printf("%v\n", bSearchV2([]int{1, 2, 2, 2, 2, 3}, 2))
	fmt.Printf("%v\n", bSearchV3([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8))

}

// 前一个
func bSearchV1(a []int, target int) int {
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == target {
			if mid == 0 || a[mid-1] != target {
				return mid
			} else {
				end = mid - 1
			}
		} else if a[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1

}

// 后一个
func bSearchV2(a []int, target int) int {
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == target {
			if mid == len(a)-1 || a[mid+1] != target {
				return mid
			} else {
				start = mid + 1
			}

		} else if a[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}
	return -1

}

// 循环有序
func bSearchV3(a []int, target int) int {
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
