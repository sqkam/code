package main

import "fmt"

// 给定一个整数数组，请实现一个函数找出数组中出现次数最多的元素。如果有多个元素出现的次数相同，返回任意一个即可。
func main() {
	arr := []int{
		1, 1, 1,
		2, 2, 2,
		3, 3, 3,
		4, 5,
	}
	solution(arr)

}

func solution(arr []int) {

	m := make(map[int]int)
	if len(arr) == 0 {
		fmt.Printf("数组为空\n")
		return
	}

	maxCountKey := arr[0]
	maxCount := 0

	for _, v := range arr {
		val, ok := m[v]
		if ok {
			m[v] = val + 1

		} else {
			m[v] = 1
		}

		if val >= maxCount {
			maxCount = val
			maxCountKey = v
		}
	}

	fmt.Printf("最大的key %v\n", maxCountKey)
}
