package main

import "fmt"

func main() {
	fmt.Printf("%v\n", myPow(2.10000, 3))
}

var m = make(map[int]float64)

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var val float64
	if v, ok := m[n]; ok {
		return v
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		val = y * y
	} else {
		val = y * y * x
	}
	m[n] = val
	return val

}
func myPowV2(x float64, n int) float64 {
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
