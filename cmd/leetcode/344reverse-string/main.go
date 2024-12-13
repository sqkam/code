package main

func main() {

}
func reverseString(s []byte) {
	start := 0
	end := len(s) - 1
	for range len(s) / 2 {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
