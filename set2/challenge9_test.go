package set2

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPKCS7Pad(t *testing.T) {
	output := PKCS7Pad([]byte("YELLOW SUBMARINE"), 20)
	if !bytes.Equal(output, []byte("YELLOW SUBMARINE\x04\x04\x04\x04")) {
		fmt.Println(len(output))
		fmt.Println(output)
		t.Fatal("err: invalid padding")
	}

	output = PKCS7Pad([]byte("hellohello"), 5)
	if !bytes.Equal(output, []byte("hellohello\x05\x05\x05\x05\x05")) {
		fmt.Println(len(output))
		fmt.Println(output)
		t.Fatal("err: invalid padding")
	}
}
