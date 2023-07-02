package set1

import (
	"errors"
	"strings"
)

func HexToBase64(hex string) (string, error) {
	if len(hex)%2 != 0 {
		return "", errors.New("invalid hex (not divisable by 2)")
	}

	hexLength := len(hex)
	bytes := make([]byte, hexLength/2)
	bytesLength := len(bytes)
	for i := 0; i < hexLength; i += 2 {
		s := hex[i : i+2]
		bytes[i/2] = (hexReverseMapping[rune(lower(s[0]))] * 16) + hexReverseMapping[rune(lower(s[1]))]
	}

	out := &strings.Builder{}
	base64Length := (bytesLength / 3)
	if bytesLength%3 > 0 {
		base64Length++
	}
	base64Length *= 4
	out.Grow(base64Length)

	it := 0
	var temp uint32

	for i := 0; i < bytesLength/3; i++ {
		temp = uint32(bytes[it]) << 16
		temp += uint32(bytes[it+1]) << 8
		temp += uint32(bytes[it+2])
		it += 3

		out.WriteByte(base64Space[(temp&0x00FC0000)>>18])
		out.WriteByte(base64Space[(temp&0x0003F000)>>12])
		out.WriteByte(base64Space[(temp&0x00000FC0)>>6])
		out.WriteByte(base64Space[(temp & 0x0000003F)])
	}

	switch bytesLength % 3 {
	case 1:
		temp = uint32(bytes[it]) << 16
		out.WriteByte(base64Space[(temp&0x00FC0000)>>18])
		out.WriteByte(base64Space[(temp&0x0003F000)>>12])
		out.WriteByte('=')
		out.WriteByte('=')
	case 2:
		temp = uint32(bytes[it]) << 16
		temp += uint32(bytes[it+1]) << 8
		out.WriteByte(base64Space[(temp&0x00FC0000)>>18])
		out.WriteByte(base64Space[(temp&0x0003F000)>>12])
		out.WriteByte(base64Space[(temp&0x00000FC0)>>6])
		out.WriteByte('=')
	}

	return out.String(), nil
}
