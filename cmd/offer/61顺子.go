package main

import "fmt"

func main() {
	//kk := isStraight([]int{1, 2, 3, 4, 5})
	//kk := isStraight([]int{1, 2, 3, 4, 5})
	//kk := isStraight([]int{0, 0, 3, 4, 5})
	kk := isStraight([]int{2, 0, 3, 4, 5})
	fmt.Printf("%v", kk)
}

func isStraight(nums []int) bool {
	dup := make([]bool, 14)
	min := 100
	max := -1
	for _, v := range nums {
		if v == 0 {
			continue
		}
		if dup[v] {
			return false
		}
		dup[v] = true
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}

	}
	return (max - min) < 5

}
