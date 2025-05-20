package templates

import (
	"bytes"
	"html/template"
	"path/filepath"

	"golang.org/x/xerrors"

	"sitegenerator/app"
)

type siteTemplates struct {
	templatesDir string
	templates    *template.Template
}

func ParseSiteTemplates(callbacks FuncCallbacks, templatesDir string) (app.SiteTemplates, error) {
	funcs := createFuncMap(callbacks)
	templates := template.New("sitegenerator").Funcs(funcs)

	_, err := templates.ParseGlob(
		filepath.Join(templatesDir, "*.html"),
	)

	if err != nil {
		return nil, xerrors.Errorf("failed to parse site templates at %q: %w", templatesDir, err)
	}

	return &siteTemplates{
		templatesDir: templatesDir,
		templates:    templates,
	}, nil
}

func (t *siteTemplates) GenerateArticlePage(d *app.ArticlePageDetails) ([]byte, error) {
	return t.generatePage("article_page.html", toArticlePageVars(d))
}

func (t *siteTemplates) GenerateSectionPage(d *app.SectionPageDetails) ([]byte, error) {
	return t.generatePage("section_page.html", toSectionPageVars(d))
}

func (t *siteTemplates) GenerateIndexPage(d *app.IndexPageData) ([]byte, error) {
	return t.generatePage("index_page.html", toIndexPageVars(d))
}

func (t *siteTemplates) generatePage(name string, data any) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = xerrors.Errorf("could not generate page %q: %w", name, r)
		}
	}()
	bytes, err = t.generatePageImpl(name, data)
	return
}

func (t *siteTemplates) generatePageImpl(name string, data any) ([]byte, error) {
	template := t.templates.Lookup(name)
	if template == nil {
		return nil, xerrors.Errorf("could not find template %q at %q", name, t.templatesDir)
	}

	var buffer bytes.Buffer
	err := template.Execute(&buffer, data)
	if err != nil {
		return nil, xerrors.Errorf("could not execute template %q: %w", name, err)
	}
	return buffer.Bytes(), nil
}
