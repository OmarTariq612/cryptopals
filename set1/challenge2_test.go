package set1

import "testing"

func TestHexXOR(t *testing.T) {
	out, err := HexXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if err != nil {
		t.Fatalf("err: %s", err.Error())
	}
	if out != "746865206b696420646f6e277420706c6179" {
		t.Fatal("output is not correct")
	}
}
