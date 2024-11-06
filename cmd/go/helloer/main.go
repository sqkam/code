package main

import "fmt"

type HelloEr interface {
	SayHello()
}

type DefaultHelloEr struct {
}

func (d *DefaultHelloEr) SayHello() {
	fmt.Printf("%v\n", "hello world")
}
func NewDefaultHelloEr() HelloEr {
	return &DefaultHelloEr{}
}

type LoggerHelloEr struct {
	h HelloEr
}

func (l *LoggerHelloEr) SayHello() {
	fmt.Printf("%v\n", "log start ")
	l.h.SayHello()
	fmt.Printf("%v\n", "log end")
}
func NewLoggerHelloEr(h HelloEr) HelloEr {
	return &LoggerHelloEr{h: h}
}

func main() {
	h := NewLoggerHelloEr(NewDefaultHelloEr())
	h.SayHello()

}
