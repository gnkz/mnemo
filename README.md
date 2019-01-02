# mnemo

A [bip-39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) implementation in Go.

# Documentation

You can get the documentation [here](https://godoc.org/github.com/gnkz/mnemo).

# Usage

```go
package main

import(
	"fmt"
	
	"github.com/gnkz/mnemo"
	"github.com/gnkz/mnemo/en"
)

func main() {
    // Generate a 128 bit entropy
	ent, err := mnemo.Entropy(mnemo.Entropy128)
	if err != nil {
		panic(err)	
	}

    // Generate the mnemonic using the entropy and the english set of words
	mnemonic, err := mnemo.NewFromEntropy(ent, en.New())
	if err != nil {
		panic(err)
	}

	fmt.Println(mnemonic)

    // Generate a master seed using the mnemonic and a passphrase
	seed := mnemo.Seed(mnemonic, "sup3rs3cr3t")

	fmt.Printf("%x\n", seed) 
}
```
# License

Copyright 2019 gnkz

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
