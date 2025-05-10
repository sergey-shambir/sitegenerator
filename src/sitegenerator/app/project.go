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

type ArticlePageData struct {
	Path    string
	Visible bool
	Meta    ArticleMetadata
}

type SectionPageData struct {
	Path    string
	Title   string
	Visible bool
	Files   []string
}

type Project interface {
	AddArticles(paths []string) error
	Save() error

	GetArticleSection(path string) *SectionPageData
	ListSections() []*SectionPageData
}
