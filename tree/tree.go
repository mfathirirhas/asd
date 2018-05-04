package main

// insertion: every node insertion if new node is smaller then put to the left, if bigger then put to the right.
// deletion: delete by node payload

import (
	"fmt"
)

type Node struct {
	left, right *Node
	payload int
}

type Tree struct {
	root *Node
}

func (t *Tree) CreateTree() {
	t.root	= nil			
}

func (t *Tree) InsertNode(payload int) {

	node := &Node {
		payload: payload,
	}

	if t.root == nil {
		t.root = node
	} else {
		p := t.root
		for p != nil {
			if node.payload <= p.payload {
				if p.left == nil {
					p.left = node
					break
				}
				p = p.left
			} else if node.payload > p.payload {
				if p.right == nil {
					p.right = node
					break
				}
				p = p.right
			}
		}
	}
}

func (t *Tree) LevelCount(node *Node) int {
	if node == nil {
		return 0
	} else {
		level := 0
		p := t.LevelCount(node.left)
		q := t.LevelCount(node.right)
		if p > q{
			level += p+1
			return level
		} else {
			level += q+1
			return level
		}
	}
}

func (t *Tree) LeafNodeCount(node *Node) int {
	if node == nil {
		return 0
	} else if node.left == nil && node.right == nil {
		return 1
	} else {
		return t.LeafNodeCount(node.left) + t.LeafNodeCount(node.right)
	}
}

func (t *Tree) AllNodeCount(node *Node) int {
	if node == nil {
		return 0
	} else {
		count := 1
		if node.left != nil {
			count += t.AllNodeCount(node.left)
		}
		if node.right != nil {
			count += t.AllNodeCount(node.right)
		}
		return count
	}
}

// PrintPrefix print root first then left and right nodes
func (t *Tree) PrintPrefix(node *Node) {
	if node == nil {
		fmt.Print()
	} else {
		fmt.Print(node.payload, " ")
		t.PrintPrefix(node.left)
		t.PrintPrefix(node.right)
	}
}

// PrintInfix print left nodes, root and then right nodes
func (t *Tree) PrintInfix(node *Node) {
	if node == nil {
		fmt.Print()
	} else {
		t.PrintInfix(node.left)
		fmt.Print(node.payload, " ")
		t.PrintInfix(node.right)
	}
}

// PrintPostfix print from left and right leaf nodes till root
func (t *Tree) PrintPostfix(node *Node) {
	if node == nil {
		fmt.Print()
	} else {
		t.PrintPostfix(node.left)
		t.PrintPostfix(node.right)
		fmt.Print(node.payload, " ")
	}
}

func main() {
	t := &Tree{}
	t.CreateTree()

	t.InsertNode(4)
	t.InsertNode(2)
	t.InsertNode(6)
	t.InsertNode(1)
	t.InsertNode(3)
	t.InsertNode(5)
	t.InsertNode(7)

	fmt.Println("Number of Levels: ", t.LevelCount(t.root))
	fmt.Println("Number of Leaf Nodes : ", t.LeafNodeCount(t.root))
	fmt.Println("Number of All Nodes:", t.AllNodeCount(t.root))
	
	fmt.Print("Print Prefix: ")
	t.PrintPrefix(t.root)
	fmt.Println()

	fmt.Print("Print Infix: ")
	t.PrintInfix(t.root)
	fmt.Println()

	fmt.Print("Print Postfix: ")
	t.PrintPostfix(t.root)
	fmt.Println()
}