package set2

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestECBDecryption(t *testing.T) {
	unknownString := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	unknownBytes, err := base64.StdEncoding.DecodeString(unknownString)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	unkownBytesResult, err := ECBDecryption(NewECBSimpleEncrypter(unknownBytes))
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	if !bytes.Equal(unkownBytesResult, unknownBytes) {
		t.Fatalf("bytes are not the same")
	}
}
