package set1

import (
	"fmt"
	"testing"
)

func TestDetectSingleXORedChar(t *testing.T) {
	hex, decreptedHex, key, err := DetectSingleXORedCharHex("4.txt")
	if err != nil {
		t.Errorf("err: %s", err.Error())
	}

	fmt.Printf("%s\n %s (key: %c)\n", hex, decreptedHex, key)
}
