package jwt

import (
	"fmt"
	. "github.com/WindomZ/go-jwt/jwt/utils"
	"path"
)

func init() {
	dir, ok := GetCurrentDirectory()
	if !ok {
		panic(ErrInit)
	}
	println(ok)
	println(dir)
	s, _ := findJwtKeyFiles(path.Join(dir, "default"))
	println(fmt.Sprintf("%#v", s))
}
