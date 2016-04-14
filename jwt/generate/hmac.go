package jwt

import (
	"crypto/hmac"
	"crypto/sha512"
	"io/ioutil"
)

func generateHmac512(message, secret string) ([]byte, error) {
	h := hmac.New(sha512.New, []byte(secret))
	if _, err := h.Write([]byte(message)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func GenerateHmacFile(filename, message, secret string) error {
	if data, err := generateHmac512(message, secret); err != nil {
		return err
	} else if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return err
	}
	return nil
}
