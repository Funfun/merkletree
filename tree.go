package merkletree

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

// Root returns root node
func (t Tree) Root() *Node {
	return t.root
}

// AddNode adds node to Tree
func (t *Tree) AddNode(node *Node) {
	if t.root == nil {
		t.root = NewNode(HashNode(node.Hash[:]), node, nil)
	}
}
