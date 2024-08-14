package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v\n", reverse(-123))

}

func reverse(x int) int {
	flag := 1
	if x <= 0 {
		flag = -1
		x = -x
	}
	tmp := 0
	for x > 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if math.MinInt32 < flag*tmp && flag*tmp < math.MaxInt32 {
		return flag * tmp
	}
	return 0

}
