package main

import "fmt"

func main() {

	fmt.Printf("%v\n", string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd')))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	start := 0
	end := len(letters) - 1
	for start <= end {
		mid := (start + end) / 2
		if letters[mid] == target {
			return letters[mid+1]
		} else if letters[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return letters[0]
}
