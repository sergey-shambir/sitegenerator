package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPageMetadata(t *testing.T) {
	metadata, err := ParsePageMetadata(testDataAbsPath("markdown-demo.md"))
	assert.NoError(t, err)
	assert.Equal(t, &PageMetadata{
		Title:       "Демонстрация возможностей Markdown",
		Description: "Тестовая страница",
		Category:    "cheatsheet",
		Keywords:    []string{"markdown", "markdown-it"},
	}, metadata)
}
