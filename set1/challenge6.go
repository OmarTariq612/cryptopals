package set1

import (
	"errors"
	"os"
)

func HammingDistance(a, b []byte) (uint, error) {
	if len(a) != len(b) {
		return 0, errors.New("mismatch length (a and b must have the same length)")
	}

	var count uint = 0

	for i := range a {
		xor := a[i] ^ b[i]
		for xor > 0 {
			count += uint(xor & 1)
			xor >>= 1
		}
	}

	return count, nil
}

func DecryptWithRepeatingXORKey(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return nil, nil
}
