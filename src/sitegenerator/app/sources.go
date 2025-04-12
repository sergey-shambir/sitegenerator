package app

type SourceType int

const (
	Unknown SourceType = iota
	Markdown
	Image
	StyleSheet
	JavaScript
)

type Sources interface {
	Root() string
	ListFiles(t SourceType) []string
}
