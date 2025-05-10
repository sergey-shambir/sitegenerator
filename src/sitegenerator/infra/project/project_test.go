package project

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"sitegenerator/app"
	"sitegenerator/infra/testdata"
)

func TestLoadProject(t *testing.T) {
	project, err := LoadProject(testdata.RootDir(), testdata.AbsPath("index.yaml"), testdata.AbsPath("sitegenerator.cache.json"))
	assert.NoError(t, err)

	sections := project.ListSections()

	assert.Equal(t, &app.SectionPageData{
		Path:    "internal",
		Title:   "Внутренние статьи",
		Visible: true,
		Files:   []string{"markdown-demo.md", "notes.md"},
	}, sections[0])

	assert.Equal(t, &app.SectionPageData{
		Path:    "drafts",
		Title:   "Черновики",
		Visible: false,
		Files:   []string{"acceptance-testing.md"},
	}, sections[1])
}

func TestEditProject(t *testing.T) {
	project, err := LoadProject(testdata.RootDir(), testdata.AbsPath("index.yaml"), testdata.AbsPath("sitegenerator.cache.json"))
	assert.NoError(t, err)

	err = project.AddArticles([]string{"drafts/acceptance-testing.md", "drafts/testing-pyramid.md"})
	assert.NoError(t, err)

	sections := project.ListSections()
	assert.Len(t, sections, 2)
	assert.Equal(t, &app.SectionPageData{
		Path:    "drafts",
		Title:   "Черновики",
		Visible: false,
		Files:   []string{"acceptance-testing.md", "testing-pyramid.md"},
	}, sections[1])

	err = project.AddArticles([]string{"golang/unicode.md", "golang/error-handling.md"})
	assert.NoError(t, err)

	sections = project.ListSections()
	assert.Len(t, sections, 3)
	assert.Equal(t, &app.SectionPageData{
		Path:    "golang",
		Title:   "golang",
		Visible: true,
		Files:   []string{"unicode.md", "error-handling.md"},
	}, sections[2])
}

func TestSaveProject(t *testing.T) {
	tempDir, err := testdata.CopyToTempDir()
	assert.NoError(t, err)

	project, err := LoadProject(tempDir, filepath.Join(tempDir, "index.yaml"), filepath.Join(tempDir, "sitegenerator.cache.json"))
	assert.NoError(t, err)

	err = project.AddArticles([]string{"drafts/acceptance-testing.md", "drafts/testing-pyramid.md", "golang/unicode.md", "golang/error-handling.md"})
	assert.NoError(t, err)

	err = project.Save()
	assert.NoError(t, err)

	// Перезагрузка сохранённого проекта.
	project, err = LoadProject(tempDir, filepath.Join(tempDir, "index.yaml"), filepath.Join(tempDir, "sitegenerator.cache.json"))
	assert.NoError(t, err)

	sections := project.ListSections()
	assert.Equal(t, &app.SectionPageData{
		Path:    "internal",
		Title:   "Внутренние статьи",
		Visible: true,
		Files:   []string{"markdown-demo.md", "notes.md"},
	}, sections[0])
	assert.Len(t, sections, 3)
	assert.Equal(t, &app.SectionPageData{
		Path:    "drafts",
		Title:   "Черновики",
		Visible: false,
		Files:   []string{"acceptance-testing.md", "testing-pyramid.md"},
	}, sections[1])
	sections = project.ListSections()
	assert.Equal(t, &app.SectionPageData{
		Path:    "golang",
		Title:   "golang",
		Visible: true,
		Files:   []string{"unicode.md", "error-handling.md"},
	}, sections[2])
}
