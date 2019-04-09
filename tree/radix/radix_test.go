package radix

import (
	"fmt"
	"testing"
)

func TestCreateTree(t *testing.T) {
	tree := &Tree{}
	tree.CreateTree()
	if tree.root != nil {
		t.Error("[TestCreateTree] failed creating new binary tree")
	}
}

func TestSetNode(t *testing.T) {
	tree := &Tree{}
	tree.root = setNode("abc", 1, nil)
	if tree.root.key != "abc" || tree.root.val != 1 {
		t.Error("[TestSetNode] failed setting node into the root")
	}
	newNode := setNode("abcdd", 2, nil)
	if newNode.key != "abcdd" || newNode.val != 2 {
		t.Error("[TestSetNode] failed setting node")
	}
}

func TestSetEdge(t *testing.T) {
	s := "key"
	newEdge1 := setEdge(s[0], nil)
	if newEdge1.label != 'k' {
		t.Error("[TestSetEdge] failed setting edge")
	}
	newEdge2 := setEdge('k', nil)
	if newEdge2.label != 107 {
		t.Error("[TestSetEdge] failed setting edge")
	}
}

func TestUpdateNode(t *testing.T) {
	n := &Node{
		key: "key",
		val: 1,
	}
	n.updateNode("abc", 2)
	if n.key != "abc" || n.val != 2 {
		t.Error("[TestUpdateNode] failed updating node")
	}
}

func TestUpdateEdge(t *testing.T) {
	e := &Edge{
		label: 'k',
	}
	e.updateEdge('a', nil)
	if e.label != 'a' {
		t.Error("[TestUpdateNode] failed updating edge")
	}
}

func TestFindPrefix(t *testing.T) {
	incisionExpect := 4
	noIncisionExpect := 0
	partlyIncisionExpect := 3
	// both incision
	n := &Node{key: "romula"}
	key := "romuda"
	pos := n.FindPrefix(key)
	// node key incision
	n2 := &Node{key: "romula"}
	key2 := "romu"
	pos2 := n2.FindPrefix(key2)
	// input key incision
	n3 := &Node{key: "romu"}
	key3 := "romuda"
	pos3 := n3.FindPrefix(key3)
	// no incision
	n4 := &Node{key: "rick"}
	key4 := "morty"
	pos4 := n4.FindPrefix(key4)
	// partly incision
	n5 := &Node{key: "roma"}
	key5 := "romuda"
	pos5 := n5.FindPrefix(key5)
	if pos != incisionExpect && pos2 != incisionExpect && pos3 != incisionExpect && pos5 != partlyIncisionExpect && pos4 != noIncisionExpect {
		t.Error("[TestFindPrefix] Error finding prefix")
	}
}

func TestFindIncision(t *testing.T) {
	nodeKey := "romula"
	inputKey := "romuda"
	pos := 4
	matchedKey, nodeKeyDiff, inputKeyDiff := findIncision(nodeKey, inputKey, pos)
	if matchedKey != "romu" {
		t.Error("[TestFindIncision] Error invalid matchedKey")
	}
	if nodeKeyDiff != "la" {
		t.Error("[TestFindIncision] Error invalid nodeKeyDiff")
	}
	if inputKeyDiff != "da" {
		t.Error("[TestFindIncision] Error invalid inputKeyDiff")
	}
}

func TestFindEdge(t *testing.T) {
	n1 := &Node{
		key: "da",
		val: 123,
	}
	n2 := &Node{
		key: "la",
		val: 234,
	}
	n3 := &Node{
		key: "wi",
		val: 345,
	}
	e1 := &Edge{
		label: n1.key[0],
		node:  n1,
	}
	e2 := &Edge{
		label: n2.key[0],
		node:  n2,
	}
	e3 := &Edge{
		label: n3.key[0],
		node:  n3,
	}

	p := &Node{
		key:   "roma",
		val:   0,
		edges: []*Edge{e1, e2, e3},
	}
	findLabel1 := n1.key[0]
	findLabel2 := n2.key[0]
	findLabel3 := n3.key[0]
	idx1 := p.findEdge(findLabel1)
	if idx1 != e1 {
		t.Error("[TestFindEdge] Error invalid edge")
	}
	idx2 := p.findEdge(findLabel2)
	if idx2 != e2 {
		t.Error("[TestFindEdge] Error invalid edge")
	}
	idx3 := p.findEdge(findLabel3)
	if idx3 != e3 {
		t.Error("[TestFindEdge] Error invalid edge")
	}
}

func TestAddEdge(t *testing.T) {
	n := &Node{
		key: "abc",
		val: 1,
	}
	kid1 := &Node{
		key: "abcd",
		val: 2,
	}
	kid2 := &Node{
		key: "abce",
		val: 3,
	}
	e1 := setEdge('d', kid1)
	e2 := setEdge('e', kid2)
	n.addEdge(e1)
	n.addEdge(e2)
	if n.edges[0].label != 'd' && n.edges[1].label != 'd' {
		t.Error("[TestAddEdge] Error invalid edge")
	}
}

func TestCheckEdges(t *testing.T) {
	n := &Node{
		key: "abc",
		val: 1,
	}
	kid1 := &Node{
		key: "def",
		val: 2,
	}
	kid2 := &Node{
		key: "efg",
		val: 3,
	}
	e1 := setEdge(kid1.key[0], kid1)
	e2 := setEdge(kid2.key[1], kid2)
	n.edges = append(n.edges, e1)
	n.edges = append(n.edges, e2)
	if c := n.checkEdges(); !c {
		t.Error("[TestCheckEdges] Error should has child")
	}
}

func TestInsertTree(t *testing.T) {
	// test 1
	fmt.Println("---test 1---")
	key1, val1 := "romuda", 1
	key2, val2 := "romula", 2
	key3, val3 := "romawi", 3
	key4, val4 := "fizzbuzz", 4
	tree := &Tree{}
	tree.CreateTree()
	tree.InsertTree(key1, val1)
	tree.InsertTree(key2, val2)
	tree.InsertTree(key3, val3)
	tree.InsertTree(key4, val4)
	fmt.Println(tree.root.key)
	for i := 0; i < len(tree.root.edges); i++ {
		fmt.Println(string(tree.root.edges[i].label), " | ", tree.root.edges[i].node.key, " | ", tree.root.edges[i].node.val)
	}
	fmt.Println(string(tree.root.edges[1].node.edges[0].label), " | ", tree.root.edges[1].node.edges[0].node.key, " | ", tree.root.edges[1].node.edges[0].node.val)
	if tree.root.edges[0].label != 'f' || tree.root.edges[0].node.key != "fizzbuzz" || tree.root.edges[0].node.val != 4 {
		t.Error("Insertion Failed")
	}
	if tree.root.key != "" {
		t.Error("Root should be empty")
	}
	if tree.root.edges[1].node.edges[0].label != 'a' || tree.root.edges[1].node.edges[0].node.key != "awi" || tree.root.edges[1].node.edges[0].node.val != 3 {
		t.Error("Insertion Failed")
	}

	// test 2
	fmt.Println("---test 2---")
	key5, val5 := "romuda", 5
	key6, val6 := "romusa", 6
	key7, val7 := "romudala", 7
	key8, val8 := "romudalax", 8
	key9, val9 := "romawi", 9
	tree2 := &Tree{}
	tree2.CreateTree()
	tree2.InsertTree(key5, val5)
	tree2.InsertTree(key6, val6)
	tree2.InsertTree(key7, val7)
	tree2.InsertTree(key8, val8)
	tree2.InsertTree(key9, val9)
	fmt.Println(tree2.root.key, " | ", tree2.root.val, " | ", tree2.root.isLeaf)
	fmt.Println(tree2.root.edges[0].node.key, " | ", tree2.root.edges[0].node.val, " | ", tree2.root.edges[0].node.isLeaf)
	fmt.Println(tree2.root.edges[1].node.key, " | ", tree2.root.edges[1].node.val, " | ", tree2.root.edges[1].node.isLeaf)
	fmt.Println(tree2.root.edges[1].node.edges[0].node.key, " | ", tree2.root.edges[1].node.edges[0].node.val, " | ", tree2.root.edges[1].node.edges[0].node.isLeaf)
	fmt.Println(tree2.root.edges[1].node.edges[0].node.edges[0].node.key, " | ", tree2.root.edges[1].node.edges[0].node.edges[0].node.val, " | ", tree2.root.edges[1].node.edges[0].node.edges[0].node.isLeaf)
	fmt.Println(tree2.root.edges[1].node.edges[0].node.edges[0].node.edges[0].node.key, " | ", tree2.root.edges[1].node.edges[0].node.edges[0].node.edges[0].node.val, " | ", tree2.root.edges[1].node.edges[0].node.edges[0].node.edges[0].node.isLeaf)
	key10, val10 := "roma", 10
	tree2.InsertTree(key10, val10)
	fmt.Println(tree2.root.edges[0].node.key, " | ", tree2.root.edges[0].node.val, " | ", tree2.root.edges[0].node.isLeaf)
	fmt.Println(tree2.root.edges[0].node.edges[0].node.key, " | ", tree2.root.edges[0].node.edges[0].node.val, " | ", tree2.root.edges[0].node.edges[0].node.isLeaf)
	if tree2.root.key != "rom" {
		t.Error("Root should be 'rom'")
	}
	if tree2.root.edges[0].label != 'a' || tree2.root.edges[0].node.key != "a" || tree2.root.edges[0].node.val != 10 {
		t.Error("Insertion Failed")
	}
	if tree2.root.edges[1].node.edges[0].node.edges[0].node.key != "la" {
		t.Error("Insertion Failed")
	}

	// test 3 for empty root
	fmt.Println("---test 3---")
	key11, val11 := "romuda", 10
	key12, val12 := "romula", 11
	key13, val13 := "foobar", 12
	key14, val14 := "foo", 13
	tree3 := &Tree{}
	tree3.CreateTree()
	tree3.InsertTree(key11, val11)
	tree3.InsertTree(key12, val12)
	tree3.InsertTree(key13, val13)
	tree3.InsertTree(key14, val14)
	fmt.Println("root: ", tree3.root.key)
	fmt.Println(tree3.root.edges[0].node.key, " | ", tree3.root.edges[0].node.val)
	fmt.Println(tree3.root.edges[0].node.edges[0].node.key, " | ", tree3.root.edges[0].node.edges[0].node.val)
	fmt.Println(tree3.root.edges[1].node.key, " | ", tree3.root.edges[1].node.val)
	fmt.Println(tree3.root.edges[1].node.edges[0].node.key, " | ", tree3.root.edges[1].node.edges[0].node.val)
	fmt.Println(tree3.root.edges[1].node.edges[1].node.key, " | ", tree3.root.edges[1].node.edges[1].node.val)
	key15, val15 := "test", 14
	tree3.InsertTree(key15, val15)
	fmt.Println(tree3.root.edges[2].node.key, " | ", tree3.root.edges[2].node.val)
	if tree3.root.edges[0].label != 'f' || tree3.root.edges[0].node.key != "foo" || tree3.root.edges[0].node.val != 13 {
		t.Error("Insertion Failed")
	}
	if tree3.root.edges[1].label != 'r' || tree3.root.edges[1].node.key != "romu" || tree3.root.edges[1].node.val != nil {
		t.Error("Insertion Failed")
	}
	if tree3.root.edges[1].node.edges[0].label != 'd' || tree3.root.edges[1].node.edges[0].node.key != "da" || tree3.root.edges[1].node.edges[0].node.val != 10 {
		t.Error("Insertion Failed")
	}
	if tree3.root.edges[2].label != 't' || tree3.root.edges[2].node.key != "test" || tree3.root.edges[2].node.val != 14 {
		t.Error("Insertion Failed")
	}
}

func TestLookUpTree(t *testing.T) {
	key5, val5 := "romuda", 5
	key6, val6 := "romusa", 6
	key7, val7 := "romudala", 7
	key8, val8 := "romudalax", 8
	key9, val9 := "romawi", 9
	key10, val10 := "roma", 10
	tree2 := &Tree{}
	tree2.CreateTree()
	tree2.InsertTree(key5, val5)
	tree2.InsertTree(key6, val6)
	tree2.InsertTree(key7, val7)
	tree2.InsertTree(key8, val8)
	tree2.InsertTree(key9, val9)
	tree2.InsertTree(key10, val10)
	checkVal5 := tree2.LookUpTree(key5)
	checkVal6 := tree2.LookUpTree(key6)
	checkVal7 := tree2.LookUpTree(key7)
	checkVal8 := tree2.LookUpTree(key8)
	checkVal9 := tree2.LookUpTree(key9)
	checkVal10 := tree2.LookUpTree(key10)
	if checkVal5 != val5 && checkVal6 != val6 && checkVal7 != val7 && checkVal8 != val8 && checkVal9 != val9 && checkVal10 != val10 {
		t.Error("Error LookUp Values")
	}
}

func TestGetKeys(t *testing.T) {
	key5, val5 := "romuda", 5
	key6, val6 := "romusa", 6
	key7, val7 := "romudala", 7
	key8, val8 := "romudalax", 8
	key9, val9 := "romawi", 9
	key10, val10 := "roma", 10
	tree2 := &Tree{}
	tree2.CreateTree()
	tree2.InsertTree(key5, val5)
	tree2.InsertTree(key6, val6)
	tree2.InsertTree(key7, val7)
	tree2.InsertTree(key8, val8)
	tree2.InsertTree(key9, val9)
	tree2.InsertTree(key10, val10)
	var l []string
	keys := tree2.getKeys(tree2.root, "", l)
	if keys[0] != "roma" && keys[1] != "romawi" && keys[2] != "romuda" && keys[3] != "romudala" && keys[4] != "romudalax" && keys[5] != "romusa" {
		t.Error("Error keys are not valid")
	}
}

func TestGetValues(t *testing.T) {
	key5, val5 := "romuda", 5
	key6, val6 := "romusa", 6
	key7, val7 := "romudala", 7
	key8, val8 := "romudalax", 8
	key9, val9 := "romawi", 9
	key10, val10 := "roma", 10
	tree2 := &Tree{}
	tree2.CreateTree()
	tree2.InsertTree(key5, val5)
	tree2.InsertTree(key6, val6)
	tree2.InsertTree(key7, val7)
	tree2.InsertTree(key8, val8)
	tree2.InsertTree(key9, val9)
	tree2.InsertTree(key10, val10)
	var l []interface{}
	values := tree2.getValues(tree2.root, l)
	if values[0].(int) != 10 && values[1].(int) != 9 && values[2].(int) != 5 && values[3].(int) != 7 && values[4].(int) != 8 && values[5].(int) != 6 {
		t.Error("Error values are not valid")
	}
}

func TestGetKeysValues(t *testing.T) {
	key5, val5 := "romuda", 5
	key6, val6 := "romusa", 6
	key7, val7 := "romudala", 7
	key8, val8 := "romudalax", 8
	key9, val9 := "romawi", 9
	key10, val10 := "roma", 10
	tree2 := &Tree{}
	tree2.CreateTree()
	tree2.InsertTree(key5, val5)
	tree2.InsertTree(key6, val6)
	tree2.InsertTree(key7, val7)
	tree2.InsertTree(key8, val8)
	tree2.InsertTree(key9, val9)
	tree2.InsertTree(key10, val10)
	m := make(map[string]interface{})
	kv := tree2.getKeysValues(tree2.root, "", m)
	if kv[key5] != val5 && kv[key6] != val6 && kv[key7] != val7 && kv[key8] != val8 && kv[key9] != val9 && kv[key10] != val10 {
		t.Error("Error keys values are not valid")
	}
}

func TestGetAllKeys(t *testing.T) {
	key5, val5 := "romuda", 5
	key6, val6 := "romusa", 6
	key7, val7 := "romudala", 7
	key8, val8 := "romudalax", 8
	key9, val9 := "romawi", 9
	key10, val10 := "roma", 10
	tree2 := &Tree{}
	tree2.CreateTree()
	tree2.InsertTree(key5, val5)
	tree2.InsertTree(key6, val6)
	tree2.InsertTree(key7, val7)
	tree2.InsertTree(key8, val8)
	tree2.InsertTree(key9, val9)
	tree2.InsertTree(key10, val10)
	keys := tree2.GetAllKeys()
	if keys[0] != "roma" && keys[1] != "romawi" && keys[2] != "romuda" && keys[3] != "romudala" && keys[4] != "romudalax" && keys[5] != "romusa" {
		t.Error("Error keys are not valid")
	}
}

func TestGetAllValues(t *testing.T) {
	key5, val5 := "romuda", 5
	key6, val6 := "romusa", 6
	key7, val7 := "romudala", 7
	key8, val8 := "romudalax", 8
	key9, val9 := "romawi", 9
	key10, val10 := "roma", 10
	tree2 := &Tree{}
	tree2.CreateTree()
	tree2.InsertTree(key5, val5)
	tree2.InsertTree(key6, val6)
	tree2.InsertTree(key7, val7)
	tree2.InsertTree(key8, val8)
	tree2.InsertTree(key9, val9)
	tree2.InsertTree(key10, val10)
	values := tree2.GetAllValues()
	if values[0].(int) != 10 && values[1].(int) != 9 && values[2].(int) != 5 && values[3].(int) != 7 && values[4].(int) != 8 && values[5].(int) != 6 {
		t.Error("Error values are not valid")
	}
}

func TestGetAllKeysValues(t *testing.T) {
	key5, val5 := "romuda", 5
	key6, val6 := "romusa", 6
	key7, val7 := "romudala", 7
	key8, val8 := "romudalax", 8
	key9, val9 := "romawi", 9
	key10, val10 := "roma", 10
	tree2 := &Tree{}
	tree2.CreateTree()
	tree2.InsertTree(key5, val5)
	tree2.InsertTree(key6, val6)
	tree2.InsertTree(key7, val7)
	tree2.InsertTree(key8, val8)
	tree2.InsertTree(key9, val9)
	tree2.InsertTree(key10, val10)
	kv := tree2.GetAllKeysValues()
	if kv[key5] != val5 && kv[key6] != val6 && kv[key7] != val7 && kv[key8] != val8 && kv[key9] != val9 && kv[key10] != val10 {
		t.Error("Error keys values are not valid")
	}
}
