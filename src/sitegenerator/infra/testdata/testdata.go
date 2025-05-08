package testdata

import (
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "data")
}

func AbsPath(relativePath string) string {
	return filepath.Join(RootDir(), relativePath)
}
