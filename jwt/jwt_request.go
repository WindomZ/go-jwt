package jwt

import (
	. "github.com/WindomZ/go-jwt/jwt/macro"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func SignedToRequest(req *http.Request, kid string, arg interface{}, minutes int) (*http.Request, error) {
	tokenString, err := Signed(kid, arg, minutes)
	if err != nil {
		return req, err
	}
	req.Header.Set(HEADER_KEY(), HEADER_VALUE(tokenString))
	return req, nil
}

func ParseFromRequest(req *http.Request) (interface{}, error) {
	if req == nil {
		return nil, ErrRequest
	} else if ah := req.Header.Get("Authorization"); ah != "" {
		if len(ah) > 6 && strings.ToUpper(ah[0:7]) == "BEARER " {
			return parseToken(jwt.Parse(ah[7:], getJwtHandlerKey))
		}
	}
	return nil, jwt.ErrNoTokenInRequest
}
