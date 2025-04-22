package app

type MarkdownConverter interface {
	ConvertToHtml(path string) ([]byte, error)
}
