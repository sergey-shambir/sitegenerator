package app

import (
	"path/filepath"
	"strings"
)

func replaceFileExtension(path string, newExt string) string {
	basePath, _ := strings.CutSuffix(path, filepath.Ext(path))
	return basePath + newExt
}
