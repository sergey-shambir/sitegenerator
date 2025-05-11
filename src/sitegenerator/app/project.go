package app

/*
 * Данные статьи, извлечённые из front matter в markdown файле
 */
type ArticleMetadata struct {
	Title       string
	Description string
	Category    string
	Keywords    []string
}

type Project interface {
	AddArticles(paths []string) error
	Save() error

	GetArticleSection(path string) *SectionPageDetails
	ListSections() []*SectionPageDetails
}
