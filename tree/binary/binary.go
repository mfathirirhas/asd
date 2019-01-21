package binary

// insertion: every node insertion if new node is smaller then put to the left, if bigger then put to the right.
// deletion: delete by node payload

import (
	"fmt"
)

var (
	PrefixList  []int
	InfixList   []int
	PostfixList []int
)

type Node struct {
	left, right *Node
	payload     int
}

type Tree struct {
	root *Node
}

func (t *Tree) CreateTree() {
	t.root = nil
}

func (t *Tree) InsertNode(payload int) {

	node := &Node{
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

func (t *Tree) DeleteNode(payload int) {

	if t.root == nil {
		fmt.Println("Tree is empty")
	} else {
		if payload == t.root.payload {
			t.root = nil
			return
		}
		n := t.FindParentNode(payload)
		if n == nil {
			fmt.Println("No node found")
		} else {
			if payload <= n.payload {
				n.left = nil
			} else {
				n.right = nil
			}
		}
	}
}

func (t *Tree) FindParentNode(payload int) (n *Node) {

	node := &Node{
		payload: payload,
	}

	if t.root == nil {
		fmt.Println("Tree is empty")
	} else {
		p := t.root
		for p != nil {
			if p.left != nil && node.payload == p.left.payload {
				n = p
				break
			}

			if p.right != nil && node.payload == p.right.payload {
				n = p
				break
			}

			if node.payload <= p.payload {
				p = p.left
			} else {
				p = p.right
			}
		}
	}

	return n
}

func (t *Tree) FindNode(payload int) bool {

	node := &Node{
		payload: payload,
	}

	if t.root == nil {
		fmt.Println("Tree is empty")
	} else {
		p := t.root
		for p != nil {
			if p.payload == payload {
				return true
			}

			if node.payload <= p.payload {
				p = p.left
			} else {
				p = p.right
			}
		}
	}
	return false
}

func (t *Tree) LevelCount(node *Node) int {
	if node == nil {
		return 0
	} else {
		level := 0
		p := t.LevelCount(node.left)
		q := t.LevelCount(node.right)
		if p > q {
			level += p + 1
			return level
		} else {
			level += q + 1
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

func (t *Tree) GetPrefix(node *Node) {
	if node != nil {
		PrefixList = append(PrefixList, node.payload)
		t.GetPrefix(node.left)
		t.GetPrefix(node.right)
	}
}

// PrintPrefix print root first then left and right nodes
func (t *Tree) PrintPrefix(node *Node) {
	if len(PrefixList) > 0 {
		for _, v := range PrefixList {
			fmt.Print(v, " ")
		}
	} else {
		fmt.Println("PrefixList is empty")
	}
}

func (t *Tree) GetInfix(node *Node) {
	if node != nil {
		t.GetInfix(node.left)
		InfixList = append(InfixList, node.payload)
		t.GetInfix(node.right)
	}
}

// PrintInfix print left nodes, root and then right nodes
func (t *Tree) PrintInfix(node *Node) {
	if len(InfixList) > 0 {
		for _, v := range InfixList {
			fmt.Print(v, " ")
		}
	} else {
		fmt.Println("InfixList is empty")
	}
}

func (t *Tree) GetPostfix(node *Node) {
	if node != nil {
		t.GetPostfix(node.left)
		t.GetPostfix(node.right)
		PostfixList = append(PostfixList, node.payload)
	}
}

// PrintPostfix print from left and right leaf nodes till root
func (t *Tree) PrintPostfix(node *Node) {
	if len(PostfixList) > 0 {
		for _, v := range PostfixList {
			fmt.Print(v, " ")
		}
	} else {
		fmt.Println("PostfixList is empty")
	}
}
