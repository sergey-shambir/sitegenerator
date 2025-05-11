package project

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sitegenerator/infra/testdata"
)

func TestLoadIndex(t *testing.T) {
	pagesIndex, err := loadPagesIndex(testdata.ContentPath("index.yaml"))
	assert.NoError(t, err)

	sections := pagesIndex.listSections()
	assert.Len(t, sections, 2)

	assert.Equal(t, &pagesIndexItem{
		Key:     "internal",
		Title:   "Внутренние статьи",
		Visible: true,
		Files:   []string{"markdown-demo.md", "notes.md"},
	}, sections[0])

	assert.Equal(t, &pagesIndexItem{
		Key:     "drafts",
		Title:   "Черновики",
		Visible: false,
		Files:   []string{"acceptance-testing.md"},
	}, sections[1])
}

func TestAddPages(t *testing.T) {
	pagesIndex, err := loadPagesIndex(testdata.ContentPath("index.yaml"))
	assert.NoError(t, err)

	pagesIndex.addArticles([]string{"drafts/acceptance-testing.md", "drafts/testing-pyramid.md"})

	sections := pagesIndex.listSections()
	assert.Len(t, sections, 2)
	assert.Equal(t, &pagesIndexItem{
		Key:     "drafts",
		Title:   "Черновики",
		Visible: false,
		Files:   []string{"acceptance-testing.md", "testing-pyramid.md"},
	}, sections[1])

	pagesIndex.addArticles([]string{"golang/unicode.md", "golang/error-handling.md"})

	sections = pagesIndex.listSections()
	assert.Len(t, sections, 3)
	assert.Equal(t, &pagesIndexItem{
		Key:     "golang",
		Title:   "golang",
		Visible: true,
		Files:   []string{"unicode.md", "error-handling.md"},
	}, sections[2])
}

// TODO: написать тест на сохранение кэша
