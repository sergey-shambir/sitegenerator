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

type metadataCacheEntry struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Keywords    []string `json:"keywords"`
}

func fromArticleMetadata(m *app.ArticleMetadata) *metadataCacheEntry {
	return &metadataCacheEntry{
		Title:       m.Title,
		Description: m.Description,
		Category:    m.Category,
		Keywords:    m.Keywords,
	}
}

func toArticleMetadata(m *metadataCacheEntry) *app.ArticleMetadata {
	if m == nil {
		return nil
	}
	return &app.ArticleMetadata{
		Title:       m.Title,
		Description: m.Description,
		Category:    m.Category,
		Keywords:    m.Keywords,
	}
}

type metadataCache struct {
	pagesDir  string
	cachePath string
	pages     map[string]*metadataCacheEntry
}

func loadGeneratorCache(pagesDir string, cachePath string) (*metadataCache, error) {
	cache := &metadataCache{
		pagesDir:  pagesDir,
		cachePath: cachePath,
		pages:     make(map[string]*metadataCacheEntry),
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

func (c *metadataCache) addArticles(paths []string) error {
	for _, path := range paths {
		if _, ok := c.pages[path]; ok {
			continue
		}
		metadata, err := parsePageMetadata(filepath.Join(c.pagesDir, path))
		if err != nil {
			return err
		}
		c.pages[path] = fromArticleMetadata(metadata)
	}
	return nil
}

func (c *metadataCache) getArticleMetadata(path string) *app.ArticleMetadata {
	return toArticleMetadata(c.pages[path])
}

func (c *metadataCache) save() error {
	data, err := json.MarshalIndent(c.pages, "", "    ")
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
