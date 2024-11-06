package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	producerCount := 2
	consumerCount := 5
	doneCh := make(chan int, producerCount)

	for i := 0; i < producerCount; i++ {
		//生产者
		go func() {
			for v := range 6 {
				ch <- v
			}
			doneCh <- 1
		}()
	}
	for i := 0; i < consumerCount; i++ {
		wg.Add(1)
		//消费者
		go func() {
			for {
				time.Sleep(time.Second)
				v, ok := <-ch
				if !ok {
					wg.Done()
					return
				}
				fmt.Printf("消费者%v 输出: %v\n", i, v)
			}
		}()
	}

	for range producerCount {
		<-doneCh
	}

	close(ch)
	wg.Wait()
}
