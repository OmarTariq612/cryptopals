package set1

import (
	"fmt"
	"testing"
)

func TestDetectAESECBMode(t *testing.T) {
	line, err := DetectAESECBMode("8.txt")
	if err != nil {
		t.Fatalf("err: %s\n", err.Error())
	}

	const blockSize = 16 * 2
	blocks := len(line) / blockSize
	for i := 0; i < blocks; i++ {
		fmt.Println(line[i*blockSize : (i+1)*blockSize])
	}
}
