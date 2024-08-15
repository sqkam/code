package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Printf("%v\n", threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	resultMap := make(map[string]bool)
	result := make([][]int, 0)
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			idx := m[-nums[i]-nums[j]]
			if idx > j {
				_, ok := resultMap[fmt.Sprintf("%d-%d-%d", nums[i], nums[j], nums[idx])]
				if !ok {
					resultMap[fmt.Sprintf("%d-%d-%d", nums[i], nums[j], nums[idx])] = true
					result = append(result, []int{nums[i], nums[j], nums[idx]})
				}
			}
		}
	}

	return result
}
