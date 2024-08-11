package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	kk := myAtoi("4193 with words")
	fmt.Printf("%v", kk)
}

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	var flag int64 = 0
	var sum int64

	for i, v := range str {
		if i == 0 && v == '-' {
			flag = -1
		} else if i == 0 && v == '+' {

		} else if v > '9' || v < '0' {
			break
		} else {

			temp := sum*10 + int64(v-'0')
			if temp >= math.MaxInt32-flag {
				sum = math.MaxInt32 - flag
				break
			}
			sum = temp

		}
	}
	if flag == -1 {
		return int(-sum)
	}
	return int(sum)
}
