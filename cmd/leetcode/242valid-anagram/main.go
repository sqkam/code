package main

func main() {

}
func isAnagram(s string, t string) bool {
	record := [26]int{}
	for _, v := range s {
		record[v-'a']++
	}
	for _, v := range t {
		record[v-'a']--
	}
	return record == [26]int{}
}
