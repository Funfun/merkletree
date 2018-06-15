package merkletree

import (
	"crypto/sha256"
	"fmt"
)

const Size = sha256.Size

type Node struct {
	Hash        [Size]byte
	Left, Right *Node
	leaf        bool
}

func NewNode(hash [Size]byte, left *Node, right *Node) *Node {
	return &Node{hash, left, right, false}
}

func NewLeaf(data string) *Node {
	hash := Hash(data)
	return &Node{hash, nil, nil, true}
}

func (n Node) String() string {
	return fmt.Sprintf("Node => Hash: %x", n.Hash)
}

// Post-order (LRN) search
func DepthFirstSearch(node *Node, result *[]*Node) {
	if node == nil {
		return
	}

	DepthFirstSearch(node.Left, result)
	DepthFirstSearch(node.Right, result)

	*result = append(*result, node)
}

func (n Node) Leaf() bool {
	return n.leaf
}

func (n *Node) SetLeaf() {
	n.leaf = true
}
