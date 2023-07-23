package set2

import (
	"fmt"
	"testing"
)

func TestPKCS7Unpad(t *testing.T) {
	res, err := PKCS7Unpad([]byte("ICE ICE BABY\x04\x04\x04\x04"), 16)
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}
	fmt.Println(string(res))
	res, err = PKCS7Unpad([]byte("ICE ICE BABY\x05\x05\x05\x05"), 16)
	if err == nil {
		t.Fatalf("considered valid even it's invalid")
	}
	fmt.Println(res)
	res, err = PKCS7Unpad([]byte("ICE ICE BABY\x01\x02\x03\x04"), 16)
	if err == nil {
		t.Fatalf("considered valid even it's invalid")
	}
	fmt.Println(res)
}
