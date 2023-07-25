package set2

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestDecryptCBCMode(t *testing.T) {
	ciphertext, err := os.ReadFile("10.txt")
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	n, err := base64.StdEncoding.Decode(ciphertext, ciphertext)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	ciphertext = ciphertext[:n]

	blockCipher, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	iv := make([]byte, 16)

	plaintext, err := DecryptCBCMode(blockCipher, iv, ciphertext)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	fmt.Println(string(plaintext))
}
