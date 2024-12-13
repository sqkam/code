package main

func main() {

}

func isHappy(n int) bool {
	m := map[int]struct{}{}
	for {
		sum := getSum(n)
		if sum == 1 {
			return true
		}
		_, ok := m[sum]
		if ok {
			return false
		}
		m[sum] = struct{}{}
		n = sum
	}
}
func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
