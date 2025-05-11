package project

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sitegenerator/app"
	"sitegenerator/infra/testdata"
)

func TestGetPageMetadata(t *testing.T) {
	cache, err := loadGeneratorCache(testdata.ContentDir(), testdata.ContentPath("sitegenerator.cache.json"))
	assert.NoError(t, err)

	metadata := cache.getArticleMetadata("markdown-demo.md")
	assert.Nil(t, metadata)

	err = cache.addArticles([]string{"internal/markdown-demo.md"})
	assert.NoError(t, err)

	metadata = cache.getArticleMetadata("internal/markdown-demo.md")
	assert.Equal(t, &app.ArticleMetadata{
		Title:       "Демонстрация возможностей Markdown",
		Description: "Тестовая страница",
		Category:    "cheatsheet",
		Keywords:    []string{"markdown", "markdown-it"},
	}, metadata)
}

func TestLoadPageMetadata(t *testing.T) {
	metadata, err := parsePageMetadata(testdata.ContentPath("internal/markdown-demo.md"))
	assert.NoError(t, err)
	assert.Equal(t, &app.ArticleMetadata{
		Title:       "Демонстрация возможностей Markdown",
		Description: "Тестовая страница",
		Category:    "cheatsheet",
		Keywords:    []string{"markdown", "markdown-it"},
	}, metadata)
}
