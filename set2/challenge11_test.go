package set2

import (
	"fmt"
	"testing"
)

func TestDetectModeOfOperation(t *testing.T) {
	mode, err := DetectModeOfOperation(NewEncrypter())
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	fmt.Printf("Prediction: %s\n", mode)
}
