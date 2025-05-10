package app

import (
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

type Generator struct {
	sources   Sources
	targets   Targets
	converter Converter
	project   Project
	logger    GeneratorLogger
}

func NewGenerator(sources Sources, targets Targets, converter Converter, project Project, logger GeneratorLogger) *Generator {
	return &Generator{
		sources:   sources,
		targets:   targets,
		converter: converter,
		project:   project,
		logger:    logger,
	}
}

func (g *Generator) Generate() error {
	// TODO: convert markdown pages using templates
	// TODO: generate section pages
	// TODO: generate main page (index.html)

	err := g.project.AddArticles(g.sources.ListFiles(Markdown))
	if err != nil {
		return err
	}

	err = g.copyAssets()
	if err != nil {
		return err
	}

	err = g.convertMarkdownFiles()
	if err != nil {
		return err
	}

	err = g.convertSassFiles()
	if err != nil {
		return err
	}

	err = g.project.Save()
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

func (g *Generator) convertSassFiles() error {
	for _, path := range g.sources.ListFiles(Sass) {
		err := g.convertSassFile(path)
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
	html, err := g.converter.ConvertMarkdownToHtml(srcAbsPath)
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

func (g *Generator) convertSassFile(path string) error {
	srcAbsPath := filepath.Join(g.sources.Root(), path)
	html, err := g.converter.ConvertSassToCss(srcAbsPath)
	if err != nil {
		return err
	}

	outputPath := replaceFileExtension(path, ".css")
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
