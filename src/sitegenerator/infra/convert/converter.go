package convert

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/xerrors"

	"sitegenerator/app"
)

const (
	markdownFormat = "markdown"
	sassFormat     = "sass"
)

type nodeConverter struct {
	converterRoot   string
	converterScript string
}

func NewConverter(converterRoot string) (app.Converter, error) {
	c := &nodeConverter{
		converterRoot:   converterRoot,
		converterScript: "index.js",
	}

	err := c.checkConverterExists()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *nodeConverter) ConvertMarkdownToHtml(path string) ([]byte, error) {
	return c.convert(markdownFormat, path)
}

func (c *nodeConverter) ConvertSassToCss(path string) ([]byte, error) {
	return c.convert(sassFormat, path)
}

func (c *nodeConverter) convert(format, path string) ([]byte, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command("node", c.converterScript, format, path)
	cmd.Dir = c.converterRoot
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, xerrors.Errorf("failed to convert %s file '%s': %w\n%s", format, path, err, stderr.String())
	}

	return stdout.Bytes(), nil
}

func (c *nodeConverter) checkConverterExists() error {
	scriptPath := filepath.Join(c.converterRoot, c.converterScript)
	_, err := os.Stat(scriptPath)
	if err != nil {
		return xerrors.Errorf("no converter script %s at '%s': %w", c.converterScript, c.converterRoot, err)
	}
	return nil
}
