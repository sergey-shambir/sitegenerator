package project

import (
	"path/filepath"

	"golang.org/x/xerrors"

	"sitegenerator/app"
)

type project struct {
	index *pagesIndex
	cache *metadataCache
}

func (p *project) AddArticles(paths []string) error {
	err := p.cache.addArticles(paths)
	if err != nil {
		return err
	}
	p.index.addArticles(paths)
	return nil
}

func (p *project) Save() error {
	err := p.index.save()
	if err != nil {
		return err
	}
	return p.cache.save()
}

func (p *project) ListSections() []*app.SectionPageDetails {
	sections := p.index.listSections()
	results := make([]*app.SectionPageDetails, len(sections))
	for i, section := range sections {
		results[i] = p.toSectionPageDetails(section)
	}
	return results
}

func (p *project) IsVisibleArticle(path string) bool {
	section := p.index.getArticleSection(path)
	return section != nil && section.Visible
}

func (p *project) GetArticleMetadata(path string) *app.ArticleMetadata {
	return p.cache.getArticleMetadata(path)
}

func (p *project) toSectionPageDetails(section *pagesIndexItem) *app.SectionPageDetails {
	pages := make([]app.UrlAndValue[*app.ArticleMetadata], len(section.Files))
	for i, file := range section.Files {
		pages[i].Url = "/" + section.Key + "/" + app.ReplaceFileExtension(file, app.HtmlExt)
		pages[i].Value = p.cache.getArticleMetadata(filepath.Join(section.Key, file))
	}

	return &app.SectionPageDetails{
		Url:       "/" + section.Key + app.HtmlExt,
		IsVisible: section.Visible,
		Title:     section.Title,
		Pages:     pages,
	}
}

func LoadProject(sourceDir, indexPath, cachePath string) (app.Project, error) {
	cache, err := loadGeneratorCache(sourceDir, cachePath)
	if err != nil {
		return nil, xerrors.Errorf("failed to load project: %w", err)
	}

	index, err := loadPagesIndex(indexPath)
	if err != nil {
		return nil, xerrors.Errorf("failed to load project: %w", err)
	}

	return &project{
		index: index,
		cache: cache,
	}, nil
}
