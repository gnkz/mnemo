// Package mnemo is an implementation of the Bip-39 mnemonic code generation scheme.
// A mnemonic code is a set of words that are used to generate seeds to create HD wallets.
package mnemo

import (
	"errors"
	"math/big"
	"strings"
)

// MnemonicLength represents the number of words to include
// in the mnemonic
type MnemonicLength int

const (
	// Words12 represents a 12 words mnemonic
	Words12 MnemonicLength = 12 + (3 * iota)

	// Words15 represents a 15 words mnemonic
	Words15

	// Words18 represents a 18 words mnemonic
	Words18

	// Words21 represents a 21 words mnemonic
	Words21

	// Words24 represents a 24 words mnemonic
	Words24
)

// New generates a random mnemonic of a determined length using words from
// a Dictionary
func New(length MnemonicLength, dict Dictionary) (string, error) {
	var entLen EntropyLength

	switch length {
	case Words12:
		entLen = Entropy128
	case Words15:
		entLen = Entropy160
	case Words18:
		entLen = Entropy192
	case Words21:
		entLen = Entropy224
	case Words24:
		entLen = Entropy256
	default:
		return "", errors.New("Invalid mnemonic length")
	}

	ent, err := Entropy(entLen)
	if err != nil {
		return "", err
	}

	return NewFromEntropy(ent, dict)
}

// NewFromEntropy creates a new mnemonic using a valid entropy and a Dictionary
func NewFromEntropy(entropy []byte, dict Dictionary) (string, error) {
	cs, csLen, err := Checksum(entropy)
	if err != nil {
		return "", err
	}

	checksumed := new(big.Int).SetBytes(entropy)
	checksumed.Lsh(checksumed, uint(csLen))
	checksumed.Or(checksumed, new(big.Int).SetBytes([]byte{cs}))

	enLen := len(entropy) * 8

	shiftRounds := (enLen + int(csLen)) / 11
	blockMask := big.NewInt(2047)
	blockMask.Lsh(blockMask, 11*uint(shiftRounds-1))
	words := []string{}

	for i := 0; i < shiftRounds; i++ {
		block := new(big.Int)
		block.And(checksumed, blockMask)
		block.Rsh(block, 11*uint(shiftRounds-(i+1)))

		word, err := dict.Word(int(block.Uint64()))
		if err != nil {
			return "", err
		}

		words = append(words, word)
		blockMask.Rsh(blockMask, 11)
	}

	return strings.Join(words, dict.Separator()), nil
}
