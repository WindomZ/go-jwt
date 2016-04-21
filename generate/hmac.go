package jwt

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	. "github.com/WindomZ/go-jwt/default"
	"github.com/WindomZ/go-random/random"
	"io/ioutil"
	"path"
	"strings"
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
func GenerateHmacFile(filepath string) error {
	if data, err := generateHmac512(
		random.RandomBytes(sha512.Size),
		random.RandomBytes(sha512.BlockSize),
	); err != nil {
		return err
	} else if err := ioutil.WriteFile(filepath, data, 0644); err != nil {
		return err
	}
	return nil
}

func CorrectHmacFileName(filepath string) string {
	name := path.Base(filepath)
	if strings.HasPrefix(name, StringHmac) {
		return filepath
	}
	return path.Join(path.Dir(filepath), fmt.Sprintf("%v_%v", StringHmac, name))
}
