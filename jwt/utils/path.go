package jwt

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func GetCurrentDirectory() (string, bool) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return filename, false
	}
	return strings.Replace(path.Dir(filename), "\\", "/", -1), true
}

func GetFileList(dir string) (paths []string, err error) {
	err = filepath.Walk(
		dir,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			} else if f == nil {
				return nil
			} else if f.IsDir() {
				return nil
			}
			paths = append(paths, path)
			return nil
		},
	)
	return
}
