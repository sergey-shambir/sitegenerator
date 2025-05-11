package templates

import (
	"html/template"
	"strings"
)

func createFuncMap() template.FuncMap {
	return template.FuncMap{
		"join":         join,
		"addAssetHash": addAssetHash,
	}
}

// Соединяет строки, используя указанный разделитель.
func join(texts []string, separator string) string {
	return strings.Join(texts, separator)
}

// Добавляет к URL Path ресурса его hash-сумму,
// чтобы при изменении ресурса кэш браузера не возвращал его старую версию.
func addAssetHash(urlPath string) string {
	return urlPath
}
