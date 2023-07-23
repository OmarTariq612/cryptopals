package set2

import (
	"fmt"
	"testing"
)

func TestProfileFor(t *testing.T) {
	if _, err := ProfileFor("foo@bar.com&f"); err == nil {
		t.Fatal("err must not be nil (contains &)")
	}

	if _, err := ProfileFor("foo@bar.com=f"); err == nil {
		t.Fatal("err must not be nil (contains =)")
	}
}

func TestMakeNewAdminAccount(t *testing.T) {
	block, err := NewAES128Cipher()
	if err != nil {
		t.Fatalf("err: %s\n", err.Error())
	}

	ciphertext, err := MakeNewAdminAccount(NewECBEncrypterFromBlockCipher(block))
	if err != nil {
		t.Fatalf("err: %s\n", err.Error())
	}

	plaintext, err := DecryptECBMode(block, ciphertext)
	if err != nil {
		t.Fatalf("err: %s\n", err.Error())
	}

	elements, err := ParseElements(string(plaintext))
	if err != nil {
		t.Fatalf("err: %s\n", err.Error())
	}

	fmt.Println(elements)

	if elements[2].Key != "role" || elements[2].Value != "admin" {
		t.Fatal("err: test didn't pass")
	}
}
