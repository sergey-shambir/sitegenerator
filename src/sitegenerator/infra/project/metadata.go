package project

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/xerrors"
	yaml "gopkg.in/yaml.v3"

	"sitegenerator/app"
)

const metadataMarker = "---"

type generatorCache struct {
	pagesDir  string
	cachePath string
	pages     map[string]*app.ArticleMetadata
}

func loadGeneratorCache(pagesDir string, cachePath string) (*generatorCache, error) {
	cache := &generatorCache{
		pagesDir:  pagesDir,
		cachePath: cachePath,
		pages:     make(map[string]*app.ArticleMetadata),
	}
	data, err := os.ReadFile(cachePath)

	if err != nil {
		if os.IsNotExist(err) {
			return cache, nil
		}
		return nil, xerrors.Errorf("failed to read metadata cache: %w", err)
	}
	err = json.Unmarshal(data, &cache.pages)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse metadata cache: %w", err)
	}
	return cache, nil
}

func (c *generatorCache) addArticles(paths []string) error {
	for _, path := range paths {
		if _, ok := c.pages[path]; ok {
			continue
		}
		metadata, err := parsePageMetadata(filepath.Join(c.pagesDir, path))
		if err != nil {
			return err
		}
		c.pages[path] = metadata
	}
	return nil
}

func (c *generatorCache) getArticleMetadata(path string) *app.ArticleMetadata {
	return c.pages[path]
}

func (c *generatorCache) save() error {
	data, err := json.Marshal(c.pages)
	if err != nil {
		return xerrors.Errorf("failed to format metadata cache: %w", err)
	}
	err = os.WriteFile(c.cachePath, data, 0644)
	if err != nil {
		return xerrors.Errorf("failed to save metadata cache: %w", err)
	}
	return nil
}

func parsePageMetadata(path string) (*app.ArticleMetadata, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, xerrors.Errorf("failed to open file '%s': %w", path, err)
	}
	contents, err := readPageMetadataYaml(file)
	if err != nil {
		return nil, xerrors.Errorf("failed to read '%s' metadata: %w", path, err)
	}

	var metadata app.ArticleMetadata
	err = yaml.Unmarshal([]byte(contents), &metadata)

	return &metadata, err
}

func readPageMetadataYaml(file io.Reader) (string, error) {
	scanner := bufio.NewScanner(file)
	var sb strings.Builder
	isStarted := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == metadataMarker {
			if isStarted {
				return sb.String(), nil
			}
			isStarted = true
		} else {
			sb.WriteString(line)
			sb.WriteRune('\n')
		}
	}

	// NOTE: блок метаданных не был закрыт
	missingHint := "opening marker"
	if isStarted {
		missingHint = "closing marker"
	}

	return "", fmt.Errorf("No YAML %s %s up to file end", missingHint, metadataMarker)
}
