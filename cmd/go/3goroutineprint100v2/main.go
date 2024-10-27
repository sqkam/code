package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1, ch2, ch3 := make(chan int, 1), make(chan int, 1), make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		i := 0
		for {
			<-ch1
			if i <= 100 {
				fmt.Printf("1 %v\n", i)
			} else {
				ch2 <- 0
				break
			}
			i += 3
			ch2 <- 0
		}
		wg.Done()
	}()

	go func() {
		i := 1
		for {
			<-ch2
			if i <= 100 {
				fmt.Printf("2 %v\n", i)
			} else {
				ch3 <- 0
				break
			}
			i += 3
			ch3 <- 0
		}
		wg.Done()
	}()

	go func() {
		i := 2
		for {
			<-ch3
			if i <= 100 {
				fmt.Printf("3 %v\n", i)
			} else {
				ch1 <- 0
				break
			}
			i += 3
			ch1 <- 0
		}
		wg.Done()
	}()

	ch1 <- 0

	wg.Wait()

}
