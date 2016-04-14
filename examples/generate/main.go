package main

import "github.com/WindomZ/go-jwt/jwt/generate"

func main() {
	// Generate a hmac key file
	if err := jwt.GenerateHmacFile("./hmac_user1"); err != nil {
		panic(err)
	}

	// Generate a rsa key pair files
	if err := jwt.GenerateRSAFile("./rsa_admin1"); err != nil {
		panic(err)
	}
}
