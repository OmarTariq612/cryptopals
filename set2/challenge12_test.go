package set2

import (
	"fmt"
	"testing"
)

func TestECBDecryption(t *testing.T) {
	unKownBytes, err := ECBDecryption(NewECBEncrypter())
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	fmt.Println(len(unKownBytes))
	fmt.Println(unKownBytes)
	fmt.Println(string(unKownBytes))
}
