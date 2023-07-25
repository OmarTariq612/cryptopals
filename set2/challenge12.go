package set2

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"errors"
	"fmt"
)

func NewECBSimpleEncrypter(unkownBytes []byte) Encrypter {
	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	return func(plaintext []byte) ([]byte, error) {
		plaintext = append(plaintext, unkownBytes...)
		return EncryptECBMode(block, plaintext), nil
	}
}

func ECBDecryption(e Encrypter) ([]byte, error) {
	input := make([]byte, 0, 32)
	lastLength := 0
	var blockSize int

	for {
		input = append(input, 0)
		ciphertext, err := e(input)
		if err != nil {
			return nil, err
		}

		if lastLength == 0 {
			lastLength = len(ciphertext)
		} else {
			if lastLength != len(ciphertext) {
				blockSize = len(ciphertext) - lastLength
				break
			}
		}
	}

	fmt.Printf("block_size = %d\n", blockSize)

	input = make([]byte, 2*blockSize)
	copy(input, bytes.Repeat([]byte{0}, 2*blockSize))
	ciphertext, err := e(input)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(ciphertext[:blockSize], ciphertext[blockSize:2*blockSize]) {
		return nil, errors.New("mode of operation is not ECB mode")
	}

	fmt.Println("ECB mode")

	unKownBytesSizeInBlocks := len(ciphertext[2*blockSize:]) / blockSize
	input = make([]byte, unKownBytesSizeInBlocks*blockSize)
	counter := 1

	for {
		ciphertext, err := e(input[:len(input)-counter])
		if err != nil {
			return nil, err
		}

		reqBlock := ciphertext[(len(ciphertext)/blockSize)/2 : ((len(ciphertext)/blockSize)+1)/2]

		for j := 0; j < 256; j++ {
			input[len(input)-1] = byte(j)
			ciphertext, err := e(input)
			if err != nil {
				return nil, err
			}

			currBlock := ciphertext[(len(ciphertext)/blockSize)/2 : ((len(ciphertext)/blockSize)+1)/2]

			if bytes.Equal(reqBlock, currBlock) {
				for i := 1; i < len(input); i++ {
					input[i-1] = input[i]
				}
				break
			}
		}

		if len(ciphertext) <= len(input) {
			input = input[len(input)-counter : len(input)-1]
			break
		}

		counter++
	}

	return input, nil
}
