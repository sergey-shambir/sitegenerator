package app

type SourceType int

const (
	Unknown SourceType = iota
	Markdown
	Sass
	Image
	StyleSheet
	JavaScript
)

/**
 * Предоставляет доступ к исходным файлам, из которых генерируется сайт.
 */
type Sources interface {
	Root() string
	ListFiles(t SourceType) []string
}
