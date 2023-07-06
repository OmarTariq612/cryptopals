package set2

// xor bytes in place (a)
// a = a ^ b
func XORTwoByteSlicesInplace(a, b []byte) {
	min := a
	if len(b) < len(min) {
		min = b
	}

	for i := range min {
		a[i] ^= b[i]
	}
}
