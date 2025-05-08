package project

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sitegenerator/app"
	"sitegenerator/data/testdata"
)

func TestGetPageMetadata(t *testing.T) {
	cache, err := LoadGeneratorCache(testdata.RootDir(), testdata.AbsPath("sitegenerator.cache.json"))

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
	metadata, err := ParsePageMetadata(testdata.AbsPath("markdown-demo.md"))
	assert.NoError(t, err)
	assert.Equal(t, &app.PageMetadata{
		Title:       "Демонстрация возможностей Markdown",
		Description: "Тестовая страница",
		Category:    "cheatsheet",
		Keywords:    []string{"markdown", "markdown-it"},
	}, metadata)
}
