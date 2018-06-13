package merkletree

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	assert.Equal(t, Hash("189"), sha256.Sum256([]byte("189")))
}
