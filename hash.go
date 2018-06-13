package merkletree

import (
	"crypto/sha256"
)

// Hash returns SHA256 checksum of the data
func Hash(data string) [32]byte {
	return sha256.Sum256([]byte(data))
}
