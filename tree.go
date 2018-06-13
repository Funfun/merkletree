package merkletree

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}
