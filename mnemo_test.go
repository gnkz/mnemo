package mnemo_test

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/gnkz/mnemo"
	"github.com/gnkz/mnemo/en"
	mock_mnemo "github.com/gnkz/mnemo/mocks"
	"github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
	cases := []struct {
		length   mnemo.MnemonicLength
		expected string
		err      error
	}{
		{
			length: mnemo.Words12,
			err:    nil,
		},
		{
			length: mnemo.Words15,
			err:    nil,
		},
		{
			length: mnemo.Words18,
			err:    nil,
		},
		{
			length: mnemo.Words21,
			err:    nil,
		},
		{
			length: mnemo.Words24,
			err:    nil,
		},
	}

	for _, cc := range cases {
		ctrl := gomock.NewController(t)
		dict := mock_mnemo.NewMockDictionary(ctrl)
		dict.EXPECT().Word(gomock.Any()).Return("test", nil).AnyTimes()
		dict.EXPECT().Separator().Return(" ")

		mnemonic, err := mnemo.New(cc.length, dict)

		if err != cc.err {
			t.Fatalf("Expected error %v but got %v\n", cc.err, err)
		}

		words := strings.Split(mnemonic, " ")

		if len(words) != int(cc.length) {
			t.Errorf("Expected mnemonic length of %d but got %d\n", cc.length, len(words))
		}
	}
}
func TestNewFromEntropy(t *testing.T) {
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
		result, err := mnemo.NewFromEntropy(input, dict)

		if err != cc.err {
			t.Fatalf("Expected error %v but got %v\n", cc.err, err)
		}

		if result != cc.expected {
			t.Fatalf("Expected result %s but got %s\n", cc.expected, result)
		}
	}
}

func ExampleNew() {
	dict := en.New()

	mnemonic, _ := mnemo.New(mnemo.Words12, dict)

	words := strings.Split(mnemonic, dict.Separator())

	fmt.Println(len(words))
	// Output: 12
}

func ExampleNewFromEntropy() {
	entropy, _ := hex.DecodeString("9e885d952ad362caeb4efe34a8e91bd2")
	mnemonic, _ := mnemo.NewFromEntropy(entropy, en.New())
	fmt.Println(mnemonic)
	// Output: ozone drill grab fiber curtain grace pudding thank cruise elder eight picnic
}
