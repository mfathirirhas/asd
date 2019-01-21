package binary

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateTree(t *testing.T) {
	tree := &Tree{}
	tree.CreateTree()
	if tree.root != nil {
		t.Error("[TestCreateTree] failed creating new binary tree")
	}
}

func TestInsertNode(t *testing.T) {
	tree := &Tree{}
	tree.CreateTree()
	tree.InsertNode(4)
	tree.InsertNode(2)
	tree.InsertNode(6)
	tree.InsertNode(3)
	if tree.root.payload != 4 {
		t.Error("[TestInsertNode] failed insert node: root incorrect")
	}
	if tree.root.left.payload != 2 {
		t.Error("[TestInsertNode] failed insert node: root left child incorrect")
	}
	if tree.root.right.payload != 6 {
		t.Error("[TestInsertNode] failed insert node: root right child incorrect")
	}
	if tree.root.left.right.payload != 3 {
		t.Error("[TestInsertNode] failed insert node: root left right child incorrect")
	}
}

func TestDeleteNode(t *testing.T) {
	tree := &Tree{}
	tree.CreateTree()
	tree.InsertNode(4)
	tree.InsertNode(2)
	tree.InsertNode(6)
	tree.InsertNode(3)
	tree.DeleteNode(3)
	tree.DeleteNode(6)
	if tree.root.right != nil {
		t.Error("[TestDeleteNode] failed delete root right")
	}
	if tree.root.left.right != nil {
		t.Error("[TestDeleteNode] failed delete root left right")
	}
}

func TestFindParentNode(t *testing.T) {
	tree := &Tree{}
	tree.CreateTree()
	tree.InsertNode(4)
	tree.InsertNode(2)
	tree.InsertNode(6)
	tree.InsertNode(3)
	n1 := tree.FindParentNode(6)
	n2 := tree.FindParentNode(3)
	if n1.payload != 4 {
		t.Error("[TestFindParentNode] parent not matched")
	}
	if n2.payload != 2 {
		t.Error("[TestFindParentNode] parent not matched")
	}
}

func TestFindNode(t *testing.T) {
	tree := &Tree{}
	tree.CreateTree()
	tree.InsertNode(4)
	tree.InsertNode(2)
	tree.InsertNode(6)
	tree.InsertNode(3)
	n1 := tree.FindNode(4)
	n2 := tree.FindNode(2)
	n3 := tree.FindNode(6)
	n4 := tree.FindNode(3)
	n5 := tree.FindNode(8)
	if !n1 && !n2 && !n3 && !n4 {
		t.Error("[TestFindNode] node not found")
	}
	if n5 {
		t.Error("[TestFindNode] this node should not be found")
	}
}

func TestLevelCount(t *testing.T) {
	tree := &Tree{}
	tree.CreateTree()
	tree.InsertNode(4)
	tree.InsertNode(2)
	tree.InsertNode(6)
	tree.InsertNode(1)
	tree.InsertNode(3)
	tree.InsertNode(5)
	tree.InsertNode(7)
	level := tree.LevelCount(tree.root)
	if level != 3 {
		t.Error("[TestLevelCount] number of level incorrect")
	}
}

func TestLeafNodeCount(t *testing.T) {
	tree := &Tree{}
	tree.CreateTree()
	tree.InsertNode(4)
	tree.InsertNode(2)
	tree.InsertNode(6)
	tree.InsertNode(1)
	tree.InsertNode(3)
	tree.InsertNode(5)
	tree.InsertNode(7)
	leafs := tree.LeafNodeCount(tree.root)
	if leafs != 4 {
		t.Error("[TestLeafNodeCount] number of leafs incorrect")
	}
}

func TestAllNodeCount(t *testing.T) {
	tree := &Tree{}
	tree.CreateTree()
	tree.InsertNode(4)
	tree.InsertNode(2)
	tree.InsertNode(6)
	tree.InsertNode(1)
	tree.InsertNode(3)
	tree.InsertNode(5)
	tree.InsertNode(7)
	allNodes := tree.AllNodeCount(tree.root)
	fmt.Println(allNodes)
	if allNodes != 7 {
		t.Error("[TestAllNodeCount] number of all nodes incorrect")
	}
}

func TestGetPrefix(t *testing.T) {
	tree := &Tree{}
	tree.CreateTree()
	tree.InsertNode(4)
	tree.InsertNode(2)
	tree.InsertNode(6)
	tree.InsertNode(1)
	tree.InsertNode(3)
	tree.InsertNode(5)
	tree.InsertNode(7)
	expected := []int{4, 2, 1, 3, 6, 5, 7}
	tree.GetPrefix(tree.root)
	if !reflect.DeepEqual(PrefixList, expected) {
		t.Error("[TestGetPrefix] prefix list incorrect")
	}
}

func TestGetInfix(t *testing.T) {
	tree := &Tree{}
	tree.CreateTree()
	tree.InsertNode(4)
	tree.InsertNode(2)
	tree.InsertNode(6)
	tree.InsertNode(1)
	tree.InsertNode(3)
	tree.InsertNode(5)
	tree.InsertNode(7)
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	tree.GetInfix(tree.root)
	if !reflect.DeepEqual(InfixList, expected) {
		t.Error("[TestGetPrefix] infix list incorrect")
	}
}

func TestGetPostfix(t *testing.T) {
	tree := &Tree{}
	tree.CreateTree()
	tree.InsertNode(4)
	tree.InsertNode(2)
	tree.InsertNode(6)
	tree.InsertNode(1)
	tree.InsertNode(3)
	tree.InsertNode(5)
	tree.InsertNode(7)
	expected := []int{1, 3, 2, 5, 7, 6, 4}
	tree.GetPostfix(tree.root)
	if !reflect.DeepEqual(PostfixList, expected) {
		t.Error("[TestGetPrefix] infix list incorrect")
	}
}
