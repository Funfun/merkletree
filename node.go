package merkletree

import (
	"fmt"
)

type Node struct {
	Hash        string
	Left, Right *Node
}

func NewNode(hash string, left *Node, right *Node) *Node {
	return &Node{hash, left, right}
}

func (n Node) String() string {
	return fmt.Sprintf("Node => Hash: %s", n.Hash)
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
