package set1

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

func DecryptAESECB(filename string, key []byte) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	n, err := base64.StdEncoding.Decode(input, input)
	if err != nil {
		return nil, err
	}

	input = input[:n]
	if len(input)%16 != 0 {
		return nil, errors.New("input length is not a multiple of the aes block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(input); i += 16 {
		block.Decrypt(input[i:i+16], input[i:i+16])
	}

	return input, nil
}
