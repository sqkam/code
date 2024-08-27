package main

import "fmt"

func main() {
	go pull()
	go push()
	for {
	}
}

var msgChan = make(chan int, 1)

func pull() {
	for v:=range msgChan{
		fmt.Printf("%v\n", v)
	}

}

func push() {
	i := 0
	for {
		msgChan <- i
		i++
		if i > 2 {
			i = 0
msgChan= make(chan int)
		}
	}
}
