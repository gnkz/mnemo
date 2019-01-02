package en_test

import (
	"encoding/hex"
	"testing"

	"github.com/gnkz/mnemo"
	"github.com/gnkz/mnemo/en"
)

func TestMnemonic(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{
			input:    "00000000000000000000000000000000",
			expected: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
			err:      nil,
		},
		{
			input:    "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			expected: "legal winner thank year wave sausage worth useful legal winner thank yellow",
			err:      nil,
		},
		{
			input:    "80808080808080808080808080808080",
			expected: "letter advice cage absurd amount doctor acoustic avoid letter advice cage above",
			err:      nil,
		},
		{
			input:    "ffffffffffffffffffffffffffffffff",
			expected: "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo wrong",
			err:      nil,
		},
		{
			input:    "9e885d952ad362caeb4efe34a8e91bd2",
			expected: "ozone drill grab fiber curtain grace pudding thank cruise elder eight picnic",
			err:      nil,
		},
		{
			input:    "000000000000000000000000000000000000000000000000",
			expected: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon agent",
			err:      nil,
		},
		{
			input:    "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			expected: "legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal will",
			err:      nil,
		},
		{
			input:    "808080808080808080808080808080808080808080808080",
			expected: "letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter always",
			err:      nil,
		},
		{
			input:    "ffffffffffffffffffffffffffffffffffffffffffffffff",
			expected: "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo when",
			err:      nil,
		},
	}

	for _, cc := range cases {
		input, _ := hex.DecodeString(cc.input)
		result, err := mnemo.New(input, en.New())

		if err != cc.err {
			t.Fatalf("Expected error %v but got %v\n", cc.err, err)
		}

		if result != cc.expected {
			t.Fatalf("Expected result %s but got %s\n", cc.expected, result)
		}
	}
}
