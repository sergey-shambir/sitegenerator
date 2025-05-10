package project

import (
	"golang.org/x/xerrors"

	"sitegenerator/app"
)

type project struct {
	index *pagesIndex
	cache *generatorCache
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

func (p *project) ListSections() []*app.SectionPageData {
	return p.index.listSections()
}

func (p *project) GetArticleSection(path string) *app.SectionPageData {
	return p.index.getArticleSection(path)
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
