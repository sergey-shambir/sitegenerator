package testdata

import (
	"os"
	"path/filepath"
	"runtime"

	"golang.org/x/xerrors"
)

func testDataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func ContentDir() string {
	return filepath.Join(testDataDir(), "content")
}

func ContentPath(relativePath string) string {
	return filepath.Join(ContentDir(), relativePath)
}

func CopyContentToTempDir() (string, error) {
	dir, err := os.MkdirTemp("", "test*")
	if err != nil {
		return "", xerrors.Errorf("failed to create temp directory: %w", dir)
	}
	err = os.CopyFS(dir, os.DirFS(ContentDir()))
	if err != nil {
		return "", xerrors.Errorf("failed to copy files into temp directory: %w", dir)
	}
	return dir, nil
}

func TemplatesDir() string {
	return filepath.Join(testDataDir(), "templates")
}

func PublicDir() string {
	return filepath.Join(testDataDir(), "public")
}

func ExpectedHtmlDir() string {
	return filepath.Join(testDataDir(), "expected_html")
}
