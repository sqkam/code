package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func lengthOfLongestSubstring(s string) int {
	maxCount := 0
	sBytes := []byte(s)
	for i, _ := range sBytes {
		m := map[byte]string{}
		startIndex := i
		for startIndex < len(sBytes) {
			if _, ok := m[sBytes[startIndex]]; ok {
				break
			}
			m[sBytes[startIndex]] = ""
			startIndex++
		}
		lenString := startIndex - i
		//lenString := len(m)
		if maxCount < lenString {
			maxCount = lenString
		}
	}
	return maxCount

}
