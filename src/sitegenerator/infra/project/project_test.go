package project

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"sitegenerator/app"
	"sitegenerator/infra/testdata"
)

func TestLoadProject(t *testing.T) {
	project, err := LoadProject(testdata.ContentDir(), testdata.ContentPath("index.yaml"), testdata.ContentPath("sitegenerator.cache.json"))
	assert.NoError(t, err)

	err = project.AddArticles([]string{
		"internal/markdown-demo.md",
		"internal/notes.md",
		"drafts/acceptance-testing.md",
	})
	assert.NoError(t, err)

	sections := project.ListSections()

	assert.Equal(t, &app.SectionPageDetails{
		Url:       "/internal.html",
		Title:     "Внутренние статьи",
		IsVisible: true,
		Pages:     testdata.ExpectedArticlePages("internal/markdown-demo.html", "internal/notes.html"),
	}, sections[0])

	assert.Equal(t, &app.SectionPageDetails{
		Url:       "/drafts.html",
		Title:     "Черновики",
		IsVisible: false,
		Pages:     testdata.ExpectedArticlePages("drafts/acceptance-testing.html"),
	}, sections[1])
}

func TestEditProject(t *testing.T) {
	project, err := LoadProject(testdata.ContentDir(), testdata.ContentPath("index.yaml"), testdata.ContentPath("sitegenerator.cache.json"))
	assert.NoError(t, err)

	err = project.AddArticles([]string{
		"drafts/acceptance-testing.md",
		"drafts/testing-pyramid.md",
	})
	assert.NoError(t, err)

	sections := project.ListSections()
	assert.Len(t, sections, 2)

	assert.Equal(t, &app.SectionPageDetails{
		Url:       "/drafts.html",
		Title:     "Черновики",
		IsVisible: false,
		Pages:     testdata.ExpectedArticlePages("drafts/acceptance-testing.html", "drafts/testing-pyramid.html"),
	}, sections[1])

	err = project.AddArticles([]string{
		"golang/unicode.md",
		"golang/error-handling.md",
	})
	assert.NoError(t, err)

	sections = project.ListSections()
	assert.Len(t, sections, 3)
	assert.Equal(t, &app.SectionPageDetails{
		Url:       "/golang.html",
		Title:     "golang",
		IsVisible: true,
		Pages:     testdata.ExpectedArticlePages("golang/unicode.html", "golang/error-handling.html"),
	}, sections[2])
}

func TestSaveProject(t *testing.T) {
	tempDir, err := testdata.CopyContentToTempDir()
	assert.NoError(t, err)

	project, err := LoadProject(tempDir, filepath.Join(tempDir, "index.yaml"), filepath.Join(tempDir, "sitegenerator.cache.json"))
	assert.NoError(t, err)

	err = project.AddArticles([]string{
		"internal/markdown-demo.md",
		"internal/notes.md",
		"drafts/acceptance-testing.md",
		"drafts/testing-pyramid.md",
		"golang/unicode.md",
		"golang/error-handling.md",
	})
	assert.NoError(t, err)

	err = project.Save()
	assert.NoError(t, err)

	// Перезагрузка сохранённого проекта.
	project, err = LoadProject(tempDir, filepath.Join(tempDir, "index.yaml"), filepath.Join(tempDir, "sitegenerator.cache.json"))
	assert.NoError(t, err)

	sections := project.ListSections()
	assert.Equal(t, &app.SectionPageDetails{
		Url:       "/internal.html",
		Title:     "Внутренние статьи",
		IsVisible: true,
		Pages:     testdata.ExpectedArticlePages("internal/markdown-demo.html", "internal/notes.html"),
	}, sections[0])
	assert.Len(t, sections, 3)
	assert.Equal(t, &app.SectionPageDetails{
		Url:       "/drafts.html",
		Title:     "Черновики",
		IsVisible: false,
		Pages:     testdata.ExpectedArticlePages("drafts/acceptance-testing.html", "drafts/testing-pyramid.html"),
	}, sections[1])
	sections = project.ListSections()
	assert.Equal(t, &app.SectionPageDetails{
		Url:       "/golang.html",
		Title:     "golang",
		IsVisible: true,
		Pages:     testdata.ExpectedArticlePages("golang/unicode.html", "golang/error-handling.html"),
	}, sections[2])
}
