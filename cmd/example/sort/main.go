package main

import (
	"fmt"
	"math/rand/v2"
)

func randomSlice(s []int) {
	for i := range len(s) / 2 {
		r := rand.IntN(len(s))
		s[r], s[i] = s[i], s[r]
	}
}
func main() {
	var randSlice []int
	for i := range 48 {
		_ = i
		randSlice = append(randSlice, rand.IntN(99))
	}
	//fmt.Printf("%v\n", Bubble([]int{18, 73, 67, 33, 50, 29}))
	//fmt.Printf("%v\n", randSlice)
	fmt.Printf("%v\n", Bubble(randSlice))
	randomSlice(randSlice)

	//fmt.Printf("%v\n", Insert([]int{18, 73, 67, 33, 50, 29}))
	//fmt.Printf("%v\n", randSlice)
	fmt.Printf("%v\n", Insert(randSlice))
	randomSlice(randSlice)

	//fmt.Printf("%v\n", randSlice)
	//fmt.Printf("%v\n", Select([]int{18, 73, 67, 33, 50, 29}))
	fmt.Printf("%v\n", Select(randSlice))
	randomSlice(randSlice)

	//fmt.Printf("%v\n", randSlice)
	MergeSort(randSlice)
	fmt.Printf("%v\n", randSlice)
	randomSlice(randSlice)

	//randSlice = stream.Of(randSlice...).Distinct(func(item int) any {
	//	return item
	//}).ToSlice()
	QuickSort(randSlice)
	fmt.Printf("%v\n", randSlice)
	randomSlice(randSlice)

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

func Select(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		minPos := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[minPos] {
				minPos = j
			}
		}
		s[minPos], s[i] = s[i], s[minPos]
	}
	return s
}

func MergeSort(s []int) {
	MergeSortR(s, 0, len(s)-1)
	return
}
func MergeSortR(s []int, p, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	MergeSortR(s, p, q)
	MergeSortR(s, q+1, r)
	MergeSortMerge(s, p, q, r)
}

func MergeSortMerge(s []int, p int, q int, r int) {
	i := p
	j := q + 1
	tmp := make([]int, 0)
	for i <= q && j <= r {
		if s[i] < s[j] {
			tmp = append(tmp, s[i])
			i++
		} else {
			tmp = append(tmp, s[j])
			j++
		}
	}
	for i <= q {
		tmp = append(tmp, s[i])
		i++
	}
	for j <= r {
		tmp = append(tmp, s[j])
		j++
	}
	for idx, v := range tmp {
		s[p+idx] = v
	}
}

func QuickSort(s []int) {
	QuickSortR(s, 0, len(s)-1)
}
func QuickSortR(s []int, p, r int) {
	if p >= r {
		return
	}

	q := partition(s, p, r)
	QuickSortR(s, p, q-1)
	QuickSortR(s, q+1, r)

}

func partition(s []int, p int, r int) int {
	i := p
	for j := p; j < r; j++ {
		if s[j] < s[r] {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[r] = s[r], s[i]
	return i
}
