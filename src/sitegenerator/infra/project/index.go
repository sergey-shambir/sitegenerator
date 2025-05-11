package project

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"golang.org/x/xerrors"
	yaml "gopkg.in/yaml.v3"
)

type pagesIndex struct {
	path  string
	items []*pagesIndexItem
}

type pagesIndexItem struct {
	Key     string   `yaml:"-"`
	Title   string   `yaml:"title"`
	Visible bool     `yaml:"visible"`
	Files   []string `yaml:"files"`
}

type pagesIndexData struct {
	items []*pagesIndexItem
}

func (pi *pagesIndexData) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.MappingNode {
		return fmt.Errorf("Unexpected YAML node %s", n.ShortTag())
	}

	sectionsMap := make([]*pagesIndexItem, 0, len(n.Content)/2)
	for i := 0; i < len(n.Content); i += 2 {
		key := n.Content[i].Value
		section := &pagesIndexItem{
			Key:     key,
			Visible: true,
		}
		err := n.Content[i+1].Decode(&section)
		if err != nil {
			return err
		}
		sectionsMap = append(sectionsMap, section)
	}

	pi.items = sectionsMap
	return nil
}

func (pi *pagesIndexData) MarshalYAML() (any, error) {
	node := &yaml.Node{
		Kind: yaml.MappingNode,
	}
	for _, item := range pi.items {
		keyNode := &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: item.Key,
		}
		valueNode := &yaml.Node{}
		if err := valueNode.Encode(item); err != nil {
			return nil, xerrors.Errorf("could not encode to YAML property %s: %w", item.Key, err)
		}

		node.Content = append(node.Content, keyNode, valueNode)
	}
	return node, nil
}

func loadPagesIndex(path string) (*pagesIndex, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data pagesIndexData
	err = yaml.Unmarshal(contents, &data)
	if err != nil {
		return nil, err
	}

	return &pagesIndex{
		path:  path,
		items: data.items,
	}, nil
}

func (pi *pagesIndex) save() error {
	data := &pagesIndexData{
		items: pi.items,
	}

	bytes, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(pi.path, bytes, 0644)
}

func (pi *pagesIndex) addArticles(paths []string) {
	for _, path := range paths {
		pi.addArticle(path)
	}
}

func (pi *pagesIndex) addArticle(path string) {
	dir, filename := filepath.Split(path)
	dir = strings.TrimSuffix(dir, string(filepath.Separator))
	section := pi.findOrCreateSection(dir)
	if !slices.Contains(section.Files, filename) {
		section.Files = append(section.Files, filename)
	}
}

func (pi *pagesIndex) findOrCreateSection(dir string) *pagesIndexItem {
	for _, section := range pi.items {
		if section.Key == dir {
			return section
		}
	}
	section := &pagesIndexItem{
		Key:     dir,
		Visible: true,
		Title:   dir,
		Files:   nil,
	}
	pi.items = append(pi.items, section)
	return section
}

func (pi *pagesIndex) getArticleSection(path string) *pagesIndexItem {
	dir := filepath.Dir(path)
	dir = strings.TrimSuffix(dir, string(filepath.Separator))

	for _, item := range pi.items {
		if item.Key == dir {
			return item
		}
	}
	return nil
}

func (pi *pagesIndex) listSections() []*pagesIndexItem {
	return pi.items
}
