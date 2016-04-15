package main

import "github.com/WindomZ/go-jwt/generate"

// Notice: in this project, key file name must prefix 'hmac_' or 'rsa_'
// Recommend to use CorrectHmacFileName and CorrectRSAFileName
func main() {
	// Generate a hmac key file
	if err := jwt.GenerateHmacFile(jwt.CorrectHmacFileName("../demo/demo")); err != nil {
		panic(err)
	}

	// Generate a rsa key pair files
	if err := jwt.GenerateRSAFile(jwt.CorrectRSAFileName("../demo/demo")); err != nil {
		panic(err)
	}
}
