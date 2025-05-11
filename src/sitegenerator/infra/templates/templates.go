package templates

import (
	"bytes"
	"html/template"
	"path/filepath"

	"golang.org/x/xerrors"

	"sitegenerator/app"
)

type siteTemplates struct {
	tpl *template.Template
}

func ParseSiteTemplates(templatesDir string) (app.SiteTemplates, error) {
	funcs := createFuncMap()
	tpl, err := template.New("sitegenerator").Funcs(funcs).ParseGlob(
		filepath.Join(templatesDir, "*.html"),
	)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse site templates at '%s': %w", templatesDir, err)
	}

	return &siteTemplates{tpl}, nil
}

func (t *siteTemplates) GenerateArticlePage(d app.ArticlePageDetails) ([]byte, error) {
	return t.generatePage("article_page.html", toArticlePageVars(d))
}

func (t *siteTemplates) GenerateSectionPage(d app.SectionPageDetails) ([]byte, error) {
	return t.generatePage("section_page.html", toSectionPageVars(d))
}

func (t *siteTemplates) GenerateIndexPage(d app.IndexPageData) ([]byte, error) {
	return t.generatePage("index_page.html", toIndexPageVars(d))
}

func (t *siteTemplates) generatePage(name string, data any) ([]byte, error) {
	var buffer bytes.Buffer
	err := t.tpl.Lookup(name).Execute(&buffer, data)
	if err != nil {
		return nil, xerrors.Errorf("failed to apply site template '%s': %w", name, err)
	}
	return buffer.Bytes(), nil
}
