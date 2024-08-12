package main

import (
	"fmt"
	"strings"
)

func main() {
	calculate("42")
}

func calculate(s string) int {
	s = strings.TrimSpace(s)
	preSign := '+'
	num := 0
	stack := make([]int, 0)
	for i, v := range s {
		isDigit := v >= '0' && v <= '9'
		if isDigit {
			num = num*10 + int(v-'0')
			if i != len(s)-1 {
				continue
			}
		}
		if v == ' ' {
			continue
		}
		switch preSign {
		case '+':
			stack = append(stack, num)
		case '-':
			stack = append(stack, -num)
		case '*':
			stack[len(stack)-1] *= num
		case '/':
			stack[len(stack)-1] /= num
		default:

		}
		preSign = v
		num = 0
	}
	result := 0
	for _, v := range stack {
		result += v
	}
	fmt.Printf("%v\n", result)
	return result
}
