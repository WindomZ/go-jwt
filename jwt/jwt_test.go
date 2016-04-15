package jwt

import "testing"

func runJWT(t *testing.T, kid string, m interface{}) {
	tokenString, err := Signed(kid, m, 72)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("token: %v\n", tokenString)
	arg, err := Parsing(tokenString)
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Log(arg)
	}
}

func TestDefaultHmac(t *testing.T) {
	var m = make(map[string]interface{})
	m["uid"] = 19
	m["username"] = "the name of user"
	m["中文"] = "这是个测试123abc"
	runJWT(t, DEFAULT_KID_HMAC, m)
}

func TestDefaultRSA(t *testing.T) {
	var m = make(map[string]interface{})
	m["uid"] = 21
	m["username"] = "the name of admin"
	m["中文"] = "这是个测试123abc"
	runJWT(t, DEFAULT_KID_RSA, m)
}
