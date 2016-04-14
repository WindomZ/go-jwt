package main

import "github.com/WindomZ/go-jwt/jwt/generate"

func main() {
	if err := jwt.GenerateHmacFile("./hmac_user1", "123456", "key"); err != nil {
		panic(err)
	}
	if err := jwt.GenerateRSAFile("./rsa_admin1"); err != nil {
		panic(err)
	}
}
