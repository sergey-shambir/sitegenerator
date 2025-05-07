package app

type Converter interface {
	ConvertMarkdownToHtml(path string) ([]byte, error)

	ConvertSassToCss(path string) ([]byte, error)
}
