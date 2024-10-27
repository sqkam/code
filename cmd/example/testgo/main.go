package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	arr := make([]int, 1000)
	for i := range len(arr) {
		arr[i] = rand.IntN(999)
	}
	m := make(map[int]bool)
	for i := range arr {
		_, ok := m[i]
		if ok {
			fmt.Printf("相同的数是%v\n", i)
			break
		}
		m[i] = true
	}
}
