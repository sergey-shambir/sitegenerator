package data

import (
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/xerrors"
)

type markdownConverter struct {
	converterRoot   string
	converterScript string
}

func NewMarkdownConverter(converterRoot string) *markdownConverter {
	return &markdownConverter{
		converterRoot:   converterRoot,
		converterScript: "index.js",
	}
}

func (c *markdownConverter) ConvertToHtml(path string) ([]byte, error) {
	err := c.checkConverterExists()
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("node", c.converterScript, path)
	cmd.Dir = c.converterRoot

	output, err := cmd.Output()
	if err != nil {
		return nil, xerrors.Errorf("failed to convert markdown file %s: %w", path, err)
	}
	return output, nil
}

func (c *markdownConverter) checkConverterExists() error {
	scriptPath := filepath.Join(c.converterRoot, c.converterScript)
	_, err := os.Stat(scriptPath)
	if err != nil {
		return xerrors.Errorf("no converter script %s at %s: %w", c.converterScript, c.converterRoot, err)
	}
	return nil
}
