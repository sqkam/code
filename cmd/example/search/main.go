package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("%v\n", BSearchV1(s, 0, len(s)-1, 5))
	fmt.Printf("%v\n", BSearchV2(s, 0, len(s)-1, 6))

}
func BSearchV1(a []int, start, end, target int) int {
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == target {
			return mid
		} else if target < a[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func BSearchV2(a []int, start, end, target int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if a[mid] == target {
		return mid
	} else if target < a[mid] {
		return BSearchV2(a, start, mid-1, target)
	} else {
		return BSearchV2(a, mid+1, end, target)
	}

}
