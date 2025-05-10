package testdata

import (
	"os"
	"path/filepath"
	"runtime"

	"golang.org/x/xerrors"
)

func RootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "data")
}

func AbsPath(relativePath string) string {
	return filepath.Join(RootDir(), relativePath)
}

func CopyToTempDir() (string, error) {
	dir, err := os.MkdirTemp("", "test*")
	if err != nil {
		return "", xerrors.Errorf("failed to create temp directory: %w", dir)
	}
	err = os.CopyFS(dir, os.DirFS(RootDir()))
	if err != nil {
		return "", xerrors.Errorf("failed to copy files into temp directory: %w", dir)
	}
	return dir, nil
}
