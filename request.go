package jwt

import (
	. "github.com/WindomZ/go-jwt/macro"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

// set key and prefix value in http.Request header entries
func SetHTTPHeaderKey(key string) {
	SetHTTPHeaderKeyAndValuePrefix(key, "")
}

// Generate the signing string, and set into http request
func SignRequest(req *http.Request, kid string, arg interface{}, minutes int) error {
	token, err := Sign(kid, arg, minutes)
	if err != nil {
		return err
	}
	req.Header.Set(HEADER_KEY(), HEADER_VALUE(token))
	return nil
}

// Generate the signing string, and set into http response writer
func SignResponse(rw http.ResponseWriter, kid string, arg interface{}, minutes int) error {
	token, err := Sign(kid, arg, minutes)
	if err != nil {
		return err
	}
	rw.Header().Set(HEADER_KEY(), HEADER_VALUE(token))
	return nil
}

// Parse http request, validate, and return a token.
func ParseRequest(req *http.Request) (interface{}, error) {
	if req == nil {
		return nil, ErrRequest
	} else if ah := req.Header.Get(HEADER_KEY()); len(ah) != 0 {
		if len(HEADER_VALUE_PREFIX()) == 0 {
			return parseToken(jwt.Parse(ah, getJwtHandlerKey))
		} else if strings.HasPrefix(ah[:len(HEADER_VALUE_PREFIX())], HEADER_VALUE_PREFIX()) {
			return parseToken(jwt.Parse(ah[(len(HEADER_VALUE_PREFIX())+1):],
				getJwtHandlerKey))
		}
	}
	return nil, jwt.ErrNoTokenInRequest
}
