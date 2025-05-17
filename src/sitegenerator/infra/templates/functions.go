package templates

import (
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/xerrors"
)

// Создаёт список функций, доступных шаблонам.
// Функции могут вызывать panic(), поэтому рендеринг должен использовать recover().
func createFuncMap(callbacks FuncCallbacks) template.FuncMap {
	return template.FuncMap{
		"join": join,
		"addAssetHash": func(urlPath string) string {
			return addAssetHash(callbacks, urlPath)
		},
	}
}

// Соединяет строки, используя указанный разделитель.
func join(texts []string, separator string) string {
	return strings.Join(texts, separator)
}

// Добавляет к URL Path ресурса его hash-сумму,
// чтобы при изменении ресурса кэш браузера не возвращал его старую версию.
func addAssetHash(callbacks FuncCallbacks, urlPath string) string {
	path := filepath.Join(callbacks.OutputRoot(), urlPathToPath(urlPath))
	hash, err := computeFileHash(path)
	if err != nil {
		panic(err)
	}

	return callbacks.AddAssetHash(urlPath, hash)
}

func urlPathToPath(urlPath string) string {
	urlParts := strings.Split(strings.TrimPrefix(urlPath, "/"), "/")
	return filepath.Join(urlParts...)
}

func computeFileHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", xerrors.Errorf("could not open file at %q: %w", path, err)
	}
	defer file.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", xerrors.Errorf("could not compute MD5 hash for file at %q: %w", path, err)
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
