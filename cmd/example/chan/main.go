package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	chs := make([]chan int, 3)
	doneCh := make(chan int)
	for i := range chs {
		chs[i] = make(chan int)
	}
	for i := range chs {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			for {
				v, ok := <-chs[i]
				if !ok {
					return
				}
				fmt.Printf("id%d %v\n", i+1, v)
				if (v + 1) > 100 {
					doneCh <- 1
				} else {
					chs[(i+1)%3] <- v + 1
				}

			}
		}()
	}

	chs[0] <- 1
	<-doneCh
	for _, v := range chs {
		close(v)
	}
	close(doneCh)

	wg.Wait()

}
