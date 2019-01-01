// Package mnemo is an implementation of the Bip-39 mnemonic code generation scheme.
//
// A mnemonic code is a set of words that are used to generate seeds to create HD wallets.
package mnemo

import (
	"math/big"
	"strings"
)

// New creates a new mnemonic using a valid entropy and a Dictionary
func New(entropy []byte, dict Dictionary) (string, error) {
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
