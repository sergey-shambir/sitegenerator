package app

import (
	"fmt"
	"os"

	"golang.org/x/xerrors"
)

type Generator struct {
	sources Sources
	targets Targets
	cache   GeneratorCache
}

func NewGenerator(sources Sources, targets Targets, cache GeneratorCache) *Generator {
	return &Generator{
		sources: sources,
		targets: targets,
		cache:   cache,
	}
}

func (g *Generator) Generate() error {
	for i, path := range g.sources.ListFiles(Markdown) {
		fmt.Printf("%d. %s\n", i, path)
	}

	err := g.CopyAssets()
	if err != nil {
		return err
	}

	err = g.cache.SaveCache()
	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) CopyAssets() error {
	for _, path := range g.sources.ListFiles(Image) {
		err := g.CopyAssetFile(path)
		if err != nil {
			return err
		}
	}
	for _, path := range g.sources.ListFiles(JavaScript) {
		err := g.CopyAssetFile(path)
		if err != nil {
			return err
		}
	}
	for _, path := range g.sources.ListFiles(StyleSheet) {
		err := g.CopyAssetFile(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) CopyAssetFile(path string) error {
	src, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return xerrors.Errorf("failed to open source file %s: %w", path, err)
	}
	return g.targets.Copy(path, src)
}
