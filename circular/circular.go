package main

import (
	"fmt"
)

type Node struct {
	next, prev *Node
	payload interface{}
}

type List struct {
	start *Node
}

func (l *List) CreateList() {
	l.start = nil
}

func (l *List) InsertFirst(payload interface{}) {

	node := &Node {
		payload: payload,
	}

	if l.start == nil {
		l.start = node
	} else {
		node.next = l.start
		node.prev = l.start.prev
		l.start.prev.next = node
		l.start.prev = node
	}
}

func (l *List) InsertAfter(payload interface{}, searchNode interface{}) {

	node := &Node {
		payload: payload,
	}

	if l.start == nil {
		fmt.Println("List is empty")
	} else {
		n := l.Find(searchNode)
		if n.payload == -1 {
			fmt.Println("No item found")
		} else {
			node.next = n.next
			node.prev = n
			n.next.prev = node
			n.next = node
		}
	}
}

func (l *List) InsertLast(payload interface{}) {

	node := &Node {
		payload: payload,
	}

	if l.start == nil {
		l.start = node
	} else {
		node.next = l.start
		node.prev = l.start.prev
		l.start.prev.next = node
		l.start.prev = node
	}
}

func (l *List) Find(payload interface{}) (n *Node) {

	n = &Node {
		payload: -1,
	}

	p := l.start
	for p != nil {
		if p.payload == payload {
			n = p
			break
		}
		p = p.next
	}

	return n
}

func (l *List) Print() {

	p := l.start
	for p != nil {
		fmt.Print(" | ", p.payload)
		p = p.next
	}
	fmt.Print(" | ")
}

func main() {

	l := &List{}
	l.CreateList()

	l.InsertFirst(1)
	l.InsertLast(3)
	l.InsertAfter(2, 1)

	l.Print()
}