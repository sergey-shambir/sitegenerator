package data

import (
	"encoding/json"
	"os"

	"golang.org/x/xerrors"
)

type PagesMetadataCache struct {
	pages map[string]*PageMetadata
}

func LoadPagesMetadataCache(path string) (*PagesMetadataCache, error) {
	cache := &PagesMetadataCache{
		pages: make(map[string]*PageMetadata),
	}
	data, err := os.ReadFile(path)

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

func SavePagesMetadataCache(cache *PagesMetadataCache, path string) error {
	data, err := json.Marshal(cache.pages)
	if err != nil {
		return xerrors.Errorf("failed to format metadata cache: %w", err)
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return xerrors.Errorf("failed to save metadata cache: %w", err)
	}
	return nil
}

func (c *PagesMetadataCache) GetPageMetadata(path string) (*PageMetadata, error) {
	metadata, ok := c.pages[path]
	if !ok {
		var err error
		metadata, err = ParsePageMetadata(path)
		if err != nil {
			return nil, err
		}
		c.pages[path] = metadata
	}
	return metadata, nil
}
