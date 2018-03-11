package basex

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	Alphabet    string `json:"alphabet"`
	Hex         string `json:"hex"`
	Comment     string `json:"comment"`
	String      string `json:"string"`
	Exception   string `json:"exception"`
	Description string `json:"description"`
}

type Fixtures struct {
	Alphabets map[string]string `json:"alphabets"`
	Valid     []*Case           `json:"valid"`
	Invalid   []*Case           `json:"invalid"`
}

var fixtures Fixtures

func init() {
	data, err := ioutil.ReadFile("fixtures/fixtures.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &fixtures)
	if err != nil {
		panic(err)
	}
}

func TestValid(t *testing.T) {
	alphabets := fixtures.Alphabets

	for _, tc := range fixtures.Valid {
		alphabet := NewAlphabet(alphabets[tc.Alphabet])
		bytes, err := hex.DecodeString(tc.Hex)
		assert.NoError(t, err)

		t.Run(fmt.Sprintf("encode %s: %s", tc.Alphabet, tc.Hex), func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(alphabet.Encode(bytes), tc.String)
		})

		t.Run(fmt.Sprintf("decode %s: %s", tc.Alphabet, tc.Hex), func(t *testing.T) {
			assert := assert.New(t)

			decoded, err := alphabet.Decode(tc.String)
			assert.NoError(err)
			assert.Equal(decoded, bytes)
		})

		t.Run(fmt.Sprintf("roundtrip %s: %s", tc.Alphabet, tc.Hex), func(t *testing.T) {
			assert := assert.New(t)

			decoded, err := alphabet.Decode(tc.String)
			assert.NoError(err)
			encoded := alphabet.Encode(decoded)
			assert.Equal(encoded, tc.String)
		})
	}
}

func TestInvalid(t *testing.T) {
	alphabets := fixtures.Alphabets

	for _, tc := range fixtures.Invalid {
		alphabet := NewAlphabet(alphabets[tc.Alphabet])

		t.Run(fmt.Sprintf("decode fails: %s", tc.Description), func(t *testing.T) {
			assert := assert.New(t)

			decoded, err := alphabet.Decode(tc.String)
			assert.Error(err)
			assert.Empty(decoded)
		})
	}
}

func TestBase32(t *testing.T) {
	assert := assert.New(t)

	base32 := NewAlphabet("qpzry9x8gf2tvdw0s3jn54khce6mua7l")

	bytes, err := base32.Decode("andtheexcludedcharacters")
	assert.NoError(err)
	assert.Equal(base32.Encode(bytes), "andtheexcludedcharacters")

	input := []byte{1, 2, 3, 4, 8, 92, 255}
	enc := base32.EncodeToBytes(input)

	back, err := base32.DecodeFromBytes(enc)
	assert.NoError(err)
	assert.Equal(back, input)
}

func TestReverse(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		[]byte{1, 2, 3},
		reverse([]byte{3, 2, 1}),
	)
}

func TestAlphabet(t *testing.T) {
	assert := assert.New(t)

	a := NewAlphabet("abcdefghjklmnopqrstuvwxyz")
	assert.Equal(a.String(), "abcdefghjklmnopqrstuvwxyz")
	assert.Equal(a.Base(), uint(25))
}
