package set1

import (
	"bufio"
	"bytes"
	"math"
	"os"
	"strings"
)

func DetectSingleXORedCharHex(filename string) (string, string, byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", "", 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currFrequencies := make([]uint64, 26)
	var sum float64
	type KeyLossHexEntry struct {
		Key  byte
		Loss float64
		Hex  string
	}

	var globalMinLossEntry KeyLossHexEntry
	var localMinLossEntry KeyLossHexEntry
	globalMinLoss := math.MaxFloat64

	for scanner.Scan() {
		hex := strings.Trim(scanner.Text(), "\r\n")
		localMinLoss := math.MaxFloat64

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

			currLoss := 0.0

			if sum == 0 {
				currLoss = math.MaxFloat64
			} else {
				for j := range currFrequencies {
					currLoss += math.Abs((float64(currFrequencies[j]) / sum) - relativeFrequencies[j])
				}
			}

			if currLoss < localMinLoss {
				localMinLoss = currLoss
				localMinLossEntry.Key = byte(i)
				localMinLossEntry.Hex = hex
				localMinLossEntry.Loss = currLoss
			}
		}

		if localMinLoss < globalMinLoss {
			globalMinLoss = localMinLoss
			globalMinLossEntry = localMinLossEntry
		}
	}

	buf := bytes.NewBuffer(make([]byte, 0, len(globalMinLossEntry.Hex)/2))
	for i := 0; i < len(globalMinLossEntry.Hex); i += 2 {
		buf.WriteByte(((hexReverseMapping[rune(lower(globalMinLossEntry.Hex[i]))] * 16) + (hexReverseMapping[rune(lower(globalMinLossEntry.Hex[i+1]))])) ^ byte(globalMinLossEntry.Key))
	}

	return globalMinLossEntry.Hex, buf.String(), globalMinLossEntry.Key, nil
}
