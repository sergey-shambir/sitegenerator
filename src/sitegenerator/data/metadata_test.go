package data

import (
	"sitegenerator/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPageMetadata(t *testing.T) {
	cache, err := LoadGeneratorCache(testDataDir(), testDataAbsPath("sitegenerator.cache.json"))

	assert.NoError(t, err)
	metadata, err := cache.GetPageMetadata("markdown-demo.md")
	assert.NoError(t, err)
	assert.Equal(t, &app.PageMetadata{
		Title:       "Демонстрация возможностей Markdown",
		Description: "Тестовая страница",
		Category:    "cheatsheet",
		Keywords:    []string{"markdown", "markdown-it"},
	}, metadata)
}

func TestLoadPageMetadata(t *testing.T) {
	metadata, err := ParsePageMetadata(testDataAbsPath("markdown-demo.md"))
	assert.NoError(t, err)
	assert.Equal(t, &app.PageMetadata{
		Title:       "Демонстрация возможностей Markdown",
		Description: "Тестовая страница",
		Category:    "cheatsheet",
		Keywords:    []string{"markdown", "markdown-it"},
	}, metadata)
}
