package main

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	head *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: &Node{}}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	p := this.head
	for range index {
		p = p.Next
	}
	return p.Next.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.size++
	this.head.Next = &Node{Val: val, Next: this.head.Next}
}

func (this *MyLinkedList) AddAtTail(val int) {

	p := this.head
	for range this.size {
		p = p.Next
	}
	this.size++
	p.Next = &Node{Val: val, Next: p.Next}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	p := this.head
	for range index {
		p = p.Next
	}
	this.size++
	p.Next = &Node{Val: val, Next: p.Next}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	p := this.head
	for range index {
		p = p.Next
	}
	this.size--
	p.Next = p.Next.Next
}

//
//type Node struct {
//	Val  int
//	Next *Node
//}
//type MyLinkedList struct {
//	head *Node
//	Size int
//}
//
//func Constructor() MyLinkedList {
//	n := &Node{}
//	return MyLinkedList{head: n}
//}
//
//func (this *MyLinkedList) Get(index int) int {
//	if index < 0 || index >= this.Size {
//		return -1
//	}
//
//	p := this.head
//	for range index {
//		p = p.Next
//	}
//
//	return p.Next.Val
//
//}
//
//func (this *MyLinkedList) AddAtHead(val int) {
//	this.Size++
//	this.head.Next = &Node{Val: val, Next: this.head.Next}
//
//}
//
//func (this *MyLinkedList) AddAtTail(val int) {
//	p := this.head
//	for range this.Size {
//		p = p.Next
//	}
//	this.Size++
//	p.Next = &Node{Val: val}
//}
//
//func (this *MyLinkedList) AddAtIndex(index int, val int) {
//	if index < 0 || index > this.Size {
//		return
//	}
//	p := this.head
//	for range index {
//		p = p.Next
//	}
//
//	this.Size++
//	p.Next = &Node{Val: val, Next: p.Next}
//
//	return
//}
//
//func (this *MyLinkedList) DeleteAtIndex(index int) {
//
//	if index < 0 || index >= this.Size {
//		return
//	}
//	p := this.head
//	for range index {
//		p = p.Next
//	}
//	this.Size--
//	p.Next = p.Next.Next
//	return
//}
