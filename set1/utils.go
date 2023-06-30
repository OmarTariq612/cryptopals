package set1

var (
	base16space = "0123456789abcdef"
	base64Space = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	hexReverseMapping map[rune]byte = map[rune]byte{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'a': 10,
		'b': 11,
		'c': 12,
		'd': 13,
		'e': 14,
		'f': 15,
	}

	relativeFrequencies = []float64{
		0.0817,
		0.0150,
		0.0278,
		0.0425,
		0.1270,
		0.0223,
		0.0202,
		0.0609,
		0.0697,
		0.0015,
		0.0077,
		0.0403,
		0.0241,
		0.0675,
		0.0751,
		0.0193,
		0.0010,
		0.0599,
		0.0633,
		0.0906,
		0.0276,
		0.0098,
		0.0236,
		0.0015,
		0.0197,
		0.0007,
	}
)

func lower(c byte) byte {
	return c | ('x' - 'X')
}
