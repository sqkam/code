package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1, ch2, ch3 := make(chan int, 1), make(chan int, 1), make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(3)
	i := 0

	go func() {
		for {
			<-ch1
			if i <= 100 {
				fmt.Printf("1 %v\n", i)
			} else {
				ch2 <- 0
				break
			}
			i += 1
			ch2 <- 0
		}
		wg.Done()
	}()

	go func() {
		for {
			<-ch2
			if i <= 100 {
				fmt.Printf("2 %v\n", i)
			} else {
				ch3 <- 0
				break
			}
			i += 1
			ch3 <- 0
		}
		wg.Done()
	}()

	go func() {
		for {
			<-ch3
			if i <= 100 {
				fmt.Printf("3 %v\n", i)
			} else {
				ch1 <- 0
				break
			}
			i += 1
			ch1 <- 0
		}
		wg.Done()
	}()

	ch1 <- 0

	wg.Wait()

}
