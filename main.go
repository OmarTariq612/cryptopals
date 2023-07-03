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

// BE CAREFUL when using for-loop variables in closures!!!
// func Print123() {
// 	var prints []func()
// 	for i := 1; i <= 3; i++ {
// 		// prints = append(prints, func(n int) func() { return func() { fmt.Println(n) } }(i))
// 		iCopy := i
// 		prints = append(prints, func() { fmt.Println(iCopy) })
// 	}
// 	for _, print := range prints {
// 		print()
// 	}
// }
