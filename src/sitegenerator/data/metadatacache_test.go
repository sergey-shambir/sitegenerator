package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPageMetadata(t *testing.T) {
	cache, err := LoadPagesMetadataCache(testDataAbsPath("sitegenerator.cache.json"))

	assert.NoError(t, err)
	metadata, err := cache.GetPageMetadata(testDataAbsPath("markdown-demo.md"))
	assert.NoError(t, err)
	assert.Equal(t, &PageMetadata{
		Title:       "Демонстрация возможностей Markdown",
		Description: "Тестовая страница",
		Category:    "cheatsheet",
		Keywords:    []string{"markdown", "markdown-it"},
	}, metadata)
}
