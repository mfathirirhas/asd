package main

import (
	"fmt"
)

type Node struct {
	next *Node
	payload interface{}
}

type List struct {
	first *Node
	last *Node
}


// methods

func (l *List) CreateList() {
	l.first = nil
	l.last 	= nil
}

func (l *List) InsertFirst(payload interface{}) {

	node := &Node{payload: payload}
	if l.first == nil {
		l.first = node
		l.last 	= node
	} else {
		node.next = l.first
		l.first = node
		p := l.first
		for p.next != nil {
			p = p.next
		}
		l.last = p
	}
}

func (l *List) InsertAfter(payload interface{}, search interface{}) {
	
	node := &Node{payload: payload}

	if l.first == nil {
		fmt.Println("List is empty")
	} else {
		n := l.Find(search)
		if n.payload == -1 {
			fmt.Println("Inserting after '", search ,"' - No data with name '", search ,"'")
		} else {
			node.next = n.next
			n.next = node
		}
	}
}

func (l *List) InsertLast(payload interface{}) {
	
	node := &Node{payload: payload}
	if l.first == nil {
		l.first = node
		l.last  = node
	} else {
		l.last.next = node
		l.last = node
	}
}

func (l *List) DeleteFirst() {

	if l.first == nil {
		fmt.Println("List is empty")
	} else {
		p := l.first
		l.first = p.next
		p.next = nil
	}
}

func (l *List) DeleteAfter(search interface{}) {

	if l.first == nil {
		fmt.Println("List is empty")
	} else {
		n := l.Find(search)
		if n.payload == -1 {
			fmt.Println("Deleting after '", search ,"' - No data with name '", search ,"'")
		} else {
			p := n.next
			n.next = p.next
			p.next = nil
		}
	}
}

func (l *List) DeleteLast() {

	if l.first == nil {
		fmt.Println("List is empty")
	} else {
		p := l.first
		for p != nil {
			if p.next == l.last {
				l.last = p
				p.next = nil
			}
			p = p.next
		}
	}
}

func (l *List) Find(search interface{}) (n *Node) {

	n = &Node{
		payload: -1,
	}

	p := l.first
	for p != nil {
		if p.payload == search {
			n = p
			break
		}
		p = p.next
	}

	return n
}

func (l *List) PrintAll() {

	p := l.first
	for p != nil {
		fmt.Print("| " ,p.payload, " ")
		p = p.next
	}
	fmt.Print("|", "\n")
}

func main() {
	l := &List{}
	l.CreateList()
 
	l.InsertFirst(1)
	l.InsertLast(3)
	l.InsertAfter(2, 1)

	l.DeleteFirst()
	l.DeleteAfter(1)
	l.DeleteLast()

	l.PrintAll()
}