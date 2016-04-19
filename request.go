package jwt

import (
	. "github.com/WindomZ/go-jwt/macro"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func SignRequest(req *http.Request, kid string, arg interface{}, minutes int) error {
	token, err := Sign(kid, arg, minutes)
	if err != nil {
		return err
	}
	req.Header.Set(HEADER_KEY(), HEADER_VALUE(token))
	return nil
}

func ParseRequest(req *http.Request) (interface{}, error) {
	if req == nil {
		return nil, ErrRequest
	} else if ah := req.Header.Get("Authorization"); ah != "" {
		if len(ah) > 6 && strings.ToUpper(ah[0:7]) == "BEARER " {
			return parseToken(jwt.Parse(ah[7:], getJwtHandlerKey))
		}
	}
	return nil, jwt.ErrNoTokenInRequest
}
