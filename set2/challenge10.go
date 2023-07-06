package set2

import (
	"crypto/cipher"
	"errors"
)

func DecryptCBCMode(cipher cipher.Block, iv []byte, ciphertext []byte) ([]byte, error) {
	blockSize := cipher.BlockSize()
	if len(iv) != blockSize {
		return nil, errors.New("length of iv must be equal to the block size")
	}

	prev := iv
	plaintext := make([]byte, len(ciphertext))

	for i := 0; i < len(ciphertext)/blockSize; i++ {
		cipher.Decrypt(plaintext[i*blockSize:(i+1)*blockSize], ciphertext[i*blockSize:(i+1)*blockSize])
		XORTwoByteSlicesInplace(plaintext[i*blockSize:(i+1)*blockSize], prev)
		prev = ciphertext[i*blockSize : (i+1)*blockSize]
	}

	// remove padding
	plaintext = plaintext[:len(plaintext)-int(plaintext[len(plaintext)-1])]

	prev = nil

	return plaintext, nil
}

func EncryptCBCMode(cipher cipher.Block, iv []byte, plaintext []byte) ([]byte, error) {
	blockSize := cipher.BlockSize()
	if len(iv) != blockSize {
		return nil, errors.New("length of iv must be equal to the block size")
	}

	prev := iv
	// add padding
	plaintext = PKCS7Pad(plaintext, blockSize)
	ciphertext := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext)/blockSize; i++ {
		XORTwoByteSlicesInplace(plaintext[i*blockSize:(i+1)*blockSize], prev)
		cipher.Encrypt(ciphertext[i*blockSize:(i+1)*blockSize], plaintext[i*blockSize:(i+1)*blockSize])
		prev = ciphertext[i*blockSize : (i+1)*blockSize]
	}

	prev = nil

	return ciphertext, nil
}
