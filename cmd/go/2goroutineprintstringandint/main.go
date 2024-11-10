package main

import (
	"fmt"
	"sync"
)

func main() {

	ch1, ch2 := make(chan int, 1), make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		i := 1
		for {
			<-ch1
			if i < 4 {
				fmt.Printf("1 %d\n", i)
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
		i := 'a'
		for {
			<-ch2
			if i < 'd' {
				fmt.Printf("2 %c\n", i)
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
