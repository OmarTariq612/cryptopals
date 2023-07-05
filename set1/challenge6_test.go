package set1

import (
	"fmt"
	"testing"
)

func TestDecryptWithRepeatingXORKey(t *testing.T) {
	key, err := DecryptWithRepeatingXORKey("6.txt")
	if err != nil {
		t.Fatalf("err: %s", err.Error())
	}

	fmt.Println(key)
	fmt.Println(len(key))
}
