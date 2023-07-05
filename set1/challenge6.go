package set1

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math"
	"math/bits"
	"os"
)

func HammingDistance(a, b []byte) (uint, error) {
	if len(a) != len(b) {
		return 0, errors.New("mismatch length (a and b must have the same length)")
	}

	// bits.OnesCount8()

	var count uint = 0

	for i := range a {
		count += uint(bits.OnesCount8(a[i] ^ b[i]))
		// for xor > 0 {
		// 	count += uint(xor & 1)
		// 	xor >>= 1
		// }
	}

	return count, nil
}

func DecryptWithRepeatingXORKey(filename string) ([]byte, error) {
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
	var reqKeySize int
	minHammingDistance := math.MaxFloat64

	for keySize := 2; keySize < 41; keySize++ {
		distance, err := HammingDistance(input[:keySize*4], input[keySize*4:keySize*8])
		if err != nil {
			return nil, err
		}

		normalizedHammingDistance := float64(distance) / float64(keySize)
		if normalizedHammingDistance < minHammingDistance {
			minHammingDistance = normalizedHammingDistance
			reqKeySize = keySize
		}
	}

	key := make([]byte, reqKeySize)
	subInputLength := len(input) / reqKeySize
	if len(input)%reqKeySize != 0 {
		subInputLength++
	}
	subInput := make([]byte, subInputLength)

	for i := 0; i < reqKeySize; i++ {
		for j := i; j < len(input); j += reqKeySize {
			subInput[(j-i)/reqKeySize] = input[j]
		}
		key[i] = DetectSingleCharXORed(subInput)
	}

	for i := range input {
		input[i] ^= key[i%reqKeySize]
	}

	fmt.Println(string(input))

	return key, nil
}

func DetectSingleCharXORed(input []byte) byte {
	var minLossKey byte
	currFrequencies := make([]float64, 26)
	var sum float64
	minLoss := math.MaxFloat64

	for i := 0; i < 256; i++ {
		for i := range currFrequencies {
			currFrequencies[i] = 0
		}
		sum = 0

		for _, b := range input {
			b ^= byte(i)
			if 'a' <= b && b <= 'z' {
				currFrequencies[b-'a']++
				sum++
			}
			// sum++
		}

		loss := 0.0

		if sum == 0 {
			loss = math.MaxFloat64
		} else {
			for j := range currFrequencies {
				loss += math.Abs((currFrequencies[j] / sum) - relativeFrequencies[j])
			}
		}

		if loss < minLoss {
			minLoss = loss
			minLossKey = byte(i)
		}
	}

	return minLossKey
}
