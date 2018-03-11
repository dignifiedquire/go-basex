// Package basex is a port of [base-x](https://github.com/cryptocoinjs/base-x) to Go.
package basex

import (
	"fmt"
	"strings"
)

// Alphabet holds an arbitrary base alphabet.
type Alphabet struct {
	raw string
	dec map[string]byte
	enc []string
}

func (a *Alphabet) String() string {
	return a.raw
}

// Base returns the base size of this alphabet.
func (a *Alphabet) Base() uint {
	return uint(len(a.enc))
}

// NewAlphabet creates a new alphabet by generating the necessariy lookup tables.
func NewAlphabet(raw string) *Alphabet {
	parts := strings.Split(raw, "")

	alph := &Alphabet{
		raw: raw,
		dec: map[string]byte{},
		enc: make([]string, len(parts)),
	}

	for i, a := range parts {
		alph.dec[a] = byte(i)
		alph.enc[i] = a
	}

	return alph
}

// EncodeToBytes encodes the given input only using the base and returns
// a byte slice with entries in the range from 0 to base.
func (a *Alphabet) EncodeToBytes(input []byte) []byte {
	if len(input) == 0 {
		return []byte{}
	}

	base := uint32(len(a.enc))

	digits := []byte{0}
	for i := 0; i < len(input); i++ {
		carry := uint32(input[i])
		for j := 0; j < len(digits); j++ {
			carry += uint32(digits[j]) << 8
			digits[j] = byte(carry % base)
			carry = carry / base
		}

		for carry > 0 {
			digits = append(digits, byte(carry%base))
			carry = carry / base
		}
	}

	var out []byte

	// deal with leading zeros
	for k := 0; input[k] == 0 && k < len(input)-1; k++ {
		out = append(out, 0)
	}

	return reverse(append(digits, out...))
}

// Encode encodes the given input using the current alphabet.
func (a *Alphabet) Encode(input []byte) string {
	alph := a.enc
	digits := a.EncodeToBytes(input)

	// convert digits to a string
	out := ""
	for _, digit := range digits {
		out += alph[digit]
	}
	return out
}

// DecodeFromBytes decodes a byte slice in the range of 0 to base, into the
// original data.
func (a *Alphabet) DecodeFromBytes(input []byte) ([]byte, error) {
	base := uint32(len(a.enc))

	bytes := []byte{0}
	for i := 0; i < len(input); i++ {
		value := input[i]

		carry := uint32(value)
		for j := 0; j < len(bytes); j++ {
			carry += uint32(bytes[j]) * base
			bytes[j] = byte(carry & 0xff)
			carry >>= 8
		}

		for carry > 0 {
			bytes = append(bytes, byte(carry&0xff))
			carry >>= 8
		}
	}

	// deal with leading zeros
	for k := 0; input[k] == 0 && k < len(input)-1; k++ {
		bytes = append(bytes, 0)
	}

	return reverse(bytes), nil
}

// Decode decodes the given string into the original bytes.
func (a *Alphabet) Decode(input string) ([]byte, error) {
	if len(input) == 0 {
		return []byte{}, nil
	}

	alphMap := a.dec

	values := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		value, ok := alphMap[string(input[i])]
		if !ok {
			return nil, fmt.Errorf("non base character: %s", string(input[i]))
		}

		values[i] = value
	}

	return a.DecodeFromBytes(values)
}

func reverse(numbers []byte) []byte {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
