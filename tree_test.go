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
