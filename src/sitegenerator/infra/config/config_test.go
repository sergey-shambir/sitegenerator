package config

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"sitegenerator/infra/testdata"
)

func TestLoadDefaultConfig(t *testing.T) {
	testDataDir := testdata.ContentDir()
	config, err := ReadConfig(testdata.ContentPath("sitegenerator-default.yaml"))

	assert.NoError(t, err)
	assert.Equal(t, &Config{
		SourceDir:    filepath.Join(testDataDir, "content"),
		TargetDir:    filepath.Join(testDataDir, "public"),
		TemplatesDir: filepath.Join(testDataDir, "templates"),
	}, config)
}

func TestLoadConfig(t *testing.T) {
	testDataDir := testdata.ContentDir()
	config, err := ReadConfig(testdata.ContentPath("sitegenerator.yaml"))

	assert.NoError(t, err)
	assert.Equal(t, &Config{
		SourceDir:            filepath.Join(testDataDir, "input"),
		TargetDir:            filepath.Join(testDataDir, "output"),
		TemplatesDir:         filepath.Join(testDataDir, "html_templates"),
		IgnoreFileExtensions: []string{".json", ".yaml"},
	}, config)
}
