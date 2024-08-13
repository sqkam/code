package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var randSlice []int
	for i := range 48 {
		_ = i
		randSlice = append(randSlice, rand.IntN(99))

	}
	fmt.Printf("%v\n", Bubble(randSlice))
	fmt.Printf("%v\n", Bubble([]int{18, 73, 67, 33, 50, 29}))
	fmt.Printf("%v\n", Insert(randSlice))

	fmt.Printf("%v\n", Insert([]int{18, 73, 67, 33, 50, 29}))

}
func Bubble(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		sorted := false
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				sorted = true
			}
		}
		if !sorted {
			break
		}

	}
	return s
}

func Insert(s []int) []int {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			} else {
				break
			}
		}
	}
	return s
}
