package main

import "fmt"

func main() {
	fmt.Printf("%v\n", multiply(3, 4))
}

func multiply(A int, B int) int {
	if B == 0 {
		return 0
	}
	return A + multiply(A, B-1)
}
