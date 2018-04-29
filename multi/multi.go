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

func (l *ParentList) DeleteFirstParent() {

	if l.firstP == nil {
		fmt.Println("Parent List is empty")
	} else {
		p := l.firstP
		l.firstP = p.nextP
		p.nextP.prevP = nil
		p.nextP = nil
	}
}

func (l *ParentList) DeleteAfterParent(searchParent interface{}) {

	if l.firstP == nil {
		fmt.Println("Parent list is empty")
	} else {
		n := l.FindParent(searchParent)
		if n.parentPayload == -1 {
			fmt.Println("No parent found")
		} else {
			p := n.nextP
			n.nextP = p.nextP
			p.nextP.prevP = n
			p.nextP = nil
			p.prevP = nil
		}
	}
}

func (l *ParentList) DeleteLastParent() {

	if l.firstP == nil {
		fmt.Println("Parent list is empty")
	} else {
		p := l.lastP
		l.lastP = p.prevP
		p.prevP.nextP = nil
		p.prevP = nil
	}
}

func (l *ParentList) InsertFirstChild(childPayload interface{}, searchParent interface{}) {

	node := &ChildNode {
		childPayload: childPayload,
	}

	parentNode := l.FindParent(searchParent)
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
		fmt.Print("No parent found \n")
	}
}

func (l *ParentList) InsertAfterChild(child interface{}, searchChild interface{}, searchParent interface{}) {

	node := &ChildNode {
		childPayload: child,
	}

	parentNode := l.FindParent(searchParent)
	if parentNode.parentPayload != -1 {
		childNode := l.FindChild(searchChild, parentNode)
		if childNode.childPayload != -1 {
			node.nextC = childNode.nextC
			node.prevC = childNode
			childNode.nextC.prevC = node
			childNode.nextC = node
		} else {
			fmt.Print("No child found \n")
		}
	} else {
		fmt.Print("No parent found \n")
	}
}

func (l *ParentList) InsertLastChild(childPayload interface{}, searchParent interface{}) {

	node := &ChildNode {
		childPayload: childPayload,
	}

	parentNode := l.FindParent(searchParent)
	if parentNode.parentPayload != -1 {
		node.prevC = parentNode.childList.lastC
		parentNode.childList.lastC.nextC = node
		parentNode.childList.lastC = node
	} else {
		fmt.Print("Inserting child to last in parent ", searchParent ," - No parent found \n")
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
			break
		}
		p = p.nextP
	}

	return n
}

func (l *ParentList) FindChild(childPayload interface{}, parentNode *ParentNode) (cn *ChildNode) {

	cn = &ChildNode {
		childPayload: -1,
	}

	p := parentNode.childList.firstC
	for p != nil {
		if p.childPayload == childPayload {
			cn = p
			break
		}
		p = p.nextC
	}

	return cn
}

func (l *ParentList) PrintAllParent() {

	p := l.firstP
	for p != nil {
		fmt.Print(" | " ,p.parentPayload)
		p = p.nextP
	}
	fmt.Print(" | ", "\n")
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

	// parent
	l.InsertFirstParent(1)
	l.InsertFirstParent(2)
	l.InsertFirstParent(3)

	l.InsertAfterParent(9, 3)
	l.InsertAfterParent(4,9)

	l.InsertLastParent(99)
	l.InsertLastParent(100)

	l.DeleteFirstParent()
	l.DeleteAfterParent(2)
	l.DeleteLastParent()

	// child
	// l.InsertFirstChild(4, 3)
	// l.InsertFirstChild(5, 3)
	// l.InsertFirstChild(90, 2)
	// l.InsertFirstChild(9, 1)
	// l.InsertFirstChild(4, 1)
	// l.InsertFirstChild(5, 1)

	// l.InsertAfterChild(6, 5, 3)
	// l.InsertAfterChild(7, 6, 3)
	// l.InsertAfterChild(6, 5, 1)
	// l.InsertAfterChild(8, 4, 1)

	// l.InsertLastChild(10, 1)
	// l.InsertLastChild(11, 1)

	l.PrintAllParent()
	l.PrintAll()
}