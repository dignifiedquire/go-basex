# base-x

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](https://protocol.ai)
[![GoDoc](https://godoc.org/github.com/dignifiedquire/go-basex?status.svg)](https://godoc.org/github.com/dignifiedquire/go-basex)
[![Coverage Status](https://coveralls.io/repos/github/dignifiedquire/go-basex/badge.svg?branch=master)](https://coveralls.io/github/dignifiedquire/go-basex?branch=master)
[![Build Status](https://travis-ci.org/dignifiedquire/go-basex.svg?branch=master)](https://travis-ci.org/dignifiedquire/go-basex)

> Fast base encoding / decoding of any given alphabet using bitcoin style leading zero compression.

## Example

Base58

``` go

import (
	"fmt"
	"testing"

	"github.com/dignifiedquire/go-basex"
)

func TestExample(t *testing.T) {
	Base58Charset := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	bs58 := basex.NewAlphabet(Base58Charset)

	decoded, err := bs58.Decode("5Kd3NBUAdUnhyzenEwVLy9pBKxSwXvE9FMPyR4UKZvpe6E3AgLr")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", decoded)
	// => 80eddbdc1168f1daeadbd3e44c1e3f8f5a284c2029f78ad26af98583a499de5b1913a4f863

	fmt.Printf("%s\n", bs58.Encode(decoded))
	// => 5Kd3NBUAdUnhyzenEwVLy9pBKxSwXvE9FMPyR4UKZvpe6E3AgLr
```

### Alphabets

See below for a list of commonly recognized alphabets, and their respective base.

Base | Alphabet
------------- | -------------
2 | `01`
8 | `01234567`
11 | `0123456789a`
16 | `0123456789abcdef`
32 | `0123456789ABCDEFGHJKMNPQRSTVWXYZ`
36 | `0123456789abcdefghijklmnopqrstuvwxyz`
58 | `123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz`
62 | `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
64 | `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`
66 | `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.!~`


## How it works

It encodes octet arrays by doing long divisions on all significant digits in the
array, creating a representation of that number in the new base. Then for every
leading zero in the input (not significant as a number) it will encode as a
single leader character. This is the first in the alphabet and will decode as 8
bits. The other characters depend upon the base. For example, a base58 alphabet
packs roughly 5.858 bits per character.

This means the encoded string 000f (using a base16, 0-f alphabet) will actually decode
to 4 bytes unlike a canonical hex encoding which uniformly packs 4 bits into each
character.

While unusual, this does mean that no padding is required and it works for bases
like 43. **If you need standard hex encoding, or base64 encoding, this module is NOT
appropriate.**

## Credit

This is a reimplementation of the great package [cryptcoinjs/base-x](https://github.com/cryptocoinjs/base-x).

## LICENSE

[MIT](LICENSE)
