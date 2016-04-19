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
	} else if ah := req.Header.Get(HEADER_KEY()); ah != "" {
		if len(HEADER_VALUE_PREFIX()) == 0 {
			return parseToken(jwt.Parse(ah, getJwtHandlerKey))
		} else if strings.HasPrefix(
			strings.ToUpper(ah[:len(HEADER_VALUE_PREFIX())]),
			HEADER_VALUE_PREFIX()) {
			return parseToken(jwt.Parse(ah[(len(HEADER_VALUE_PREFIX())+1):],
				getJwtHandlerKey))
		}
	}
	return nil, jwt.ErrNoTokenInRequest
}