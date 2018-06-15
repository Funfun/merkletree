package merkletree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	assert.Exactly(t,
		NewNode(Hash("hashed-data"), nil, nil),
		&Node{Hash("hashed-data"), nil, nil, false})
}

func TestDepthFirstSearch(t *testing.T) {
	h := NewNode(Hash("H"), nil, nil)
	a := NewNode(Hash("A"), nil, nil)
	e := NewNode(Hash("E"), nil, nil)
	c := NewNode(Hash("C"), nil, nil)
	i := NewNode(Hash("I"), h, nil)
	g := NewNode(Hash("G"), nil, i)
	d := NewNode(Hash("D"), c, e)
	b := NewNode(Hash("B"), a, d)
	f := NewNode(Hash("F"), b, g)

	actualHashes := []*Node{}
	expectedHashes := []*Node{a, c, e, d, b, h, i, g, f}
	DepthFirstSearch(f, &actualHashes)

	assert.Equal(t, expectedHashes, actualHashes)
}
