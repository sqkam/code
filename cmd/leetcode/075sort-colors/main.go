package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	sortColors(slice)
	fmt.Printf("%v\n", slice)
}

func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, p int, r int) {
	if p >= r {
		return
	}
	q := partition(nums, p, r)
	quickSort(nums, p, q-1)
	quickSort(nums, q+1, r)
}

func partition(nums []int, p int, r int) int {
	i := p
	for j := p; j < r; j++ {
		if nums[j] < nums[r] {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[r], nums[i] = nums[i], nums[r]
	return i
}
