package project

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	yaml "gopkg.in/yaml.v3"

	"sitegenerator/app"
)

type pagesIndex struct {
	data []pagesIndexItem
}

type pagesIndexItem struct {
	Key   string
	Value *pagesIndexSectionData
}

type pagesIndexSectionData struct {
	Title   string   `yaml:"title"`
	Visible *bool    `yaml:"visible"`
	Files   []string `yaml:"files"`
}

func LoadPagesIndex(path string) (app.PagesIndex, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pagesIndex pagesIndex
	err = yaml.Unmarshal(data, &pagesIndex)
	if err != nil {
		return nil, err
	}

	return &pagesIndex, nil
}

func (pi *pagesIndex) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.MappingNode {
		return fmt.Errorf("Unexpected YAML node %s", n.ShortTag())
	}

	sectionsMap := make([]pagesIndexItem, 0, len(n.Content)/2)
	for i := 0; i < len(n.Content); i += 2 {
		key := n.Content[i].Value
		section := new(pagesIndexSectionData)
		err := n.Content[i+1].Decode(&section)
		if err != nil {
			return err
		}
		sectionsMap = append(sectionsMap, pagesIndexItem{
			Key:   key,
			Value: section,
		})
	}

	pi.data = sectionsMap
	return nil
}

func (pi *pagesIndex) Save(path string) error {
	data, err := yaml.Marshal(pi.data)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func (pi *pagesIndex) AddPages(paths []string) {
	for _, path := range paths {
		pi.AddPage(path)
	}
}

func (pi *pagesIndex) AddPage(path string) {
	dir, filename := filepath.Split(path)
	dir = strings.TrimSuffix(dir, string(filepath.Separator))
	section := pi.findOrCreateSection(dir)
	if !slices.Contains(section.Files, filename) {
		section.Files = append(section.Files, filename)
	}
}

func (pi *pagesIndex) findOrCreateSection(dir string) *pagesIndexSectionData {
	for _, section := range pi.data {
		if section.Key == dir {
			return section.Value
		}
	}
	section := &pagesIndexSectionData{
		Title: dir,
		Files: nil,
	}
	pi.data = append(pi.data, pagesIndexItem{
		Key:   dir,
		Value: section,
	})
	return section
}

func (pi *pagesIndex) ListSections() []app.PagesSection {
	results := make([]app.PagesSection, len(pi.data))
	for i, kv := range pi.data {
		results[i] = app.PagesSection{
			Key:     kv.Key,
			Title:   kv.Value.Title,
			Visible: kv.Value.Visible == nil || *kv.Value.Visible,
			Files:   kv.Value.Files,
		}
	}

	return results
}
