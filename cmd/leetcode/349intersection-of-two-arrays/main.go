package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]struct{}{}
	for _, v := range nums1 {
		m[v] = struct{}{}
	}
	resultm := map[int]struct{}{}
	for _, v := range nums2 {
		_, ok := m[v]
		if ok {
			resultm[v] = struct{}{}
		}
	}
	result := make([]int, 0, len(resultm))
	for k := range resultm {
		result = append(result, k)
	}
	return result
}
