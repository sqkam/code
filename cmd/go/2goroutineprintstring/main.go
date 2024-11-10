package main

import (
	"fmt"
	"sync"
)

func main() {
	s := "1A2B3C4D5E"
	sBytes := []byte(s)
	lenSBytes := len(sBytes)
	ch1, ch2 := make(chan int, 1), make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	i := 0

	go func() {
		for {
			<-ch1
			if i < lenSBytes {
				fmt.Printf("1 %c\n", sBytes[i])
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
			if i < lenSBytes {
				fmt.Printf("2 %c\n", sBytes[i])
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
