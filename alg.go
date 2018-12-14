package main

import "fmt"

// LinkedList in go
type LinkedList struct {
	val  int
	next *LinkedList
}

// Init just 0 length list
func (head *LinkedList) Init() {
	head.val = 0
	head.next = nil
}

// AddAt add val at pos, and will return immediately if pos is larger than len of the list
func (head *LinkedList) AddAt(pos int, val int) {
	if pos > head.val { // for head node, the val record the length
		return // need add error returning
	}

	p := head

	for i := 0; i < pos; {
		p = p.next
	}
	p.next = &LinkedList{val, p.next}
	head.val++
}

func helloWorld() int {
	fmt.Println("Hello")
	return 0
}
