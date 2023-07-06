package set2

import (
	"crypto/cipher"
	"errors"
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

func EncryptECBMode(b cipher.Block, plaintext []byte) []byte {
	blockSize := b.BlockSize()
	// add padding
	plaintext = PKCS7Pad(plaintext, blockSize)
	ciphertext := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i += blockSize {
		b.Decrypt(ciphertext[i:i+blockSize], plaintext[i:i+blockSize])
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
	plaintext = plaintext[:len(plaintext)-int(plaintext[len(plaintext)-1])]

	return plaintext, nil
}
