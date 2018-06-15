package merkletree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTree(t *testing.T) {
	assert.Equal(t, NewTree(), &Tree{nil})
}

func TestGetRoot(t *testing.T) {
	tree := NewTree()
	assert.Nil(t, tree.Root())
}

func TestAddNodeToTree(t *testing.T) {
	// when tree has no nodes
	tree := NewTree()
	node := NewLeaf("0")
	tree.AddNode(node)
	assert.NotNil(t, node)
	assert.Nil(t, tree.Root().Right)
	assert.Equal(t, node, tree.Root().Left)
}
