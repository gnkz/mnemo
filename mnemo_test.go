package mnemo_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/gnkz/mnemo"
	"github.com/gnkz/mnemo/en"
	"github.com/gnkz/mnemo/mocks"
	"github.com/golang/mock/gomock"
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
			expected: "test test test test test test test test test test test test",
			err:      nil,
		},
		{
			input:    "0000000000000000000000000000000000000000",
			expected: "test test test test test test test test test test test test test test test",
			err:      nil,
		},
		{
			input:    "000000000000000000000000000000000000000000000000",
			expected: "test test test test test test test test test test test test test test test test test test",
			err:      nil,
		},
	}

	for _, cc := range cases {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		dict := mock_mnemo.NewMockDictionary(ctrl)

		dict.EXPECT().Word(gomock.Any()).Return("test", nil).AnyTimes()
		dict.EXPECT().Separator().Return(" ")

		input, _ := hex.DecodeString(cc.input)
		result, err := mnemo.New(input, dict)

		if err != cc.err {
			t.Fatalf("Expected error %v but got %v\n", cc.err, err)
		}

		if result != cc.expected {
			t.Fatalf("Expected result %s but got %s\n", cc.expected, result)
		}
	}
}

func ExampleNew() {
	entropy, _ := hex.DecodeString("9e885d952ad362caeb4efe34a8e91bd2")
	mnemonic, _ := mnemo.New(entropy, en.New())
	fmt.Println(mnemonic)
	// Output: ozone drill grab fiber curtain grace pudding thank cruise elder eight picnic
}
