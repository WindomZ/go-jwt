package main

import (
	"fmt"
	"github.com/WindomZ/go-jwt/jwt"
	"os"
)

const (
	KeyNameHmac string = "hmac_demo" // The name of a hmac key file
	KeyNameRSA         = "rsa_demo"  // The name of a rsa key file
)

func main() {
	// Load key files and take effect
	if dir, err := os.Getwd(); err != nil {
		panic(err)
	} else if err := jwt.NewConfig(dir).Effect(); err != nil {
		panic(err)
	}

	// run demoHmac
	if err := demoHmac(); err != nil {
		panic(err)
	}

	// run demoRSA
	if err := demoRSA(); err != nil {
		panic(err)
	}

	fmt.Println("Success!")
}

func sign(keyName string, m interface{}) (token string, err error) {
	return jwt.Sign(keyName, m, 72)
}

func verify(token string) (err error) {
	_, err = jwt.Parse(token)
	return
}

var test_case = map[string]interface{}{
	"number":  19,
	"english": "This is the English test.",
	"中文":      "这是个中文测试。",
}

func demoHmac() error {
	if token, err := sign(KeyNameHmac, test_case); err != nil {
		return err
	} else if err := verify(token); err != nil {
		return err
	}
	return nil
}

func demoRSA() error {
	if token, err := sign(KeyNameRSA, test_case); err != nil {
		return err
	} else if err := verify(token); err != nil {
		return err
	}
	return nil
}
