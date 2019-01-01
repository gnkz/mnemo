package mnemo

import (
	"crypto/sha256"
	"errors"
)

// Checksum returns the first N bits of the sha256 hash of the entropy.
//
// The entropy must be of 128, 160, 192, 224 or 256 bits.
//
// 128 bits entropy generates a 4 bits checksum
//
// 160 bits entropy generates a 5 bits checksum
//
// 192 bits entropy generates a 6 bits checksum
//
// 224 bits entropy generates a 7 bits checksum
//
// 256 bits entropy generates a 8 bits checksum
func Checksum(entropy []byte) (byte, uint8, error) {
	enLen := EntropyLength(len(entropy) * 8)

	var (
		csMask byte
		csLen  uint8
	)

	switch enLen {
	case Entropy128:
		csMask = byte(0xf0)
		csLen = 4
	case Entropy160:
		csMask = byte(0xf8)
		csLen = 5
	case Entropy192:
		csMask = byte(0xfc)
		csLen = 6
	case Entropy224:
		csMask = byte(0xfe)
		csLen = 7
	case Entropy256:
		csMask = byte(0xff)
		csLen = 8
	default:
		return 0, 0, errors.New("Invalid entropy length")
	}

	sum := sha256.Sum256(entropy)
	cs := (sum[0] & csMask) >> (8 - csLen)

	return cs, csLen, nil
}
