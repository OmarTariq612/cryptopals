package set1

import (
	"errors"
	"strings"
)

func HexXOR(hex1, hex2 string) (string, error) {
	if len(hex1)%2 != 0 || len(hex2)%2 != 0 {
		return "", errors.New("invalid hex (not divisable by 2)")
	}
	if len(hex1) != len(hex2) {
		return "", errors.New("hex1 and hex2 must have the same length")
	}

	out := &strings.Builder{}
	out.Grow(len(hex1))

	for i := 0; i < len(hex1); i++ {
		out.WriteByte(base16space[hexReverseMapping[rune(lower(hex1[i]))]^hexReverseMapping[rune(lower(hex2[i]))]])
	}

	return out.String(), nil
}
