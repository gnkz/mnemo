package mnemo

import (
	"crypto/sha512"

	"golang.org/x/crypto/pbkdf2"
)

const (
	hashIterations int    = 2048
	seedLength     int    = 64
	saltPrefix     string = "mnemonic"
)

// Seed creates a master seed from a rawMnemonic and a passphrase
func Seed(rawMenmonic, passphrase string) []byte {
	salt := []byte(saltPrefix)

	if passphrase != "" {
		salt = append(salt, []byte(passphrase)...)
	}

	return pbkdf2.Key([]byte(rawMenmonic), salt, hashIterations, seedLength, sha512.New)
}
