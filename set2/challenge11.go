package set2

import (
	"crypto/aes"
	cryptoRand "crypto/rand"
	"fmt"
	mathRand "math/rand"
)

type Encrypter func([]byte) ([]byte, error)

func NewEncrypter() Encrypter {
	return func(plaintext []byte) ([]byte, error) {
		key := make([]byte, 16)
		if _, err := cryptoRand.Read(key); err != nil {
			return nil, err
		}

		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}

		prependLength := mathRand.Intn(6) + 5
		appendLength := mathRand.Intn(6) + 5
		plaintextAppended := make([]byte, prependLength+len(plaintext)+appendLength)
		copy(plaintextAppended[prependLength:], plaintext)

		if _, err := cryptoRand.Read(plaintextAppended[:prependLength]); err != nil {
			return nil, err
		}
		if _, err := cryptoRand.Read(plaintextAppended[prependLength+len(plaintext):]); err != nil {
			return nil, err
		}

		if mathRand.Intn(2) == 0 {
			iv := make([]byte, 16)
			if _, err := cryptoRand.Read(iv); err != nil {
				return nil, err
			}
			fmt.Println("Oracle: CBC Mode")
			return EncryptCBCMode(block, iv, plaintextAppended)
		}

		fmt.Println("Oracle: ECB Mode")
		return EncryptECBMode(block, plaintextAppended), nil
	}
}

type Mode byte

const (
	ECB Mode = iota
	CBC
)

func (m Mode) String() string {
	switch m {
	case ECB:
		return "ECB Mode"
	case CBC:
		return "CBC Mode"
	}

	return "UNKNOWN Mode"
}

func DetectModeOfOperation(e Encrypter) (Mode, error) {
	input := make([]byte, 64)
	output, err := e(input)
	if err != nil {
		return 0, err
	}

	isECB, err := IsItECB(output, 16)
	if err != nil {
		return 0, nil
	}
	if isECB {
		return ECB, nil
	}

	return CBC, nil
}
