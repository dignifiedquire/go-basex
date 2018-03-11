package basex_test

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

	if bs58.Encode(decoded) != "5Kd3NBUAdUnhyzenEwVLy9pBKxSwXvE9FMPyR4UKZvpe6E3AgLr" {
		t.Fatal("unexpected roundtrip result")
	}
}
