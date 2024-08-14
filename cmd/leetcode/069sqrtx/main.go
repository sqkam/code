package main

import "fmt"

func main() {

	fmt.Printf("%v\n", mySqrt(9))
}

func mySqrt(x int) int {
	start := 0
	end := x
	ans := x
	for start <= end {
		mid := (start + end) / 2
		if mid*mid <= x {
			ans = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return ans
}
