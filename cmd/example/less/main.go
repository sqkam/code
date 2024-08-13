package main

import (
	"fmt"
	uuid "github.com/lithammer/shortuuid/v4"
	"math/rand"
	"slices"
)

type Node struct {
	Age  int64
	Name string
}

func main() {
	s := make([]Node, 0)
	s = append(s,
		Node{
			Age:  rand.Int63n(99),
			Name: uuid.New(),
		},
		Node{
			Age:  rand.Int63n(99),
			Name: uuid.New(),
		},
		Node{
			Age:  rand.Int63n(99),
			Name: uuid.New(),
		},
		Node{
			Age:  rand.Int63n(99),
			Name: uuid.New(),
		},
	)

	slices.SortFunc(s, func(a, b Node) int {
		if a.Age != b.Age {
			return int(a.Age - b.Age)
		}
		return len(a.Name) - len(b.Name)
	})
	fmt.Printf("%v\n", s)
}
