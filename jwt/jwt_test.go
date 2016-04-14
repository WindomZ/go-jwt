package jwt

//import (
//	"testing"
//)
//
//func runJWT(t *testing.T, kid string, m interface{}) {
//	tokenString, err := Signed(kid, m, 72)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	t.Logf("token:%v\n", tokenString)
//	arg, err := ParseFromString(tokenString)
//	if err != nil {
//		t.Error(err)
//		return
//	} else {
//		t.Log(arg)
//	}
//}
//
//func TestUser1(t *testing.T) {
//	var m = make(map[string]interface{})
//	m["uid"] = 19
//	m["username"] = "the name of user"
//	m["中文"] = "这是个测试123abc"
//	runJWT(t, KID_USER_V1, m)
//}
//
//func TestAdmin1(t *testing.T) {
//	var m = make(map[string]interface{})
//	m["uid"] = 21
//	m["username"] = "the name of admin"
//	m["中文"] = "这是个测试123abc"
//	runJWT(t, KID_ADMIN_V1, m)
//}

//func TestGenerateUser1(t *testing.T) {
//	err := jwt.SaveHmac256ToFile("./hmac_test1", "secretofgeneplayer", "messageofgeneplayer")
//	if err != nil {
//		panic(err)
//	}
//}
