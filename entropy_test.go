package mnemo_test

import (
	"fmt"
	"testing"

	"github.com/gnkz/mnemo"
)

func TestEntropy(t *testing.T) {
	cases := []struct {
		input     mnemo.EntropyLength
		outputLen int
		err       error
	}{
		{
			input:     mnemo.Entropy128,
			outputLen: 16,
			err:       nil,
		},
		{
			input:     mnemo.Entropy160,
			outputLen: 20,
			err:       nil,
		},
		{
			input:     mnemo.Entropy192,
			outputLen: 24,
			err:       nil,
		},
		{
			input:     mnemo.Entropy224,
			outputLen: 28,
			err:       nil,
		},
		{
			input:     mnemo.Entropy256,
			outputLen: 32,
			err:       nil,
		},
	}

	for _, cc := range cases {
		en, err := mnemo.Entropy(cc.input)

		if err != cc.err {
			t.Fatalf("Expected error %v but got %v\n", cc.err, err)
		}

		if len(en) != cc.outputLen {
			t.Errorf("Expected length %d but got %d\n", cc.outputLen, len(en))
		}

	}
}

func ExampleEntropy() {
	en, _ := mnemo.Entropy(mnemo.Entropy128)

	fmt.Println(len(en))
	// Output: 16
}
