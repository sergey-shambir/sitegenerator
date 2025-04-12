package data

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testDataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "test_data")
}

func testDataAbsPath(relativePath string) string {
	return filepath.Join(testDataDir(), relativePath)
}

func TestLoadDefaultConfig(t *testing.T) {
	testDataDir := testDataDir()
	config, err := ReadConfig(testDataAbsPath("sitegenerator-default.yaml"))

	assert.NoError(t, err)
	assert.Equal(t, &Config{
		SourceDir:    filepath.Join(testDataDir, "content"),
		TargetDir:    filepath.Join(testDataDir, "public"),
		TemplatesDir: filepath.Join(testDataDir, "templates"),
	}, config)
}

func TestLoadConfig(t *testing.T) {
	testDataDir := testDataDir()
	config, err := ReadConfig(testDataAbsPath("sitegenerator.yaml"))

	assert.NoError(t, err)
	assert.Equal(t, &Config{
		SourceDir:            filepath.Join(testDataDir, "input"),
		TargetDir:            filepath.Join(testDataDir, "output"),
		TemplatesDir:         filepath.Join(testDataDir, "html_templates"),
		IgnoreFileExtensions: []string{".json", ".yaml"},
	}, config)
}
