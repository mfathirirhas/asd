package main

import (
	"fmt"
)

type ParentNode struct {
	nextP, prevP *ParentNode
	parentPayload interface{}
	childList *ChildList
}

type ChildNode struct {
	nextC, prevC *ChildNode
	childPayload interface{}
}

type ParentList struct {
	firstP, lastP *ParentNode
}

type ChildList struct {
	firstC, lastC *ChildNode
}

func (l *ParentList) CreateParentList() {
	l.firstP = nil
	l.lastP  = nil
}

func (l *ParentList) InsertFirstParent(payload interface{}) {

	node := &ParentNode {
		parentPayload: payload,
		childList: &ChildList {
			firstC: nil,
			lastC: nil,
		},
	}

	if l.firstP == nil {
		l.firstP = node
		l.lastP = node
	} else {
		node.nextP = l.firstP
		l.firstP.prevP = node
		l.firstP = node
	}
}

func (l *ParentList) InsertAfterParent(payload interface{}, search interface{}) {
	
	node := &ParentNode {
		parentPayload: payload,
		childList: &ChildList {
			firstC: nil,
			lastC: nil,
		},
	}

	n := l.FindParent(search)
	if n.parentPayload == -1 {
		fmt.Println("Inserting after '",search,"' - No data with name '", search ,"'")
	} else {
		node.nextP = n.nextP
		n.nextP = node
	}
}

func (l *ParentList) InsertLastParent(payload interface{}) {

	node := &ParentNode {
		parentPayload: payload,
		childList: &ChildList {
			firstC: nil,
			lastC: nil,
		},
	}

	if l.firstP == nil {
		l.firstP = node
		l.lastP = node
	} else {
		node.prevP = l.lastP
		l.lastP.nextP = node
		l.lastP = node
	}

}

func (l *ParentList) InsertFirstChild(parent interface{} ,payload interface{}) {

	node := &ChildNode {
		childPayload: payload,
	}

	parentNode := l.FindParent(parent)
	if parentNode.parentPayload != -1 {
		if parentNode.childList.firstC == nil {
			parentNode.childList.firstC = node
			parentNode.childList.lastC = node
		} else {
			node.nextC = parentNode.childList.firstC
			parentNode.childList.firstC.prevC = node
			parentNode.childList.firstC = node
		}
	} else {
		fmt.Print("No parent found")
	}
}

func (l *ParentList) FindParent(payload interface{}) (n *ParentNode) {

	n = &ParentNode {
		parentPayload: -1,
	}

	p := l.firstP
	for p != nil {
		if p.parentPayload == payload {
			n = p
		}
		p = p.nextP
	}

	return n
}

func (l *ParentList) PrintAllParent() {

	p := l.firstP
	for p != nil {
		fmt.Print(" | " ,p.parentPayload)
		p = p.nextP
	}
	fmt.Print(" | ")
}

func (l *ParentList) PrintAll() {

	p := l.firstP
	for p != nil {
		fmt.Print(" | ", p.parentPayload , " | ",  " -> ")
		pc := p.childList.firstC
		for pc != nil {
			fmt.Print(" | ", pc.childPayload)
			pc = pc.nextC
		}
		fmt.Println(" | ")
		p = p.nextP
	}
}

func main() {
	l := &ParentList{}
	l.CreateParentList()
	l.InsertFirstParent(1)
	l.InsertFirstParent(2)
	l.InsertFirstParent(3)

	l.InsertAfterParent(9, 3)
	l.InsertAfterParent(4,9)

	l.InsertLastParent(99)
	l.InsertLastParent(100)

	l.InsertFirstChild(3, 4)
	l.InsertFirstChild(3, 5)
	l.InsertFirstChild(2, 90)
	l.InsertFirstChild(1, 9)
	l.InsertFirstChild(1, 4)
	l.InsertFirstChild(1, 5)

	// l.PrintAllParent()
	l.PrintAll()
}