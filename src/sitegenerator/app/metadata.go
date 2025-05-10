package app

type GeneratorCache interface {
	AddArticles(paths []string) error
	GetArticleMetadata(path string) *ArticleMetadata
	Save() error
}
