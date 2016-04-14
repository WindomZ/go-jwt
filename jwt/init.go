package jwt

import (
	. "github.com/WindomZ/go-jwt/jwt/utils"
	"path"
)

func init() {
	dir, ok := GetCurrentDirectory()
	if !ok {
		panic(ErrInit)
	}
	s, _ := findJwtKeyFiles(path.Join(dir, "default"))
	hs, err := filesToHandlers(s)
	if err != nil {
		panic(err)
	}
	for _, h := range hs {
		println(h.Kid)
	}
}
