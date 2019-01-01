package mnemo

import (
	"crypto/rand"
	"errors"
)

type (
	// EntropyLength defines a set of defined entropy length in bits
	EntropyLength int
)

const (
	// Entropy128 represents a 128 bits length entropy
	Entropy128 EntropyLength = 32 * (4 + iota)

	// Entropy160 represents a 160 bits length entropy
	Entropy160

	// Entropy192 represents a 192 bits length entropy
	Entropy192

	// Entropy224 represents a 224 bits length entropy
	Entropy224

	// Entropy256 represents a 256 bits length entropy
	Entropy256
)

// Entropy returns an entropy of length bits. If the length is different from 128, 160, 192, 224 or 256 it will return an error
func Entropy(length EntropyLength) ([]byte, error) {
	switch length {
	case
		Entropy128,
		Entropy160,
		Entropy192,
		Entropy224,
		Entropy256:
	default:
		return nil, errors.New("Invalid entropy length")
	}

	entropy := make([]byte, length/8)

	_, err := rand.Read(entropy)
	if err != nil {
		return nil, err
	}

	return entropy, nil
}
