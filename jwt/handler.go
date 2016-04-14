package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	KID = "kid"
	ARG = "arg"
	EXP = "exp"
)

const (
	TagHmac string = "hmac"
	TagRSA         = "rsa"
)

type jwtHandler struct {
	Kid    string
	Method jwt.SigningMethod
	enKey  jwt.Keyfunc
	deKey  jwt.Keyfunc
}

var handlers map[string]*jwtHandler

func setJwtHandler(m *jwtHandler) error {
	if m == nil {
		return ErrJwtHandler
	} else if len(m.Kid) == 0 {
		return ErrExistKID
	}
	handlers[m.Kid] = m
	return nil
}

func getJwtHandler(kid string) (*jwtHandler, error) {
	if m, ok := handlers[kid]; m == nil || !ok {
		return nil, ErrExistKID
	} else {
		return m, nil
	}
}

func getJwtHandlerKey(token *jwt.Token) (interface{}, error) {
	if kid, ok := token.Header[KID].(string); !ok {
		return nil, ErrToken
	} else if m, err := getJwtHandler(kid); err != nil {
		return nil, err
	} else {
		return m.deKey(token)
	}
}
