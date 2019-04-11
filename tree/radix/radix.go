package radix

import (
	"sort"
)

type Node struct {
	key    string      // define every unique path in the tree
	val    interface{} // define the value of every key/path
	edges  Edges       // slice of all the edges each node has
	isLeaf bool        // is node has value
}

type Edge struct {
	label byte  // first char of node key it's pointing to
	node  *Node // each edge connect to one node
}

type Edges []*Edge

type Tree struct {
	root *Node // define the root of tree
}

func (t *Tree) CreateTree() {
	t.root = nil
}

// to order slice
func (e Edges) Len() int {
	return len(e)
}

func (e Edges) Less(i, j int) bool {
	return e[i].label < e[j].label
}

func (e Edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// end ---

func setNode(key string, val interface{}, edges Edges) *Node {
	isLeaf := false
	if val != nil {
		isLeaf = true
	}
	return &Node{
		key:    key,
		val:    val,
		edges:  edges,
		isLeaf: isLeaf,
	}
}

func setEdge(label byte, node *Node) *Edge {
	return &Edge{
		label: label,
		node:  node,
	}
}

func (n *Node) updateNode(key string, val interface{}) {
	if key != "" {
		n.key = key
	}
	if val != nil {
		n.val = val
	}
}

func (e *Edge) updateEdge(label byte, node *Node) {
	if label != 0 {
		e.label = label
	}
	if node != nil {
		e.node = node
	}
}

// FindPrefix find the prefix match between input key and root key.
// Return the index of the next to the last char of matched prefix
func (n *Node) FindPrefix(key string) (pos int) {
	max := len(key)
	if nodeKeyLen := len(n.key); nodeKeyLen < max {
		max = nodeKeyLen
	}
	for pos = 0; pos < max; pos++ {
		if key[pos] != n.key[pos] {
			break
		}
	}
	return pos
}

func findIncision(nodeKey, inputKey string, pos int) (matchedKey, nodeKeyDiff, inputKeyDiff string) {
	if pos > 0 {
		matchedKey = nodeKey[:pos]
		if len(nodeKey) == pos && len(inputKey) > pos { // all nodekey in inputkey
			inputKeyDiff = inputKey[pos:]
		} else if len(nodeKey) > pos && len(inputKey) == pos { // all inputkey in nodekey
			nodeKeyDiff = nodeKey[pos:]
		} else if len(nodeKey) > pos && len(inputKey) > pos { // both have diff
			nodeKeyDiff = nodeKey[pos:]
			inputKeyDiff = inputKey[pos:]
		}
	}
	return
}

// FindEdge find edge of the node based on label
func (n *Node) findEdge(label byte) *Edge {
	// this function implementation is from Armon Dadgar - go-radix
	num := len(n.edges)
	idx := sort.Search(num, func(i int) bool {
		return n.edges[i].label >= label
	})
	if idx < num && n.edges[idx].label == label {
		return n.edges[idx]
	}
	return nil
}

// addEdge add edge to a node
func (n *Node) addEdge(edges ...*Edge) {
	for _, edge := range edges {
		n.edges = append(n.edges, edge)
	}
	sort.Sort(n.edges)
}

// checkEdges check if a node has childs(true) or not(false)
func (n *Node) checkEdges() bool {
	if len(n.edges) > 0 {
		return true
	}
	return false
}

// Insert Insert key & val into the tree recursively
// params:
// - node 		: current destination for insertion
// - parentNode	: parent of current node
// - key  		: string
// - val  		: any value
func (t *Tree) insert(node, parentNode *Node, key string, val interface{}) {
	if node == nil {
		t.root = setNode(key, val, nil)
		return
	} else {
		pos := node.FindPrefix(key) // match input key with current node
		p := node                   // pointer for current node

		// if no matched found between node key and input key then it must be root
		// if root is not empty, then create new empty root and add edge to old and new node
		// if root is empty, continue insertion on the next matched edge
		if pos == 0 {
			if p == t.root {
				if p.key != "" {
					edgeP := setEdge(p.key[0], p)
					newNode := setNode(key, val, nil)
					newEdge := setEdge(key[0], newNode)
					t.root = setNode("", nil, nil)
					t.root.addEdge(edgeP, newEdge)
					return
				} else {
					e := p.findEdge(key[0])
					if e != nil {
						t.insert(e.node, p, key, val)
					} else {
						newNode := setNode(key, val, nil)
						newEdge := setEdge(newNode.key[0], newNode)
						p.addEdge(newEdge)
						return
					}
				}
			}
		}
		matchedKey, nodeKeyDiff, inputKeyDiff := findIncision(p.key, key, pos)

		// remainder on both node key and input key
		// IF Root: new root and its 2 edge to old root and to the new node from input
		// IF Not Root: new node and parent of current node point to new node and new node has 2 edge to new node and old node
		if nodeKeyDiff != "" && inputKeyDiff != "" {
			if p == t.root {
				p.key = nodeKeyDiff
				edgeP := setEdge(p.key[0], p)

				newNode := setNode(inputKeyDiff, val, nil)
				edgeN := setEdge(newNode.key[0], newNode)

				t.root = setNode(matchedKey, nil, nil)
				t.root.addEdge(edgeP, edgeN)
			} else {
				newMatchedNode := setNode(matchedKey, nil, nil)
				newEdgeN := setEdge(nodeKeyDiff[0], p)

				parentNode.findEdge(p.key[0]).updateEdge(newMatchedNode.key[0], newMatchedNode)

				p.updateNode(nodeKeyDiff, nil)

				newNode := setNode(inputKeyDiff, val, nil)
				newEdge := setEdge(newNode.key[0], newNode)

				newMatchedNode.addEdge(newEdgeN, newEdge)
			}
			return
		}

		// remainder on node key only
		// this will result in new root with only 1 edge to the old root
		if nodeKeyDiff != "" && inputKeyDiff == "" {
			if p == t.root {
				p.key = nodeKeyDiff
				newEdge := setEdge(p.key[0], p)
				t.root = setNode(matchedKey, val, nil)
				t.root.addEdge(newEdge)
			} else {
				newMatchedNode := setNode(matchedKey, val, nil)
				newEdgeN := setEdge(nodeKeyDiff[0], p)
				newMatchedNode.addEdge(newEdgeN)

				parentNode.findEdge(p.key[0]).updateEdge(newMatchedNode.key[0], newMatchedNode)

				p.updateNode(nodeKeyDiff, nil)
			}
			return
		}

		// remainder on input key only
		// this will result in checking for current root edge by first indice of input key, if not exist input into new node
		if nodeKeyDiff == "" && inputKeyDiff != "" {
			e := p.findEdge(inputKeyDiff[0])
			if e != nil {
				t.insert(e.node, p, inputKeyDiff, val)
			} else {
				newNode := setNode(inputKeyDiff, val, nil)
				newEdge := setEdge(newNode.key[0], newNode)
				p.addEdge(newEdge)
				return
			}
		}

		// if both remainder empty
		if nodeKeyDiff == "" && inputKeyDiff == "" {
			if p == t.root {
				t.root.val = val
			} else {
				p.val = val
			}
			return
		}
	}
}

func (t *Tree) InsertTree(key string, val interface{}) {
	t.insert(t.root, nil, key, val)
}

// LookUp find value by key and if exist also return true
func (t *Tree) lookUp(node *Node, key string) (interface{}, bool) {
	var (
		val     interface{}
		isExist bool
	)
	if node == nil {
		val, isExist = nil, false
	} else {
		pos := node.FindPrefix(key)
		p := node
		// if none match
		if pos == 0 {
			if p == t.root {
				if p.key != "" {
					val, isExist = nil, false
				} else {
					e := p.findEdge(key[0])
					if e != nil {
						val, isExist = t.lookUp(e.node, key)
					} else {
						val, isExist = nil, false
					}
				}
			}
		}
		_, nodeKeyDiff, inputKeyDiff := findIncision(p.key, key, pos)

		if nodeKeyDiff != "" && inputKeyDiff != "" {
			val, isExist = nil, false
		}

		if nodeKeyDiff != "" && inputKeyDiff == "" {
			val, isExist = nil, false
		}

		if nodeKeyDiff == "" && inputKeyDiff != "" {
			e := p.findEdge(inputKeyDiff[0])
			if e != nil {
				val, isExist = t.lookUp(e.node, inputKeyDiff)
			} else {
				val, isExist = nil, false
			}
		}

		// if both remainder empty
		if nodeKeyDiff == "" && inputKeyDiff == "" {
			if p.isLeaf {
				val, isExist = p.val, true
			}
		}
	}
	return val, isExist
}

func (t *Tree) LookUpTree(key string) interface{} {
	if val, isExist := t.lookUp(t.root, key); isExist {
		return val
	}
	return nil
}

func (t *Tree) getKeys(n *Node, key string, keys []string) []string {

	if n == nil {
		return nil
	} else {
		if n.isLeaf {
			key += n.key
			keys = append(keys, key)
		} else {
			key += n.key
		}
		if len(n.edges) > 0 {
			for _, e := range n.edges {
				keys = t.getKeys(e.node, key, keys)
			}
		}
	}
	return keys
}

func (t *Tree) getValues(n *Node, values []interface{}) []interface{} {

	if n == nil {
		return nil
	} else {
		if n.isLeaf {
			values = append(values, n.val)
		}
		if len(n.edges) > 0 {
			for _, e := range n.edges {
				values = t.getValues(e.node, values)
			}
		}
	}
	return values
}

func (t *Tree) getKeysValues(n *Node, key string, kv map[string]interface{}) map[string]interface{} {
	if n == nil {
		return nil
	} else {
		if n.isLeaf {
			key += n.key
			kv[key] = n.val
		} else {
			key += n.key
		}
		if len(n.edges) > 0 {
			for _, e := range n.edges {
				kv = t.getKeysValues(e.node, key, kv)
			}
		}
	}
	return kv
}

func (t *Tree) GetAllKeys() []string {
	var key string
	var keys []string
	return t.getKeys(t.root, key, keys)
}

func (t *Tree) GetAllValues() []interface{} {
	var vals []interface{}
	return t.getValues(t.root, vals)
}

func (t *Tree) GetAllKeysValues() map[string]interface{} {
	var key string
	m := make(map[string]interface{})
	return t.getKeysValues(t.root, key, m)
}
