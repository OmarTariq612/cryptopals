package set2

func PKCS7Pad(input []byte, blockSize int) []byte {
	reqPaddingLength := (((len(input) + blockSize) / blockSize) * blockSize) - len(input)
	output := make([]byte, len(input)+reqPaddingLength)
	copy(output, input)

	for i := range output[len(input):] {
		output[i+len(input)] = byte(reqPaddingLength)
	}

	return output
}
