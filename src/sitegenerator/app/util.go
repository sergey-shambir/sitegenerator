package app

import (
	"path/filepath"
	"strings"
)

const (
	IndexHtmlPath = "index.html"
	HtmlExt       = ".html"
	CssExt        = ".css"
)

func ReplaceFileExtension(path string, newExt string) string {
	basePath, _ := strings.CutSuffix(path, filepath.Ext(path))
	return basePath + newExt
}

func UrlPathToPath(urlPath string) string {
	urlParts := strings.Split(strings.TrimPrefix(urlPath, "/"), "/")
	return filepath.Join(urlParts...)
}
