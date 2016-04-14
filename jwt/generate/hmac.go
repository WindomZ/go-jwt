package jwt

import (
	"crypto/hmac"
	"crypto/sha512"
	"github.com/WindomZ/go-random/random"
	"io/ioutil"
)

// Generate a hmac key
func generateHmac512(message, secret []byte) ([]byte, error) {
	h := hmac.New(sha512.New, secret)
	if _, err := h.Write(message); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// Generate a hmac key, and save to a file with name '@filename'
func GenerateHmacFile(filename string) error {
	if data, err := generateHmac512(
		random.RandomBytes(sha512.Size),
		random.RandomBytes(sha512.BlockSize),
	); err != nil {
		return err
	} else if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return err
	}
	return nil
}
