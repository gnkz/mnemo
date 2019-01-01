package mnemo_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/gnkz/mnemo"
)

func TestChecksum(t *testing.T) {
	cases := []struct {
		input    string
		checksum byte
		length   uint8
		err      error
	}{
		{
			input:    "00000000000000000000000000000000",
			checksum: 0x03,
			length:   4,
			err:      nil,
		},
		{
			input:    "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			checksum: 0x08,
			length:   4,
			err:      nil,
		},
		{
			input:    "000000000000000000000000000000000000000000000000",
			checksum: 0x27,
			length:   6,
			err:      nil,
		},
		{
			input:    "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			checksum: 0x19,
			length:   6,
			err:      nil,
		},
	}

	for _, cc := range cases {
		input, _ := hex.DecodeString(cc.input)
		cs, l, err := mnemo.Checksum(input)

		if err != cc.err {
			t.Fatalf("Expected error %v but got %v\n", cc.err, err)
		}

		if cs != cc.checksum {
			t.Errorf("Expected checksum %x but got %x\n", cc.checksum, cs)
		}

		if l != cc.length {
			t.Errorf("Expected checksum length of %d but got %d\n", cc.length, l)
		}
	}
}

func ExampleChecksum() {
	entropy, _ := hex.DecodeString("9e885d952ad362caeb4efe34a8e91bd2")
	cs, cslen, _ := mnemo.Checksum(entropy)

	fmt.Println(cs)
	fmt.Println(cslen)
	// Output: 1
	// 4
}
