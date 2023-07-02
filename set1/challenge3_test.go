package set1

import (
	"fmt"
	"testing"
)

func TestSingleCharXORed(t *testing.T) {
	out, key, err := SingleCharXORed("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		t.Fatalf("err: %s", err.Error())
	}
	fmt.Printf("%s (key: %c)\n", out, key)
}
