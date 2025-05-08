package app

type ArticleMetadata struct {
	Title       string
	Description string
	Category    string
	Keywords    []string
}

type GeneratorCache interface {
	GetArticleMetadata(path string) (*ArticleMetadata, error)
	SaveCache() error
}
