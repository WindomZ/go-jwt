package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// Generate the signing string.
func Sign(kid string, arg interface{}, minutes int) (string, error) {
	if minutes < 0 {
		minutes = 0
	}
	handler, err := getJwtHandler(kid)
	if err != nil {
		return "", err
	}
	token := jwt.New(handler.Method)
	token.Header[KID] = handler.Kid
	token.Claims[ARG] = arg
	token.Claims[EXP] = time.Now().Add(time.Minute * time.Duration(minutes)).Unix()
	return token.SignedString(handler.enKey)
}

func parseToken(token *jwt.Token, err error) (interface{}, error) {
	if token == nil {
		return nil, ErrToken
	} else if token.Valid {
		return token.Claims[ARG], nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			err = ErrToken
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			err = ErrTokenExpired
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			err = ErrTokenNotActive
		} else {
			err = ErrHandleTokenFunc(err)
		}
	} else {
		err = ErrHandleTokenFunc(err)
	}
	return nil, err
}

// Parse, validate, and return a token.
func Parse(token string) (interface{}, error) {
	return parseToken(jwt.Parse(token, getJwtHandlerKey))
}
