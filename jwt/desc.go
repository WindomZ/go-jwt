package jwt

const (
	MSG_ERR_JWT_INVALID_TOKEN       string = "That's not even a token"
	MSG_ERR_JWT_EXPIRED_TOKEN              = "Token is expired"
	MSG_ERR_JWT_NOT_VALID_YET_TOKEN        = "Token is not active yet"
	MSG_ERR_JWT_CANNOT_HANDLE_TOKEN        = "Couldn't handle this token:"
	MSG_ERR_JWT_REQUEST                    = "Request error"
)
