package set1

import (
	"bytes"
	"errors"
	"math"
	"sort"
)

func SingleCharXORed(hex string) (string, byte, error) {
	if len(hex)%2 != 0 {
		return "", 0, errors.New("invalid hex (not divisable by 2")
	}

	res := make([]struct {
		Loss float64
		Key  byte
	}, 256)

	currFrequencies := make([]float64, 26)
	var sum float64

	for i := 0; i < 256; i++ {
		for i := range currFrequencies {
			currFrequencies[i] = 0
		}
		sum = 0

		for j := 0; j < len(hex); j += 2 {
			b := lower((hexReverseMapping[rune(lower(hex[j]))]*16)+(hexReverseMapping[rune(lower(hex[j+1]))])) ^ byte(i)
			if 'a' <= b && b <= 'z' {
				currFrequencies[b-'a']++
				sum++
			}
		}

		res[i].Key = byte(i)
		if sum == 0 {
			res[i].Loss = math.MaxFloat64
		} else {
			for j := range currFrequencies {
				res[i].Loss += math.Abs((currFrequencies[j] / sum) - relativeFrequencies[j])
			}
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Loss < res[j].Loss
	})

	buf := bytes.NewBuffer(make([]byte, 0, len(hex)/2))

	for i := 0; i < len(hex); i += 2 {
		buf.WriteByte(((hexReverseMapping[rune(lower(hex[i]))] * 16) + (hexReverseMapping[rune(lower(hex[i+1]))])) ^ byte(res[0].Key))
	}

	return buf.String(), res[0].Key, nil
}
