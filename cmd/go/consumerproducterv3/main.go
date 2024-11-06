package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	productorCount := 2
	consumerCount := 5
	doneCh := make(chan int, productorCount)

	for range productorCount {
		go func() {
			for i := range 6 {
				ch <- i
			}
			doneCh <- 0
		}()

	}
	for range consumerCount {
		wg.Add(1)
		go func() {

			for v := range ch {
				fmt.Printf("%v\n", v)
			}
			wg.Done()
		}()
	}

	for range productorCount {
		<-doneCh
	}
	close(ch)
	wg.Wait()

}
