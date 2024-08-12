package main

func main() {
	myQueue := Constructor()
	myQueue.Push(1) // queue is: [1]
	myQueue.Push(2) // queue is: [1, 2] (leftmost is front of the queue)
	myQueue.Peek()  // return 1
	myQueue.Pop()   // return 1, queue is [2]
	myQueue.Empty() // return false
}

type MyQueue struct {
	InStack  []int
	OutStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) In2Out() {
	for len(this.InStack) > 0 {
		this.OutStack = append(this.OutStack, this.InStack[len(this.InStack)-1])
		this.InStack = this.InStack[:len(this.InStack)-1]
	}

}
func (this *MyQueue) Push(x int) {
	this.InStack = append(this.InStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.OutStack) == 0 {
		this.In2Out()
	}

	val := this.OutStack[len(this.OutStack)-1]
	this.OutStack = this.OutStack[:len(this.OutStack)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.OutStack) == 0 {
		this.In2Out()
	}
	return this.OutStack[len(this.OutStack)-1]

}

func (this *MyQueue) Empty() bool {
	return len(this.InStack) == 0 && len(this.OutStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
