package main

func main() {

}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0

	m := map[int]int{}
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}

	for _, v1 := range nums3 {
		for _, v2 := range nums4 {
			hitCount, ok := m[-v1-v2]
			if ok {
				count += hitCount
			}
		}
	}
	return count
}
