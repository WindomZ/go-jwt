package jwt

import (
	. "github.com/WindomZ/go-jwt/jwt/macro"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	KID = "kid"
	ARG = "arg"
	EXP = "exp"
)

const (
	TagHmac string = StrHmac
	TagRSA         = StrRSA
)

type jwtHandler struct {
	Kid    string
	Method jwt.SigningMethod
	enKey  interface{}
	deKey  interface{}
}

var handlers = make(map[string]*jwtHandler, 5)

func setJwtHandler(h *jwtHandler) error {
	if h == nil {
		return ErrJwtHandler
	} else if len(h.Kid) == 0 {
		return ErrExistKID
	}
	handlers[h.Kid] = h
	return nil
}

func setJwtHandlers(hs []*jwtHandler) error {
	for _, h := range hs {
		if err := setJwtHandler(h); err != nil {
			return err
		}
	}
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
		return m.deKey, nil
	}
}
