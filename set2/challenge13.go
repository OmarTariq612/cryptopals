package set2

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Element struct {
	Key, Value string
}

type Elements []Element

func (es Elements) String() string {
	res := strings.Builder{}
	for i := range es {
		if _, err := res.WriteString(fmt.Sprintf("%s=%s", es[i].Key, es[i].Value)); err != nil {
			panic("this must not happen")
		}
		if i < len(es)-1 {
			res.WriteRune('&')
		}
	}
	return res.String()
}

func ParseElements(s string) (Elements, error) {
	queryParserRes, err := url.ParseQuery(s)
	if err != nil {
		return nil, err
	}

	res := make([]Element, 0, len(queryParserRes))
	for key, value := range queryParserRes {
		res = append(res, Element{Key: key, Value: value[0]})
	}

	return res, nil
}

var userCount int

func init() {
	userCount = 9
}

func ProfileFor(email string) (Elements, error) {
	if strings.ContainsAny(email, "&=") {
		return nil, errors.New("string must not contain invalid chars {&, =}")
	}

	// race condition (I don't care)

	userCount++

	return []Element{
		{"email", email},
		{"uid", strconv.Itoa(userCount)},
		{"role", "user"},
	}, nil
}

// returns ciphertext (that when decrypted will represent an admin account)
func MakeNewAdminAccount(e Encrypter) ([]byte, error) {
	email := []byte("omar@omar.dev") // email=omar@omar.dev&uid=xxxx&role=user
	var ciphertext []byte
	lastLength := -1
	var blockSize int

	for {
		elements, err := ProfileFor(string(email))
		if err != nil {
			return nil, err
		}
		ciphertext, err = e([]byte(elements.String()))
		if err != nil {
			return nil, err
		}
		if lastLength == -1 {
			lastLength = len(ciphertext)
		} else {
			if diff := len(ciphertext) - lastLength; diff > 1 {
				blockSize = diff
				// this will make "user" has its own block
				email = append(email, []byte("oooo")...)
				elements, err := ProfileFor(string(email))
				if err != nil {
					return nil, err
				}
				ciphertext, err = e([]byte(elements.String()))
				if err != nil {
					return nil, err
				}
				break
			}
		}

		email = append(email, 'o')
	}

	adminInput := []byte("admin")
	paddingContent := blockSize - len(adminInput)
	adminInput = append(bytes.Repeat([]byte{0}, blockSize-len("email=")), adminInput...)
	adminInput = append(adminInput, bytes.Repeat([]byte{byte(paddingContent)}, paddingContent)...)

	adminElements, err := ProfileFor(string(adminInput))
	if err != nil {
		return nil, err
	}
	adminCiphertext, err := e([]byte(adminElements.String()))
	if err != nil {
		return nil, err
	}

	// replace the user block with the admin block
	copy(ciphertext[len(ciphertext)-blockSize:], adminCiphertext[blockSize:2*blockSize])

	return ciphertext, nil
}
