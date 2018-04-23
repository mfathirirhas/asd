package main

import (
	"fmt"
)

type Payload struct {
	name string
	age int
	address string
	phone string
}

type Node struct {
	next *Node
	payload *Payload
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

func (l *List) InsertFirst(payload *Payload) {

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

func (l *List) InsertAfter(payload *Payload, name string) {
	
	node := &Node{payload: payload}

	if l.first == nil {
		fmt.Println("List is empty")
	} else {
		n := l.FindByName(name)
		if n.payload.name == "" {
			fmt.Println("Inserting after '"+ name +"' - No data with name '"+ name +"'")
		} else {
			node.next = n.next
			n.next = node
		}
	}
}

func (l *List) InsertLast(payload *Payload) {
	
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

func (l *List) DeleteAfter(name string) {

	if l.first == nil {
		fmt.Println("List is empty")
	} else {
		n := l.FindByName(name)
		if n.payload.name == "" {
			fmt.Println("Deleting after '"+ name +"' - No data with name '"+ name +"'")
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

func (l *List) FindByName(name string) (n *Node) {

	n = &Node{
		payload: &Payload {
			name:"",
		},
	}

	p := l.first
	for p != nil {
		if p.payload.name == name {
			n = p
			break
		}
		p = p.next
	}

	return n
}

func (l *List) PrintAll() {

	p := l.first
	i := 1
	for p != nil {
		fmt.Println(i ,".", p.payload.name ,"|", p.payload.age ,"|", p.payload.address ,"|", p.payload.phone)
		i++
		p = p.next
	}
}

func main() {
	l := &List{}
	l.CreateList()

	p := &Payload {
		name: "Muhammad",
		age: 22,
		address: "ldlaksmdlkamsdlk",
		phone: "0842342342343",
	}
	l.InsertFirst(p)

	p2 := &Payload {
		name: "Irhas",
		age: 24,
		address: "jalan jalan",
		phone: "0842342342343",
	}
	l.InsertLast(p2)

	p3 := &Payload {
		name: "Fathir",
		age: 23,
		address: "ldlaksmdlkamsdlk",
		phone: "0842342342343",
	}
	l.InsertAfter(p3, "Muhammad")

	l.DeleteFirst()
	l.DeleteAfter("Muhammad")
	l.DeleteLast()

	l.PrintAll()
}