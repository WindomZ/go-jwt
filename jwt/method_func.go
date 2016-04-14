package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"os"
	"path"
)

const (
	KID_USER_V1  = "user1"
	KID_ADMIN_V1 = "admin1"
	KID_DEFAULT  = KID_USER_V1
)

var (
//keyUser1    []byte
//keyEnAdmin1 []byte
//keyDeAdmin1 []byte
)

var (
	handlerUser = &jwtHandler{
		Kid:    KID_USER_V1,
		Method: jwt.SigningMethodHS256,
		enKey:  func(t *jwt.Token) (interface{}, error) { return keyUser1, nil },
		deKey:  func(t *jwt.Token) (interface{}, error) { return keyUser1, nil },
	}
	handlerAdmin = &jwtHandler{
		Kid:    KID_ADMIN_V1,
		Method: jwt.SigningMethodRS256,
		enKey:  func(t *jwt.Token) (interface{}, error) { return keyEnAdmin1, nil },
		deKey:  func(t *jwt.Token) (interface{}, error) { return keyDeAdmin1, nil },
	}
)

func initMethods() {
	var e error
	var file string
	if file, e = os.Getwd(); e != nil {
		panic(e)
	}
	//if strings.HasSuffix(file, "/libs/jwt") {
	//	file = ""
	//} else if strings.HasSuffix(file, "/libs") {
	//	file = "jwt"
	//} else {
	//	file = "./libs/jwt"
	//}
	if keyUser1, e = ioutil.ReadFile(path.Join(file, "hmac_user1")); e != nil {
		panic(e)
	}
	if keyEnAdmin1, e = ioutil.ReadFile(path.Join(file, "rsa_admin1")); e != nil {
		panic(e)
	}
	if keyDeAdmin1, e = ioutil.ReadFile(path.Join(file, "rsa_admin1.pub")); e != nil {
		panic(e)
	}
}
