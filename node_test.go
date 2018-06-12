package merkletree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	assert.Exactly(t, NewNode("hashed-data", nil, nil), &Node{"hashed-data", nil, nil})
}

func TestDepthFirstSearch(t *testing.T) {
	h := NewNode("H", nil, nil)
	a := NewNode("A", nil, nil)
	e := NewNode("E", nil, nil)
	c := NewNode("C", nil, nil)
	i := NewNode("I", h, nil)
	g := NewNode("G", nil, i)
	d := NewNode("D", c, e)
	b := NewNode("B", a, d)
	f := NewNode("F", b, g)

	actualHashes := []*Node{}
	expectedHashes := []*Node{a, c, e, d, b, h, i, g, f}
	DepthFirstSearch(f, &actualHashes)

	assert.Equal(t, expectedHashes, actualHashes)
}
