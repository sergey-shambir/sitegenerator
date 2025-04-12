package data

import (
	"sitegenerator/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadSources(t *testing.T) {
	dir := testDataDir()
	sources, err := ReadSources(dir, []string{".yaml"})

	assert.NoError(t, err)

	expectedImages := []string{"images/1.gif", "images/2.jpg", "images/3.webp", "images/4.png"}
	expectedMarkdown := []string{"markdown-demo.md"}

	assert.Equal(t, dir, sources.Root())
	assert.Equal(t, expectedMarkdown, sources.ListFiles(app.Markdown))
	assert.Equal(t, expectedImages, sources.ListFiles(app.Image))
	assert.Equal(t, []string(nil), sources.ListFiles(app.Unknown))
}
