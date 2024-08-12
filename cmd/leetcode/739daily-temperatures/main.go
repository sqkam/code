package main

import (
	"fmt"
)

func main() {
	s := []int{99, 99, 100}
	dailyTemperaturesV3(s)
}
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, len(temperatures))
	for i, v := range temperatures {
		for si := i + 1; si < n; si++ {
			if v < temperatures[si] {
				result[i] = si - i
				break
			}
		}
	}
	fmt.Printf("%v\n", result)
	return result

}

func dailyTemperaturesV2(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i, v := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		if v > temperatures[stack[len(stack)-1]] {
			for si := len(stack) - 1; si > -1; si-- {
				if temperatures[stack[si]] < v {
					result[stack[si]] = i - stack[si]
					stack = stack[:len(stack)-1]
				}
			}
		}
		stack = append(stack, i)
	}
	//fmt.Printf("%v\n", result)
	return result

}

func dailyTemperaturesV3(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i, v := range temperatures {

		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	//fmt.Printf("%v\n", result)
	return result

}
