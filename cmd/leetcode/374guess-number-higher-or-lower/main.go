package main

func guess(num int) int {
	return -1
}
func main() {
	guessNumber(1)
}
func guessNumber(n int) int {
	start := 0

	for {
		mid := (start + n) / 2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == -1 {
			n = mid - 1
		} else {
			start = mid + 1
		}

	}

}
