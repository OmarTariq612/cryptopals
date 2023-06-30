package main

import (
	"fmt"

	"github.com/OmarTariq612/cryptopals/set1"
)

func main() {
	out, key, err := set1.SingleCharXORed("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	fmt.Println(key)
}
