package set1

import (
	"errors"
	"strings"
)

func EncryptWithRepeatingXORKey(input, key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key length must be at least one byte")
	}

	inputLength := len(input)
	keyLength := len(key)

	bytes := make([]byte, inputLength)
	for i := 0; i < inputLength; i++ {
		bytes[i] = input[i] ^ key[i%keyLength]
	}

	out := &strings.Builder{}
	out.Grow(inputLength / 2)
	for i := 0; i < inputLength; i++ {
		out.WriteByte(base16space[bytes[i]>>4])
		out.WriteByte(base16space[bytes[i]&0b00001111])
	}

	return out.String(), nil
}
