package set2

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
)

// xor bytes in place (a)
// a = a ^ b
func XORTwoByteSlicesInplace(a, b []byte) {
	min := a
	if len(b) < len(min) {
		min = b
	}

	for i := range min {
		a[i] ^= b[i]
	}
}

func IsItECB(input []byte, blockSize int) (bool, error) {
	if len(input)%blockSize != 0 {
		return false, errors.New("input must be a multiple of the block size")
	}

	seen := make(map[string]struct{})
	for i := 0; i*blockSize < len(input); i++ {

		if _, ok := seen[string(input[i*blockSize:(i+1)*blockSize])]; ok {
			return true, nil
		}

		seen[string(input[i*blockSize:(i+1)*blockSize])] = struct{}{}
	}

	return false, nil
}

func PKCS7Unpad(input []byte, blockSize int) []byte {
	if len(input)%blockSize != 0 {
		panic("len of input is not divisible by block size")
	}
	paddingLength := input[len(input)-1]
	if paddingLength > byte(blockSize) {
		panic(fmt.Sprintf("padding length is greater than block size (%d, %d)", paddingLength, blockSize))
	}

	return input[:len(input)-int(paddingLength)]
}

func EncryptECBMode(b cipher.Block, plaintext []byte) []byte {
	blockSize := b.BlockSize()
	// add padding
	plaintext = PKCS7Pad(plaintext, blockSize)
	ciphertext := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i += blockSize {
		b.Encrypt(ciphertext[i:i+blockSize], plaintext[i:i+blockSize])
	}

	return ciphertext
}

func DecryptECBMode(b cipher.Block, ciphertext []byte) ([]byte, error) {
	blockSize := b.BlockSize()
	if len(ciphertext)%blockSize != 0 {
		return nil, errors.New("ciphertext length must be a multiple of the block size")
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += blockSize {
		b.Decrypt(plaintext[i:i+blockSize], ciphertext[i:i+blockSize])
	}

	// remove padding
	// plaintext = plaintext[:len(plaintext)-int(plaintext[len(plaintext)-1])]
	plaintext = PKCS7Unpad(plaintext, blockSize)

	return plaintext, nil
}

func NewAES128Cipher() (cipher.Block, error) {
	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func NewECBEncrypterFromBlockCipher(block cipher.Block) Encrypter {
	return func(b []byte) ([]byte, error) {
		return EncryptECBMode(block, b), nil
	}
}
