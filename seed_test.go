package mnemo_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/gnkz/mnemo"
)

func TestSeed(t *testing.T) {
	cases := []struct {
		mnemonic   string
		passphrase string
		expected   string
	}{
		{
			mnemonic:   "gravity machine north sort system female filter attitude volume fold club stay feature office ecology stable narrow fog",
			passphrase: "TREZOR",
			expected:   "628c3827a8823298ee685db84f55caa34b5cc195a778e52d45f59bcf75aba68e4d7590e101dc414bc1bbd5737666fbbef35d1f1903953b66624f910feef245ac",
		},
	}

	for _, cc := range cases {
		got := mnemo.Seed(cc.mnemonic, cc.passphrase)

		if hex.EncodeToString(got) != cc.expected {
			t.Errorf("Expected seed %s but got %x\n", cc.expected, got)
		}
	}
}

func ExampleSeed() {
	mnemonic := "gravity machine north sort system female filter attitude volume fold club stay feature office ecology stable narrow fog"
	passphrase := "TREZOR"

	seed := mnemo.Seed(mnemonic, passphrase)

	fmt.Printf("%x", seed)
	// Output: 628c3827a8823298ee685db84f55caa34b5cc195a778e52d45f59bcf75aba68e4d7590e101dc414bc1bbd5737666fbbef35d1f1903953b66624f910feef245ac
}
