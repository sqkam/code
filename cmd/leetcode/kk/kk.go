package main

import "fmt"

func main() {
	fmt.Printf("%v\n", solution("aba"))
}
func solution(s string) bool {
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}
