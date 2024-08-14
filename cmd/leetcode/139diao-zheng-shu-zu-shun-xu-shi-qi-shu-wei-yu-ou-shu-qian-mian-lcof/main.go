package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", trainingPlan([]int{1, 2, 3, 4, 5}))
}

func trainingPlan(actions []int) []int {
	i := 0
	j := len(actions) - 1
	for i < j {
		if actions[i]%2 == 1 {
			i++
			continue
		}
		if actions[j]%2 == 0 {
			j--
			continue
		}
		actions[i], actions[j] = actions[j], actions[i]
		//可以不用下面两个
		i++
		j--
	}

	return actions
}
func trainingPlanV2(actions []int) []int {
	sort(actions, 0, len(actions)-1)
	return actions
}

func sort(actions []int, p int, r int) {
	if p >= r {
		return
	}
	q := partition(actions, p, r)
	sort(actions, p, q-1)
	sort(actions, q+1, r)
}

func partition(actions []int, p int, r int) int {
	i := p
	for j := p; j < r; j++ {
		if actions[j]%2 == 1 {
			actions[i], actions[j] = actions[j], actions[i]
			i++
		}
	}
	actions[i], actions[r] = actions[r], actions[i]
	return i
}
