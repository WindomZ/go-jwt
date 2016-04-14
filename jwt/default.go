package jwt

import (
	. "github.com/WindomZ/go-jwt/jwt/utils"
	"path"
)

const (
	DEFAULT_KID_HMAC string = "hmac_default"
	DEFAULT_KID_RSA         = "rsa_default"
)

func init() {
	dir, ok := GetCurrentDirectory()
	if !ok {
		panic(ErrInit)
	}
	fs, _ := findJwtKeyFiles(path.Join(dir, "default"))
	if hs, err := filesToHandlers(fs); err != nil {
		panic(err)
	} else {
		setJwtHandlers(hs)
	}
}
