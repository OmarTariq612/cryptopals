package set1

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	out, err := HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		t.Fatalf("err: %s", err.Error())
	}
	if out != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Fatalf("output is not correct")
	}
}
