package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5,Next: &ListNode{Val: 6}}}}}}
	reverseList(&list)
}

func reverseList(list *ListNode) {
	newList:=ListNode{Next: list}
	listP:=&newList
	arr:=make([]int,0)
	for listP.Next!=nil{
		arr=append(arr,listP.Next.Val)
		listP=listP.Next
	}
	newArr:=make([]int,0)
	p:=0
	q:=len(arr)-1
	for p<=q {
		if p==q {
			newArr=append(newArr,arr[p])
			break
		}
		newArr=append(newArr,arr[p])
		newArr=append(newArr,arr[q])
		p++
		q--
	}
	fmt.Printf("%v\n", newArr)

}
