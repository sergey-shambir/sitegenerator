package app

type PageMetadata struct {
	Title       string
	Description string
	Category    string
	Keywords    []string
}

type GeneratorCache interface {
	GetPageMetadata(path string) (*PageMetadata, error)
	SaveCache() error
}
