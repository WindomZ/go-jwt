package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	. "github.com/WindomZ/go-jwt/jwt/macro"
	"os"
	"path"
	"strings"
)

const (
	type_private_key string = "RSA PRIVATE KEY"
	type_public_key  string = "PUBLIC KEY"
)

// Generate a rsa key pair
func generateRSAKey(filepath string, bits int) error {
	if len(filepath) == 0 {
		return os.ErrNotExist
	} else if privateKey, err := rsa.GenerateKey(rand.Reader, bits); err != nil {
		return err
	} else if file, err := os.Create(filepath); err != nil {
		return err
	} else if err := pem.Encode(
		file,
		&pem.Block{
			Type:  type_private_key,
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	); err != nil {
		return err
	} else if pkix, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey); err != nil {
		return err
	} else if file, err := os.Create(filepath + ".pub"); err != nil {
		return err
	} else if err = pem.Encode(
		file,
		&pem.Block{
			Type:  type_public_key,
			Bytes: pkix,
		},
	); err != nil {
		return err
	}
	return nil
}

// Generate a rsa key pair files, and save to a file with name '@filename', like '@filename' and '@filename.pub'
func GenerateRSAFile(filepath string) error {
	return generateRSAKey(filepath, 1024)
}

func CorrectRSAFileName(filepath string) string {
	name := path.Base(filepath)
	if strings.HasPrefix(name, StrRSA) {
		return filepath
	}
	return path.Join(path.Dir(filepath), fmt.Sprintf("%v_%v", StrRSA, name))
}
