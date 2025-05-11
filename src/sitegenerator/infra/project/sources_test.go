package project

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sitegenerator/app"
	"sitegenerator/infra/testdata"
)

func TestReadSources(t *testing.T) {
	dir := testdata.ContentDir()
	sources, err := ReadSources(dir, []string{".yaml"})

	assert.NoError(t, err)

	expectedImages := []string{"images/1.gif", "images/2.jpg", "images/3.webp", "images/4.png"}
	expectedMarkdown := []string{"drafts/acceptance-testing.md", "drafts/testing-pyramid.md", "golang/error-handling.md", "golang/unicode.md", "internal/markdown-demo.md", "internal/notes.md"}
	expectedSass := []string{"main.scss"}

	assert.Equal(t, dir, sources.Root())
	assert.Equal(t, expectedMarkdown, sources.ListFiles(app.Markdown))
	assert.Equal(t, expectedImages, sources.ListFiles(app.Image))
	assert.Equal(t, expectedSass, sources.ListFiles(app.Sass))
	assert.Equal(t, []string(nil), sources.ListFiles(app.Unknown))
}
