package set2

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

func NewECBEncrypter() Encrypter {
	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	unKnownString := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	unKnownBytes, err := base64.StdEncoding.DecodeString(unKnownString)
	// unKnownBytes := []byte("Hello Teacher Hello Teacher How are you How are you ??? I'm fine thank you I'm fine thank you HOW ARE YOUUUUUU!")
	if err != nil {
		panic(err)
	}

	// fmt.Println(len(unKnownBytes))

	return func(plaintext []byte) ([]byte, error) {
		blockSize := block.BlockSize()
		plaintext = append(plaintext, unKnownBytes...)
		plaintext = PKCS7Pad(plaintext, blockSize)
		ciphertext := make([]byte, len(plaintext))

		for i := 0; i < len(plaintext)/blockSize; i++ {
			block.Encrypt(ciphertext[i*blockSize:(i+1)*blockSize], plaintext[i*blockSize:(i+1)*blockSize])
		}

		return ciphertext, nil
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
