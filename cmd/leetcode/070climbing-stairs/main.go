package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	now := time.Now()
	fmt.Printf("%v\n", climbStairsV2(999999))
	fmt.Printf("%v\n", time.Since(now).Microseconds())

}

var m = make(map[int]int)

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	val1, ok := m[n-1]
	if !ok {
		val1 = climbStairs(n - 1)
		m[n-1] = val1
	}
	val2, ok := m[n-2]
	if !ok {
		val2 = climbStairs(n - 2)
		m[n-2] = val2
	}
	return val2 + val1
}
func climbStairsV2(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	val, ok := m[n]
	if !ok {

		val = climbStairsV2(n-1) + climbStairsV2(n-2)
		m[n] = val
	}

	return val
}

func climbStairsV3(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	val1 := climbStairsV3(n - 1)

	val2 := climbStairsV3(n - 2)

	return val2 + val1
}
