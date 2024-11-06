package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	producerCount := 2
	consumerCount := 5
	doneCh := make(chan int, producerCount)

	for range producerCount {
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

	for range producerCount {
		<-doneCh
	}
	close(ch)
	wg.Wait()

}
