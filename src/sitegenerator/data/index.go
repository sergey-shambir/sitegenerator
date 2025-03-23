package data

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

type PagesIndex struct {
	data []PagesIndexItem
}

type PagesSection struct {
	Key     string
	Title   string
	Visible bool
	Files   []string
}

type PagesIndexItem struct {
	Key   string
	Value *PagesIndexSectionData
}

func (pi *PagesIndex) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.MappingNode {
		return fmt.Errorf("Unexpected YAML node %s", n.ShortTag())
	}

	sectionsMap := make([]PagesIndexItem, 0, len(n.Content)/2)
	for i := 0; i < len(n.Content); i += 2 {
		key := n.Content[i].Value
		section := new(PagesIndexSectionData)
		err := n.Content[i+1].Decode(&section)
		if err != nil {
			return err
		}
		sectionsMap = append(sectionsMap, PagesIndexItem{
			Key:   key,
			Value: section,
		})
	}

	pi.data = sectionsMap
	return nil
}

type PagesIndexSectionData struct {
	Title   string   `yaml:"title"`
	Visible *bool    `yaml:"visible"`
	Files   []string `yaml:"files"`
}

/*
 * Добавляет страницы, заданные списком путей.
 * Пути отсчитываются от каталога контента сайта.
 */
func (pi *PagesIndex) AddPages(paths []string) {
	for _, path := range paths {
		pi.AddPage(path)
	}
}

/*
 * Добавляет страницу по заданному пути файла.
 * Путь отсчитывается от каталога контента сайта.
 */
func (pi *PagesIndex) AddPage(path string) {
	dir, filename := filepath.Split(path)
	dir = strings.TrimSuffix(dir, string(filepath.Separator))
	section := pi.findOrCreateSection(dir)
	if !slices.Contains(section.Files, filename) {
		section.Files = append(section.Files, filename)
	}
}

func (pi *PagesIndex) findOrCreateSection(dir string) *PagesIndexSectionData {
	for _, section := range pi.data {
		if section.Key == dir {
			return section.Value
		}
	}
	section := &PagesIndexSectionData{
		Title: dir,
		Files: nil,
	}
	pi.data = append(pi.data, PagesIndexItem{
		Key:   dir,
		Value: section,
	})
	return section
}

func (pi *PagesIndex) ListSections() []PagesSection {
	results := make([]PagesSection, len(pi.data))
	for i, kv := range pi.data {
		results[i] = PagesSection{
			Key:     kv.Key,
			Title:   kv.Value.Title,
			Visible: kv.Value.Visible == nil || *kv.Value.Visible,
			Files:   kv.Value.Files,
		}
	}

	return results
}

func LoadPagesIndex(path string) (*PagesIndex, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return loadPagesIndexImpl(data)
}

func loadPagesIndexImpl(data []byte) (*PagesIndex, error) {
	var pagesIndex PagesIndex
	err := yaml.Unmarshal(data, &pagesIndex)
	if err != nil {
		return nil, err
	}

	return &pagesIndex, nil
}

func SavePagesIndex(path string, pagesIndex *PagesIndex) error {
	data, err := savePagesIndexImpl(pagesIndex)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func savePagesIndexImpl(pagesIndex *PagesIndex) ([]byte, error) {
	return yaml.Marshal(pagesIndex.data)
}
