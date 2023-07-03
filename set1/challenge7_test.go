package set1

import (
	"fmt"
	"testing"
)

func TestDecryptAESECB(t *testing.T) {
	out, err := DecryptAESECB("7.txt", []byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Fatalf("err: %s\n", err.Error())
	}

	fmt.Println(string(out))
}
