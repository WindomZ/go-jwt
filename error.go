package jwt

import (
	"errors"
	"fmt"
)

var (
	ErrInit error = errors.New("jwt: Fail to initialize jwt")
	ErrNil        = errors.New("jwt: This is nil")
)

var (
	ErrToken          error = errors.New("jwt: That's not even a token")
	ErrTokenExpired         = errors.New("jwt: Token is expired")
	ErrTokenNotActive       = errors.New("jwt: Token is not active yet")
	ErrRequest              = errors.New("jwt: Request error")
	ErrResponse             = errors.New("jwt: Response error")
)

var (
	ErrJwtHandler  error = errors.New("jwt: That's not even a jwt handler")
	ErrJwtHandlers       = errors.New("jwt: These are not jwt handler")
	ErrExistKID          = errors.New("jwt: There is no this KID")
)

var (
	ErrHandleTokenFunc = func(err error) error {
		return errors.New(fmt.Sprintf("jwt: Couldn't handle this token with error(%v)", err.Error()))
	}
)

func IsTimeOutErr(err error) bool {
	return err == ErrTokenExpired
}
