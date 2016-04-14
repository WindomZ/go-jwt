package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func generateRSAKey(filename string, bits int) error {
	if len(filename) == 0 {
		return os.ErrNotExist
	} else if privateKey, err := rsa.GenerateKey(rand.Reader, bits); err != nil {
		return err
	} else if file, err := os.Create(filename); err != nil {
		return err
	} else if err := pem.Encode(
		file,
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	); err != nil {
		return err
	} else if pkix, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey); err != nil {
		return err
	} else if file, err := os.Create(filename + ".pub"); err != nil {
		return err
	} else if err = pem.Encode(
		file,
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pkix,
		},
	); err != nil {
		return err
	}
	return nil
}

func GenerateRSAFile(filename string) error {
	return generateRSAKey(filename, 1024)
}
