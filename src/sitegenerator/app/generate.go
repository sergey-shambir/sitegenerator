package app

import (
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

type Generator struct {
	sources   Sources
	targets   Targets
	templates SiteTemplates
	converter Converter
	project   Project
	logger    GeneratorLogger
}

func NewGenerator(sources Sources, targets Targets, converter Converter, templates SiteTemplates, project Project, logger GeneratorLogger) *Generator {
	return &Generator{
		sources:   sources,
		targets:   targets,
		templates: templates,
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

	err = g.convertSassFiles()
	if err != nil {
		return err
	}

	err = g.generatePages()
	if err != nil {
		return err
	}

	err = g.project.Save()
	if err != nil {
		return err
	}

	return nil
}

// NOTE: Генерация статей должна быть в конце, чтобы иметь доступ к файлам, которые будут созданы в процессе конвертации
func (g *Generator) generatePages() error {
	for _, path := range g.sources.ListFiles(Markdown) {
		err := g.generateArticlePage(path)
		if err != nil {
			return err
		}
	}

	sections := g.project.ListSections()
	for _, section := range sections {
		err := g.generateSectionPage(section)
		if err != nil {
			return err
		}
	}

	err := g.generateIndexPage(sections)
	if err != nil {
		return err
	}

	g.project.ListSections()

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

func (g *Generator) generateArticlePage(path string) error {
	srcAbsPath := filepath.Join(g.sources.Root(), path)
	articleHtml, err := g.converter.ConvertMarkdownToHtml(srcAbsPath)
	if err != nil {
		return err
	}

	html, err := g.templates.GenerateArticlePage(&ArticlePageDetails{
		IsVisible: g.project.IsVisibleArticle(path),
		Meta:      g.project.GetArticleMetadata(path),
		Content:   articleHtml,
	})
	if err != nil {
		return err
	}

	outputPath := ReplaceFileExtension(path, HtmlExt)
	err = g.targets.Write(outputPath, html)
	if err != nil {
		return err
	}

	g.logger.LogConvertedFile(path, outputPath)
	return nil
}

func (g *Generator) generateSectionPage(section *SectionPageDetails) error {
	outputPath := UrlPathToPath(section.Url)
	html, err := g.templates.GenerateSectionPage(section)
	if err != nil {
		return err
	}

	err = g.targets.Write(outputPath, html)
	if err != nil {
		return err
	}

	g.logger.LogGeneratedPage(outputPath)
	return nil
}

func (g *Generator) generateIndexPage(sections []*SectionPageDetails) error {
	sectionsData := make([]UrlAndValue[string], 0, len(sections))
	for _, section := range sections {
		if section.IsVisible {
			sectionsData = append(sectionsData, UrlAndValue[string]{
				Url:   section.Url,
				Value: section.Title,
			})
		}
	}

	html, err := g.templates.GenerateIndexPage(&IndexPageData{
		Sections: sectionsData,
	})
	if err != nil {
		return err
	}

	err = g.targets.Write(IndexHtmlPath, html)
	if err != nil {
		return err
	}

	g.logger.LogGeneratedPage(IndexHtmlPath)
	return nil
}

func (g *Generator) convertSassFile(path string) error {
	srcAbsPath := filepath.Join(g.sources.Root(), path)
	html, err := g.converter.ConvertSassToCss(srcAbsPath)
	if err != nil {
		return err
	}

	outputPath := ReplaceFileExtension(path, CssExt)
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
