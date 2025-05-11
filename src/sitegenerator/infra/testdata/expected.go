package testdata

import (
	"os"
	"path/filepath"
	"sitegenerator/app"

	"golang.org/x/xerrors"
)

func ExpectedArticlePages(paths ...string) []app.UrlAndValue[*app.ArticleMetadata] {
	var result []app.UrlAndValue[*app.ArticleMetadata]
	for _, path := range paths {
		result = append(result, app.UrlAndValue[*app.ArticleMetadata]{
			Url:   "/" + path,
			Value: ExpectedMetadata(path),
		})
	}
	return result
}

func ExpectedMetadata(path string) *app.ArticleMetadata {
	switch path {
	case "internal/markdown-demo.md":
		return &app.ArticleMetadata{
			Title:       "Демонстрация возможностей Markdown",
			Description: "Тестовая страница",
			Category:    "cheatsheet",
			Keywords:    []string{"markdown", "markdown-it"},
		}
	case "internal/notes.md":
		return &app.ArticleMetadata{
			Title:       "Заметки в формате Markdown",
			Description: "Тестовая страница",
			Category:    "cheatsheet",
			Keywords:    []string{"markdown", "markdown-it"},
		}
	case "drafts/acceptance-testing.md":
		return &app.ArticleMetadata{
			Title:       "Приёмочные тесты",
			Description: "Что такое приёмочные тесты и как понять, какие тесты приёмочные",
			Category:    "tutorial",
			Keywords:    []string{"ATDD"},
		}
	case "drafts/testing-pyramid.md":
		return &app.ArticleMetadata{
			Title:       "Пирамида тестирования",
			Description: "Уровни автоматизированных тестов",
			Category:    "tutorial",
			Keywords:    []string{"ATDD"},
		}
	case "golang/unicode.md":
		return &app.ArticleMetadata{
			Title:       "Unicode в Go",
			Description: "Как работает поддержка Unicode в Go",
			Category:    "tutorial",
			Keywords:    []string{"Go"},
		}
	case "golang/error-handling.md":
		return &app.ArticleMetadata{
			Title:       "Обработка ошибок в Go",
			Description: "Как выстроить обработку ошибок в Go",
			Category:    "tutorial",
			Keywords:    []string{"Go"},
		}
	default:
		panic(xerrors.Errorf("no expected metadata for article path %q", path))
	}
}

func ExpectedHtml(filename string) string {
	path := filepath.Join(ExpectedHtmlDir(), filename)
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(xerrors.Errorf("could not to read expected HTML from file %q: %w", path, err))
	}
	return string(bytes)
}
