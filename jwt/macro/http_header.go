package jwt

import "fmt"

var (
	headerKey         string = "Authorization"
	headerValuePrefix        = "Bearer"
)

func SetHeadKeyAndValuePrefix(key, pre string) {
	headerKey = key
	headerValuePrefix = pre
}

// the key in http.Request header entries
func HEADER_KEY() string {
	return headerKey
}

// the value in http.Request header entries
func HEADER_VALUE(v string) string {
	return fmt.Sprintf("%v %v", headerValuePrefix, v)
}
