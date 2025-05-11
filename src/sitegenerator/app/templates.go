package app

// UrlAndValue - связь между URL и ассоциированным значением.
type UrlAndValue[V any] struct {
	Url   string
	Value V
}

// ArticlePageDetails - данные для генерации страницы статьи.
type ArticlePageDetails struct {
	IsVisible bool
	Meta      *ArticleMetadata
	Content   []byte
}

// SectionPageDetails - данные для генерации страницы раздела.
type SectionPageDetails struct {
	Url       string
	IsVisible bool
	Title     string
	Pages     []UrlAndValue[*ArticleMetadata]
}

// IndexPageData - данные для генерации главной страницы сайта.
type IndexPageData struct {
	Title    string
	Sections []UrlAndValue[string]
}

// SiteTemplates - инструмент для генерации HTML-страниц сайта.
type SiteTemplates interface {
	GenerateArticlePage(d ArticlePageDetails) ([]byte, error)
	GenerateSectionPage(d SectionPageDetails) ([]byte, error)
	GenerateIndexPage(d IndexPageData) ([]byte, error)
}
