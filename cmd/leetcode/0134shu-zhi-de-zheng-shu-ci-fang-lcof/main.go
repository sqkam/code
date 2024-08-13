package main

import "fmt"

func main() {
	fmt.Printf("%v\n", myPow(2.10000, 3))
}

var m = make(map[int]float64)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var val float64
	if v, ok := m[n]; ok {
		return v

	}
	if n > 0 {
		val = x * myPow(x, n-1)
	}
	if n < 0 {
		val = 1.0 / x * myPow(x, n+1)
	}
	m[n] = val
	return val

}
