package jwt

import "errors"

var (
	ErrToken          error = errors.New(MSG_ERR_JWT_INVALID_TOKEN)
	ErrTokenExpired         = errors.New(MSG_ERR_JWT_EXPIRED_TOKEN)
	ErrTokenNotActive       = errors.New(MSG_ERR_JWT_NOT_VALID_YET_TOKEN)
	ErrRequest              = errors.New(MSG_ERR_JWT_REQUEST)
)
