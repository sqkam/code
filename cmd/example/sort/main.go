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
	fmt.Printf("%v\n", BubbleV2(randSlice))
	randomSlice(randSlice)

	//fmt.Printf("%v\n", Insert([]int{18, 73, 67, 33, 50, 29}))
	//fmt.Printf("%v\n", randSlice)
	fmt.Printf("%v\n", "--------------------------insert---------------")
	fmt.Printf("%v\n", Insert(randSlice))
	randomSlice(randSlice)
	fmt.Printf("%v\n", InsertV2(randSlice))
	randomSlice(randSlice)

	//fmt.Printf("%v\n", randSlice)
	//fmt.Printf("%v\n", Select([]int{18, 73, 67, 33, 50, 29}))
	fmt.Printf("%v\n", "--------------------------select---------------")
	fmt.Printf("%v\n", Select(randSlice))
	randomSlice(randSlice)

	fmt.Printf("%v\n", SelectV2(randSlice))
	randomSlice(randSlice)
	//fmt.Printf("%v\n", randSlice)
	fmt.Printf("%v\n", "--------------------------merge---------------")
	MergeSort(randSlice)
	fmt.Printf("%v\n", randSlice)
	randomSlice(randSlice)
	MergeSortV2(randSlice)
	fmt.Printf("%v\n", randSlice)
	randomSlice(randSlice)

	//randSlice = stream.Of(randSlice...).Distinct(func(item int) any {
	//	return item
	//}).ToSlice()
	fmt.Printf("%v\n", "--------------------------quick---------------")
	QuickSort(randSlice)
	fmt.Printf("%v\n", randSlice)
	randomSlice(randSlice)
	QuickSortV2(randSlice)
	fmt.Printf("%v\n", randSlice)
	randomSlice(randSlice)

}
func BubbleV2(s []int) []int {
	for i := range len(s) {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func Bubble(s []int) []int {
	// on2
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

func InsertV2(s []int) []int {
	for i := range len(s) {
		for j := i; j > 0; j-- {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}

	}
	return s
}

func Insert(s []int) []int {
	// on2
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

func SelectV2(s []int) []int {
	for i := range len(s) {
		minPos := i
		for j := i; j < len(s); j++ {
			if s[j] < s[minPos] {
				minPos = j
			}
		}
		s[i], s[minPos] = s[minPos], s[i]

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

func MergeSortV2(s []int) {
	MergeSortRV2(s, 0, len(s)-1)
}
func MergeSortRV2(s []int, start, end int) {
	if start >= end {
		return
	}
	q := (start + end) / 2
	MergeSortRV2(s, start, q)
	MergeSortRV2(s, q+1, end)
	MergeSortMergeV2(s, start, q, end)
}
func MergeSortMergeV2(s []int, p, q, r int) {
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

func MergeSort(s []int) {
	// 传入开始,结束下标
	MergeSortR(s, 0, len(s)-1)
	return
}
func MergeSortR(s []int, p, r int) {
	// 结束条件
	if p >= r {
		return
	}
	q := (p + r) / 2
	// 前半段有序,后半段有序,再合并
	MergeSortR(s, p, q)
	MergeSortR(s, q+1, r)
	MergeSortMerge(s, p, q, r)
}

func MergeSortMerge(s []int, p int, q int, r int) {
	//r 是结尾idx
	// q相对 前一个数组是结尾idx
	//双指针合并两个有序数组
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

func QuickSortV2(s []int) {
	QuickSortRV2(s, 0, len(s)-1)
}
func QuickSortRV2(s []int, p, r int) {
	if p >= r {
		return
	}
	q := partitionV2(s, p, r)
	QuickSortRV2(s, p, q-1)
	QuickSortRV2(s, q+1, r)
}
func partitionV2(s []int, q, r int) int {
	i := q
	for j := q; j < r; j++ {
		if s[j] < s[r] {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[r], s[i] = s[i], s[r]
	return i
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
