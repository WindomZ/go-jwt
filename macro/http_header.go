package jwt

import "fmt"

var (
	headerKey         string = "Authorization"
	headerValuePrefix        = "Bearer"
)

// set key and prefix value in http.Request header entries
func SetHTTPHeaderKeyAndValuePrefix(key, pre string) {
	headerKey = key
	headerValuePrefix = pre
}

// the key in http.Request header entries
func HEADER_KEY() string {
	return headerKey
}

// the prefix value in http.Request header entries
func HEADER_VALUE_PREFIX() string {
	return headerValuePrefix
}

// the value in http.Request header entries
func HEADER_VALUE(v string) string {
	if len(headerValuePrefix) != 0 {
		return fmt.Sprintf("%v %v", headerValuePrefix, v)
	}
	return v
}
