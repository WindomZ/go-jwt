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
	if req == nil {
		return ErrRequest
	} else if token, err := Sign(kid, arg, minutes); err != nil {
		return err
	} else {
		req.Header.Set(HEADER_KEY(), HEADER_VALUE(token))
	}
	return nil
}

// Generate the signing string, and set into http response writer
func SignResponse(rw http.ResponseWriter, kid string, arg interface{}, minutes int) error {
	if rw == nil {
		return ErrResponse
	} else if token, err := Sign(kid, arg, minutes); err != nil {
		return err
	} else {
		rw.Header().Set(HEADER_KEY(), HEADER_VALUE(token))
	}
	return nil
}

func parseHTTPHeader(header http.Header) (interface{}, error) {
	if ah := header.Get(HEADER_KEY()); len(ah) != 0 {
		if len(HEADER_VALUE_PREFIX()) == 0 {
			return parseToken(jwt.Parse(ah, getJwtHandlerKey))
		} else if strings.HasPrefix(ah[:len(HEADER_VALUE_PREFIX())], HEADER_VALUE_PREFIX()) {
			return parseToken(jwt.Parse(ah[(len(HEADER_VALUE_PREFIX())+1):],
				getJwtHandlerKey))
		}
	}
	return nil, jwt.ErrNoTokenInRequest
}

// Parse http request, validate, and return a token.
func ParseRequest(req *http.Request) (interface{}, error) {
	if req == nil {
		return nil, ErrRequest
	}
	return parseHTTPHeader(req.Header)
}

// Parse http response, validate, and return a token.
func ParseResponse(resp *http.Response) (interface{}, error) {
	if resp == nil {
		return nil, ErrResponse
	}
	return parseHTTPHeader(resp.Header)
}
