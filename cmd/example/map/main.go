package main

import "fmt"

func main() {
	m := make(map[string]int)
	val := m["adsfsadf"]
	//不存在就是0
	fmt.Printf("%v\n", val)
}
