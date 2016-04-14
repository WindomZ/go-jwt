package jwt

import (
	. "github.com/WindomZ/go-jwt/jwt/utils"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"path"
	"strings"
)

type jwtKeyFile struct {
	Tag  string
	Ext  bool
	Kid  string
	Data []byte
}

func findJwtKeyFiles(dir string) ([]*jwtKeyFile, int) {
	paths, err := GetFileList(dir)
	if err != nil {
		return nil, -1
	} else if len(paths) == 0 {
		return nil, 0
	}
	files := make([]*jwtKeyFile, 0, len(paths))
	for _, p := range paths {
		if f, ok := parseJwtKeyFile(p); ok {
			files = append(files, f)
		}
	}
	return files, len(files)
}

func parseJwtKeyFile(file string) (*jwtKeyFile, bool) {
	f := path.Base(file)
	ext := path.Ext(f)
	f = f[:len(f)-len(ext)]
	if len(f) == 0 {
		return nil, false
	} else if data, err := ioutil.ReadFile(file); err != nil {
		return nil, false
	} else if strings.HasPrefix(f, TagHmac) {
		if len(f) > len(TagHmac)+1 {
			return &jwtKeyFile{Tag: TagHmac, Ext: (len(ext) != 0), Kid: f, Data: data}, true
		}
	} else if strings.HasPrefix(f, TagRSA) {
		if len(f) > len(TagRSA)+1 {
			return &jwtKeyFile{Tag: TagRSA, Ext: (len(ext) != 0), Kid: f, Data: data}, true
		}
	}
	return nil, false
}

func filesToHandlers(files []*jwtKeyFile) ([]*jwtHandler, error) {
	if files == nil {
		return nil, ErrNil
	}
	hm := make(map[string]*jwtHandler, len(files))
	for _, f := range files {
		kf := func(t *jwt.Token) (interface{}, error) { return f.Data, nil }
		switch f.Tag {
		case TagHmac:
			hm[f.Kid] = &jwtHandler{
				Kid:    f.Kid,
				Method: jwt.SigningMethodHS512,
				enKey:  kf,
				deKey:  kf,
			}
		case TagRSA:
			if h, ok := hm[f.Kid]; ok {
				if f.Ext {
					h.deKey = kf
				} else {
					h.enKey = kf
				}
			} else {
				hm[f.Kid] = &jwtHandler{
					Kid:    f.Kid,
					Method: jwt.SigningMethodRS512,
					enKey:  kf,
					deKey:  kf,
				}
			}
		}
	}
	if len(hm) == 0 {
		return nil, ErrJwtHandlers
	}
	hs := make([]*jwtHandler, 0, len(hm))
	for _, h := range hm {
		hs = append(hs, h)
	}
	return hs, nil
}
