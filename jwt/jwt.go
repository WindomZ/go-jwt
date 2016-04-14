package jwt

import (
	"errors"
	"fmt"
	. "github.com/WindomZ/go-jwt/jwt/macro"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

func init() {
	initMethods()
}

func Signed(kid string, arg interface{}, minutes int) (string, error) {
	if minutes < 0 {
		minutes = 0
	}
	method := lookupMethod(kid)
	token := jwt.New(method.Method)
	token.Header[KID] = method.Kid
	token.Claims[ARG] = arg
	token.Claims[EXP] = time.Now().Add(time.Minute * time.Duration(minutes)).Unix()
	key, err := method.enKey(token)
	if err == nil {
		return token.SignedString(key)
	}
	return "", err
}

func SignedToRequest(req *http.Request, kid string, arg interface{}, minutes int) (*http.Request, error) {
	tokenString, err := Signed(kid, arg, minutes)
	if err != nil {
		return req, err
	}
	req.Header.Set(HEADER_KEY(), HEADER_VALUE(tokenString))
	return req, nil
}

func parseToken(token *jwt.Token, err error) (interface{}, error) {
	if token == nil {
		return nil, ErrToken
	}
	if token.Valid {
		return token.Claims[ARG], nil
	}
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			err = ErrToken
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			err = ErrTokenExpired
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			err = ErrTokenNotActive
		} else {
			err = errors.New(fmt.Sprintln(MSG_ERR_JWT_CANNOT_HANDLE_TOKEN, err))
		}
	} else {
		err = errors.New(fmt.Sprintln(MSG_ERR_JWT_CANNOT_HANDLE_TOKEN, err))
	}
	return nil, err
}

func ParseFromString(tokenString string) (interface{}, error) {
	return parseToken(jwt.Parse(tokenString, lookupKey))
}

func ParseFromRequest(req *http.Request) (interface{}, error) {
	if req == nil {
		return nil, ErrRequest
	} else if ah := req.Header.Get("Authorization"); ah != "" {
		if len(ah) > 6 && strings.ToUpper(ah[0:7]) == "BEARER " {
			return parseToken(jwt.Parse(ah[7:], lookupKey))
		}
	}
	return nil, jwt.ErrNoTokenInRequest
}
