package main

import "fmt"

func main() {
	total := 0
	fmt.Scan(&total)
	arr := make([]int, 0, total)
	arrSum := make([]int, 0, total)
	for i := range total {
		fmt.Scan(&arr[i])
	}

	for i := range arrSum {
		if i > 0 {
			arrSum[i] = arrSum[i-1] + arr[i]
		} else {
			arrSum[i] = arr[i]
		}
	}
	for {
		a, b := 0, 0
		fmt.Scan(&a, &b)
		if a==0 {
			fmt.Printf("%v\n", arrSum[b])
		}else {
			fmt.Printf("%v\n", arrSum[b]-arrSum[a-1])
		}

	}

}