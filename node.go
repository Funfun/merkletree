package merkletree

import (
	"fmt"
)

type Node struct {
	Hash      string
	LeftNode  *Node
	RightNode *Node
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

	DepthFirstSearch(node.LeftNode, result)
	DepthFirstSearch(node.RightNode, result)

	*result = append(*result, node)
}
