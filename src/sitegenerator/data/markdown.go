package data

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/xerrors"
)

type markdownConverter struct {
	converterRoot   string
	converterScript string
}

func NewMarkdownConverter(converterRoot string) (*markdownConverter, error) {
	c := &markdownConverter{
		converterRoot:   converterRoot,
		converterScript: "index.js",
	}

	err := c.checkConverterExists()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *markdownConverter) ConvertToHtml(path string) ([]byte, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command("node", c.converterScript, path)
	cmd.Dir = c.converterRoot
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, xerrors.Errorf("failed to convert markdown file %s: %w\n%s", path, err, stderr.String())
	}

	return stdout.Bytes(), nil
}

func (c *markdownConverter) checkConverterExists() error {
	scriptPath := filepath.Join(c.converterRoot, c.converterScript)
	_, err := os.Stat(scriptPath)
	if err != nil {
		return xerrors.Errorf("no converter script %s at %s: %w", c.converterScript, c.converterRoot, err)
	}
	return nil
}
