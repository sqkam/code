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
	fmt.Printf("%v\n", waysToStep(999999))
	fmt.Printf("%v\n", time.Since(now).Microseconds())

}

var m = make(map[int]int)

func waysToStep(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	if n == 3 {
		return 4
	}
	val, ok := m[n]
	if !ok {

		val = waysToStep(n-1) + waysToStep(n-2) + waysToStep(n-3)
		val %= 1000000007
		m[n] = val
	}

	return val
}
