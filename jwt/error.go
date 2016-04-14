package jwt

import "errors"

var (
	ErrInit error = errors.New("Fail to initialize jwt")
	ErrNil        = errors.New("This is nil")
)

var (
	ErrToken          error = errors.New("That's not even a token")
	ErrTokenExpired         = errors.New("Token is expired")
	ErrTokenNotActive       = errors.New("Token is not active yet")
	ErrRequest              = errors.New("Request error")
)

var (
	ErrJwtHandler  error = errors.New("That's not even a jwt handler")
	ErrJwtHandlers       = errors.New("These are not jwt handler")
	ErrExistKID          = errors.New("There is no this KID")
)
