package set1

import (
	"bytes"
	"errors"
	"math"
)

func SingleCharXORed(hex string) (string, byte, error) {
	if len(hex)%2 != 0 {
		return "", 0, errors.New("invalid hex (not divisable by 2")
	}

	var minLossKey byte
	currFrequencies := make([]float64, 26)
	var sum float64
	minLoss := math.MaxFloat64

	for i := 0; i < 256; i++ {
		for i := range currFrequencies {
			currFrequencies[i] = 0
		}
		sum = 0

		for j := 0; j < len(hex); j += 2 {
			b := ((hexReverseMapping[rune(lower(hex[j]))] * 16) + (hexReverseMapping[rune(lower(hex[j+1]))])) ^ byte(i)
			if 'a' <= b && b <= 'z' {
				currFrequencies[b-'a']++
			}
			sum++
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

	buf := bytes.NewBuffer(make([]byte, 0, len(hex)/2))

	for i := 0; i < len(hex); i += 2 {
		buf.WriteByte(((hexReverseMapping[rune(lower(hex[i]))] * 16) + (hexReverseMapping[rune(lower(hex[i+1]))])) ^ byte(minLossKey))
	}

	return buf.String(), minLossKey, nil
}
