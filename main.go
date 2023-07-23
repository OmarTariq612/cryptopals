package main

func main() {
	// out, key, err := set1.SingleCharXORed("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(out)
	// fmt.Println(key)

	// mode, err := set2.DetectModeOfOperation(set2.NewEncrypter())
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Prediction: %s\n", mode)

	// msg := []byte("Hello World")

	// publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	// if err != nil {
	// 	panic(err)
	// }

	// signature, err := privateKey.Sign(nil, msg, crypto.Hash(0))
	// if err != nil {
	// 	panic(err)
	// }

	// if ed25519.Verify(publicKey, msg, signature) {
	// 	fmt.Println("valid signature")
	// } else {
	// 	fmt.Println("invalid signature")
	// }

	// msg := []byte("Hello World")
	// digest := sha3.Sum256(msg)

	// privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// if err != nil {
	// 	panic(err)
	// }

	// publicKey := privateKey.PublicKey

	// signature, err := privateKey.Sign(rand.Reader, digest[:], nil)
	// if err != nil {
	// 	panic(err)
	// }

	// if ecdsa.VerifyASN1(&publicKey, digest[:], signature) {
	// 	fmt.Println("valid signature")
	// } else {
	// 	fmt.Println("invalid signature")
	// }

	// msg := []byte("Hello World")
	// digest := sha3.Sum256(msg)

	// privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	// if err != nil {
	// 	panic(err)
	// }

	// publicKey := privateKey.PublicKey

	// signature, err := privateKey.Sign(rand.Reader, digest[:], &rsa.PSSOptions{Hash: crypto.SHA3_256, SaltLength: rsa.PSSSaltLengthAuto})
	// if err != nil {
	// 	panic(err)
	// }

	// if rsa.VerifyPSS(&publicKey, crypto.SHA3_256, digest[:], signature, &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto}) == nil {
	// 	fmt.Println("valid signature")
	// } else {
	// 	fmt.Println("invalid signature")
	// }

	// msg := []byte("Hello World")

	// publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	// if err != nil {
	// 	panic(err)
	// }

	// signature, err := privateKey.Sign(nil, msg, crypto.Hash(0))
	// if err != nil {
	// 	panic(err)
	// }

	// if ed25519.Verify(publicKey, msg, signature) {
	// 	fmt.Println("valid signature")
	// } else {
	// 	fmt.Println("invalid signature")
	// }

	// fmt.Println(bytes.Equal([]byte{1, 2, 3}, []byte{1, 2, 3, 0}))

	// var dst [5]byte
	// CountToN(5, dst[:0])
	// fmt.Println(len(dst))
	// fmt.Println(dst)

	// newDst := CountToN(10, nil)
	// fmt.Println(len(newDst))
	// fmt.Println(newDst)
	// packet := make([]byte, 0)
	// packet = append(packet, 1, 2, 3)
	// fmt.Println(packet)

	// fmt.Printf("Allocs: %d\n", int(testing.AllocsPerRun(1, func() {
	// 	fmt.Println(&packet)
	// })))

	// res, err := set2.ParseToOrderedJSON("foo=bar&baz=qux&zap=zazzle")
	// if err != nil {
	// 	panic(err)
	// }

	// res, err := set2.ProfileFor("omar@tariq.com")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(res)

	// output := set2.PKCS7Pad([]byte("omar"), 16)
	// fmt.Println(len(output))
	// fmt.Println(output)
	// output = set2.PKCS7Unpad(output, 16)
	// fmt.Println(len(output))
	// fmt.Println(output)
	// fmt.Println(string(output))

	// plaintext := "email=omar@omar.dev&uid=20&role=admin"
	// block := set2.NewAES128Cipher()
	// ciphertext := set2.EncryptECBMode(block, []byte(plaintext))

	// newPlaintext, err := set2.DecryptECBMode(block, ciphertext)
	// if err != nil {
	// 	panic(err)
	// }

	// res, err := set2.ParseElements(string(newPlaintext))
	// if err != nil {
	// 	panic(err)
	// }
}

// func Func(in [][]byte) [][]byte {
// 	for i, slice := range in {
// 		for j := range slice {
// 			slice[j] = byte(i)
// 		}
// 	}

// 	return in
// }

// func CountToN(n int, dst []byte) []byte {
// 	for i := 0; i < n; i++ {
// 		dst = append(dst, byte(i))
// 	}

// 	return dst
// }

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
