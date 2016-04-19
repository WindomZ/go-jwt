package jwt

import (
	. "github.com/WindomZ/go-jwt/utils"
	"path"
)

const (
	DEFAULT_KID_HMAC string = "hmac_default"
	DEFAULT_KID_RSA         = "rsa_default"
)

// Initialize the default configuration
func init() {
	if dir, ok := GetCurrentDirectory(); !ok {
		panic(ErrInit)
	} else if err := NewConfig(path.Join(dir, "default")).Effect(); err != nil {
		panic(err)
	}
}
