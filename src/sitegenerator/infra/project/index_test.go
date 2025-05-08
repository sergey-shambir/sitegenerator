package project

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sitegenerator/infra/testdata"
)

func TestLoadIndex(t *testing.T) {
	pagesIndex, err := LoadPagesIndex(testdata.AbsPath("index.yaml"))
	assert.NoError(t, err)

	sections := pagesIndex.ListSections()
	assert.Len(t, sections, 2)

	assert.Equal(t, PagesSection{
		Key:     "internal",
		Title:   "Внутренние статьи",
		Visible: true,
		Files:   []string{"markdown-demo.md", "notes.md"},
	}, sections[0])

	assert.Equal(t, PagesSection{
		Key:     "drafts",
		Title:   "Черновики",
		Visible: false,
		Files:   []string{"acceptance-testing.md"},
	}, sections[1])
}

func TestAddPages(t *testing.T) {
	pagesIndex, err := LoadPagesIndex(testdata.AbsPath("index.yaml"))
	assert.NoError(t, err)

	pagesIndex.AddPages([]string{"drafts/acceptance-testing.md", "drafts/testing-pyramid.md"})

	sections := pagesIndex.ListSections()
	assert.Len(t, sections, 2)
	assert.Equal(t, PagesSection{
		Key:     "drafts",
		Title:   "Черновики",
		Visible: false,
		Files:   []string{"acceptance-testing.md", "testing-pyramid.md"},
	}, sections[1])

	pagesIndex.AddPages([]string{"golang/unicode.md", "golang/error-handling.md"})

	sections = pagesIndex.ListSections()
	assert.Len(t, sections, 3)
	assert.Equal(t, PagesSection{
		Key:     "golang",
		Title:   "golang",
		Visible: true,
		Files:   []string{"unicode.md", "error-handling.md"},
	}, sections[2])
}
