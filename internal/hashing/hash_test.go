package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashStringNotEquals(t *testing.T) {
	message1 := "super secret message."
	message2 := "even more secret message."

	hash1, _ := HashString(message1)
	hash2, _ := HashString(message2)

	assert.NotEqual(t, hash1, hash2)
}

func TestHashStringEquals(t *testing.T) {
	message1 := "super secret message."
	message2 := "super secret message."

	hash1, _ := HashString(message1)
	hash2, _ := HashString(message2)

	assert.Equal(t, hash1, hash2)
}

func TestHashStringCorrectLength(t *testing.T) {
	message1 := "super secret message."

	hash, _ := HashString(message1)

	assert.Equal(t, 64, len(hash))
}
