package app

type PagesIndex interface {
	Save() error

	/*
	 * Добавляет страницы, заданные списком путей.
	 * Пути отсчитываются от каталога контента сайта.
	 */
	AddArticles(paths []string)

	/**
	 * Возвращает раздел, к которому относится статья по заданному пути.
	 */
	GetArticleSection(path string) *SectionPageDetails

	/**
	 * Возвращает список всех разделов.
	 */
	ListSections() []*SectionPageDetails
}
