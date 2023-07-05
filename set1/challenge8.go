package set1

import (
	"bufio"
	"errors"
	"os"
)

func DetectAESECBMode(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		isECB, err := IsItECB(line, 16)
		if err != nil {
			return "", err
		}
		if isECB {
			return line, nil
		}
	}

	return "", errors.New("could not detect aes ecb encrypted line")
}

func IsItECB(hex string, blockSize int) (bool, error) {
	if (len(hex)/2)%blockSize != 0 {
		return false, errors.New("input must be a multiple of the block size")
	}

	seen := make(map[string]struct{})
	for i := 0; i*2*blockSize < len(hex); i++ {

		if _, ok := seen[hex[i*2*blockSize:(i+1)*2*blockSize]]; ok {
			return true, nil
		}

		seen[hex[i*2*blockSize:(i+1)*2*blockSize]] = struct{}{}
	}

	return false, nil
}

// inefficient:
// ===================
// func IsItECB(input []byte, blockSize int) (bool, error) {
// 	if len(input)%blockSize != 0 {
// 		return false, errors.New("input must be a multiple of the block size")
// 	}

// 	for i := 0; i*blockSize < len(input); i++ {
// 		for j := i + 1; j*blockSize < len(input); j++ {
// 			if bytes.Equal(input[i*blockSize:(i+1)*blockSize], input[j*blockSize:(j+1)*blockSize]) {
// 				return true, nil
// 			}
// 		}
// 	}

// 	return false, nil
// }
