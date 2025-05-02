package app

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

type Generator struct {
	sources   Sources
	targets   Targets
	converter MarkdownConverter
	cache     GeneratorCache
	logger    GeneratorLogger
}

func NewGenerator(sources Sources, targets Targets, converter MarkdownConverter, cache GeneratorCache, logger GeneratorLogger) *Generator {
	return &Generator{
		sources:   sources,
		targets:   targets,
		converter: converter,
		cache:     cache,
		logger:    logger,
	}
}

func (g *Generator) Generate() error {
	for i, path := range g.sources.ListFiles(Markdown) {
		fmt.Printf("%d. %s\n", i, path)
	}

	err := g.copyAssets()
	if err != nil {
		return err
	}

	err = g.convertMarkdownFiles()
	if err != nil {
		return err
	}

	err = g.cache.SaveCache()
	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) convertMarkdownFiles() error {
	for _, path := range g.sources.ListFiles(Markdown) {
		err := g.convertMarkdownFile(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) copyAssets() error {
	for _, path := range g.sources.ListFiles(Image) {
		err := g.copyAssetFile(path)
		if err != nil {
			return err
		}
	}
	for _, path := range g.sources.ListFiles(JavaScript) {
		err := g.copyAssetFile(path)
		if err != nil {
			return err
		}
	}
	for _, path := range g.sources.ListFiles(StyleSheet) {
		err := g.copyAssetFile(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) convertMarkdownFile(path string) error {
	srcAbsPath := filepath.Join(g.sources.Root(), path)
	html, err := g.converter.ConvertToHtml(srcAbsPath)
	if err != nil {
		return err
	}

	outputPath := replaceFileExtension(path, ".html")
	err = g.targets.Write(outputPath, html)
	if err != nil {
		return err
	}

	g.logger.LogConvertedFile(path, outputPath)
	return nil
}

func (g *Generator) copyAssetFile(path string) error {
	srcAbsPath := filepath.Join(g.sources.Root(), path)
	src, err := os.OpenFile(srcAbsPath, os.O_RDONLY, 0)
	if err != nil {
		return xerrors.Errorf("failed to open source file %s: %w", path, err)
	}

	err = g.targets.Copy(path, src)
	if err != nil {
		return err
	}

	g.logger.LogCopiedFile(path)
	return nil
}
