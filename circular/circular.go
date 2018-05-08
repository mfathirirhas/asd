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
	end *Node
}


// methods

func (l *List) CreateList() {
	l.start = nil
	l.end 	= nil
}

func (l *List) InsertFirst(payload interface{}) {

	node := &Node{payload: payload}
	if l.start == nil {
		l.start = node
		l.end 	= node
		node.next = node
		node.prev = node
	} else {
		node.next = l.start
		l.start.prev = node
		l.start = node
		p := l.start
		for p.next != nil {
			if p == l.end {
				p.next = l.start
				break
			}
			p = p.next
		}
	}
}

func (l *List) InsertAfter(payload interface{}, search interface{}) {
	
	node := &Node{payload: payload}

	if l.start == nil {
		fmt.Println("List is empty")
	} else {
		n := l.Find(search)
		if n.payload == -1 {
			fmt.Println("Inserting after '", search ,"' - No data with name '", search ,"'")
		} else {
			node.next = n.next
			n.next.prev = node
			n.next = node
		}
	}
}

func (l *List) InsertLast(payload interface{}) {
	
	node := &Node{payload: payload}
	if l.start == nil {
		l.start = node
		l.end  = node
		node.next = node
		node.prev = node
	} else {
		l.end.next = node
		l.start.prev = node
		node.prev = l.end
		node.next = l.start
		l.end = node
	}
}

func (l *List) DeleteFirst() {

	if l.start == nil {
		fmt.Println("List is empty")
	} else {
		p := l.start
		l.start = p.next
		l.start.prev = l.end
		l.end.next = l.start
		p.next = nil
		l.start.prev = nil
	}
}

func (l *List) DeleteAfter(search interface{}) {

	if l.start == nil {
		fmt.Println("List is empty")
	} else {
		n := l.Find(search)
		if n.payload == -1 {
			fmt.Println("Deleting after '", search ,"' - No data with name '", search ,"'")
		} else {
			p := n.next
			n.next = p.next
			p.next.prev = n
			p.next = nil
			p.prev = nil			
		}
	}
}

func (l *List) DeleteLast() {

	if l.start == nil {
		fmt.Println("List is empty")
	} else {
		p := l.end
		l.end = p.prev
		l.end.next = l.start
		l.start.prev = l.end
		p.prev.next = nil
		p.prev = nil
	}
}

func (l *List) Find(search interface{}) (n *Node) {

	n = &Node{
		payload: -1,
	}

	p := l.start
	for p != nil {
		if p.payload == search {
			n = p
			break
		}

		if p == l.end {
			break
		}
		p = p.next
	}

	return n
}

func (l *List) PrintAll() {

	p := l.start
	for p != nil {
		fmt.Print("| " ,p.payload, " ")
		if p == l.end{
			break
		}
		p = p.next
	}
	fmt.Print("|", "\n")
}

func main() {
	l := &List{}
	l.CreateList()
 
	l.InsertFirst(3)
	l.InsertFirst(1)
	l.InsertLast(4)
	l.InsertLast(5)
	l.InsertLast(6)
	l.InsertAfter(2, 1)
	l.InsertAfter(2.5, 2)

	l.DeleteFirst()
	l.DeleteAfter(1)
	l.DeleteLast()

	l.PrintAll()
}