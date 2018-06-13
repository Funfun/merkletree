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
