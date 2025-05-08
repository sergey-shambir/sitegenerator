package app

type PagesSection struct {
	Key     string
	Title   string
	Visible bool
	Files   []string
}

type PagesIndex interface {
	Save(path string) error

	/*
	 * Добавляет страницы, заданные списком путей.
	 * Пути отсчитываются от каталога контента сайта.
	 */
	AddPages(paths []string)

	/*
	 * Добавляет страницу по заданному пути файла.
	 * Путь отсчитывается от каталога контента сайта.
	 */
	AddPage(path string)

	ListSections() []PagesSection
}
