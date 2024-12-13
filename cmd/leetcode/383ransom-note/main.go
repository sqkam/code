package main

func canConstruct(ransomNote string, magazine string) bool {
	m := [26]int{}
	for _, v := range magazine {
		m[v-'a']++
	}
	for _, v := range ransomNote {
		if m[v-'a'] > 0 {
			m[v-'a']--
		} else {
			return false
		}

	}
	return true
}
